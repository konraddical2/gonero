package base58

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math"
)

const (
	b58base              = 58
	fullBlockSize        = 8
	fullEncodedBlockSize = 11
)

var (
	alphabet          = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
	uint64Max         = math.Pow(2, 64)
	encodedBlockSizes = []int{0, 2, 3, 5, 6, 7, 9, 10, 11}
)

func getEncodedBlockSizesIndex(l int) int {
	for i := 0; i < len(encodedBlockSizes); i++ {
		if encodedBlockSizes[i] == l {
			return i
		}
	}
	return -1
}

func hexToBin(hexStr string) ([]byte, error) {
	return hex.DecodeString(hexStr)
}

func binToHex(bin []byte) string {
	return hex.EncodeToString(bin)
}

func uint8beTo64(data []byte) (uint64, error) {
	l := len(data)
	if !((1 <= l) && (l <= 8)) {
		return 0, fmt.Errorf("invalid input length: %d", len(data))
	}
	res := uint64(0)
	for _, b := range data {
		res = res<<8 | uint64(b)
	}
	return res, nil
}

func uint64To8be(num uint64, size uint) ([]byte, error) {
	if (size < 1) || (size > 8) {
		return nil, fmt.Errorf("invalid input length: %d", size)
	}
	res := make([]byte, size)
	twopow8 := uint64(256)
	for i := int(size) - 1; i > -1; i-- {
		res[i] = byte(num % twopow8)
		num = num / twopow8
	}
	return res, nil
}

func encodeBlock(data, buf []byte, index int) ([]byte, error) {
	lData := len(data)

	if (lData < 1) || (lData > fullEncodedBlockSize) {
		return nil, fmt.Errorf("invalid block length: %d", lData)
	}
	num, err := uint8beTo64(data)
	if err != nil {
		return nil, err
	}

	i := encodedBlockSizes[lData] - 1

	for num > 0 {
		remainder := num % b58base
		num = num / b58base
		buf[index+i] = alphabet[remainder]
		i = i - 1
	}
	return buf, nil
}

// Encode hexadecimal string as base58 (ex: encoding a Monero address).
func Encode(data []byte) string {

	lData := len(data)

	if lData == 0 {
		return ""
	}
	fullBlockCount := lData / fullBlockSize
	lastBlockSize := lData % fullBlockSize
	resSize := fullBlockCount*fullEncodedBlockSize + encodedBlockSizes[lastBlockSize]

	res := make([]byte, resSize) // [alphabet[0]]*)

	for i := 0; i < fullBlockCount; i++ {
		res, _ = encodeBlock(data[(i*fullBlockSize):(i*fullBlockSize+fullBlockSize)], res, i*fullEncodedBlockSize)
	}
	if lastBlockSize > 0 {
		res, _ = encodeBlock(data[(fullBlockCount*fullBlockSize):(fullBlockCount*fullBlockSize+lastBlockSize)], res, fullBlockCount*fullEncodedBlockSize)
	}
	return string(res)
}

func decodeBlock(data, buf []byte, index int) ([]byte, error) {
	lData := len(data)

	if (lData < 1) || (lData > fullEncodedBlockSize) {
		return nil, fmt.Errorf("Invalid block length: %d", lData)
	}
	resSize := getEncodedBlockSizesIndex(lData)
	if resSize <= 0 {
		return nil, fmt.Errorf("Invalid block size: %d", resSize)
	}
	resNum := 0
	order := 1
	for i := lData - 1; i > -1; i-- {
		digit := bytes.Index(alphabet, []byte{data[i]})
		if digit < 0 {
			return nil, fmt.Errorf("invalid symbol: %d", data[i])
		}
		product := order*digit + resNum
		if float64(product) > uint64Max {
			return nil, fmt.Errorf("Overflow: %d * %d + %d = %d", order, digit, resNum, product)
		}
		resNum = product
		order = order * b58base
	}

	if (resSize < fullBlockSize) && (math.Pow(2, float64(8*resSize)) <= float64(resNum)) {
		return nil, fmt.Errorf("Overflow: %d doesn't fit in %d bit(s)", resNum, resSize)
	}
	tmpBuf, err := uint64To8be(uint64(resNum), uint(resSize))
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(tmpBuf); i++ {
		buf[i+index] = tmpBuf[i]
	}

	return buf, nil
}

// Decode a base58 string (ex: a Monero address) into hexadecimal form.
func Decode(b58Str string) []byte {

	enc := []byte(b58Str) //bytearray(enc, encoding="ascii")
	lEnc := len(enc)

	if lEnc == 0 {
		return []byte{}
	}

	fullBlockCount := lEnc / fullEncodedBlockSize
	lastBlockSize := lEnc % fullEncodedBlockSize

	//try:
	lastBlockDecodedSize := getEncodedBlockSizesIndex(lastBlockSize)
	//except ValueError:
	//    raise ValueError("Invalid encoded length: %d" % lEnc)

	dataSize := fullBlockCount*fullBlockSize + lastBlockDecodedSize

	data := make([]byte, dataSize)
	for i := 0; i < fullBlockCount; i++ {
		data, _ = decodeBlock(enc[(i*fullEncodedBlockSize):(i*fullEncodedBlockSize+fullEncodedBlockSize)], data, i*fullBlockSize)
	}

	if lastBlockSize > 0 {
		data, _ = decodeBlock(enc[(fullBlockCount*fullEncodedBlockSize):(fullBlockCount*fullEncodedBlockSize+lastBlockSize)], data, fullBlockCount*fullBlockSize)
	}
	return data
}
