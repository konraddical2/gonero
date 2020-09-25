package gonero

import (
	"bytes"
	"crypto/ed25519"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/konraddical2/gonero/base58"
	"golang.org/x/crypto/sha3"
)

// valid leading bytes for addresses
var (
	masterAddrNetBytes = []byte{18, 53, 24}
	subAddrNetBytes    = []byte{42, 63, 36}
	intAddrNetBytes    = []byte{19, 54, 25}

	mainnet  = "main"
	testnet  = "test"
	stagenet = "stage"
	networks = []string{mainnet, testnet, stagenet}
	netIndex = map[string]int{mainnet: 0, testnet: 1, stagenet: 2}
)

// address regular expressions
var (
	addrRegex  = regexp.MustCompile("^[123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{95}$")
	iaddrRegex = regexp.MustCompile("^[123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{106}$")
)

type baseAddress struct {
	Addr          string
	Label         string
	validNetbytes []byte
	decoded       []byte
	regex         *regexp.Regexp
}

func (b *baseAddress) decode() error {
	decoded := base58.Decode(b.Addr)
	l := len(decoded)
	if l == 0 {
		return fmt.Errorf("could not decode %s", b.Addr)
	}
	// validate net byte
	netbyte := decoded[0]
	if !bytes.Contains(b.validNetbytes, []byte{netbyte}) {
		return fmt.Errorf("invalid net byte")
	}
	// validate checksum
	checksum := decoded[l-4:]

	h := sha3.NewLegacyKeccak256()
	_, err := h.Write(decoded[:l-4])
	if err != nil {
		return err
	}

	hash := h.Sum(nil)
	if !bytes.Equal(checksum, hash[:4]) {
		return fmt.Errorf("invalid checksum")
	}

	b.decoded = decoded
	return nil
}

func (b *baseAddress) Valid() bool {
	if !b.regex.Match([]byte(b.Addr)) {
		return false
	}
	if err := b.decode(); err != nil {
		return false
	}
	return true
}

func (b *baseAddress) Net() string {
	if len(b.decoded) == 0 {
		return ""
	}
	// if b.decoded is set, we know that this will not be -1
	i := bytes.Index(b.validNetbytes, []byte{b.decoded[0]})

	return networks[i]
}

func (b *baseAddress) SpendKey() string {
	if len(b.decoded) == 0 {
		return ""
	}
	return fmt.Sprintf("%x", b.decoded[1:33])
}

func (b *baseAddress) ViewKey() string {
	if len(b.decoded) == 0 {
		return ""
	}
	return fmt.Sprintf("%x", b.decoded[33:65])
}

func (b *baseAddress) IsMainnet() bool {
	return b.Net() == mainnet
}

func (b *baseAddress) IsStagenet() bool {
	return b.Net() == stagenet
}

func (b *baseAddress) IsTestnet() bool {
	return b.Net() == testnet
}

// Address is a Monero address
type Address struct {
	baseAddress
}

// NewAddress creates a new Address struct from a hex address string
func NewAddress(addr string) *Address {
	a := &Address{baseAddress: baseAddress{
		Addr: addr, validNetbytes: masterAddrNetBytes, regex: addrRegex},
	}
	err := a.decode()
	if err != nil {
		return nil
	}

	return a
}

// CheckPrivateViewKey checks if private view key matches this address.
func (a *Address) CheckPrivateViewKey(key string) bool {
	data, err := hex.DecodeString(key)
	if err != nil {
		return false
	}
	pk := ed25519.NewKeyFromSeed(data)
	public := []byte(pk.Public().(ed25519.PublicKey))
	fmt.Printf("%x  %x\n", a.decoded[33:65], public)
	return bytes.Equal(a.decoded[33:65], public)
}

// CheckPrivateSpendKey checks if private spend key matches this address.
func (a *Address) CheckPrivateSpendKey(key string) bool {
	data, err := hex.DecodeString(key)
	if err != nil {
		return false
	}
	pk := ed25519.NewKeyFromSeed(data)
	public := []byte(pk.Public().(ed25519.PublicKey))
	fmt.Printf("%x  %x\n", a.decoded[1:33], public)
	return bytes.Equal(a.decoded[1:33], public)
}

// WithPaymentID creates a new integrated address with the given hex string payment ID
func (a *Address) WithPaymentID(paymentID string) *IntegratedAddress {
	prefix := intAddrNetBytes[netIndex[a.Net()]]
	data := append([]byte{prefix}, a.decoded[1:65]...)

	// append payment id to data
	cleaned := strings.Replace(paymentID, "0x", "", -1)
	pid, err := strconv.ParseUint(cleaned, 16, 64)
	if err != nil {
		return nil
	}

	pidData := make([]byte, 8)
	binary.BigEndian.PutUint64(pidData, pid)
	data = append(data, pidData...)

	h := sha3.NewLegacyKeccak256()
	_, err = h.Write(data)
	if err != nil {
		return nil
	}
	checksum := h.Sum(nil)[:4]
	data = append(data, checksum...)

	return NewIntegratedAddress(base58.Encode(data))
}

// SubAddress is a Monero subaddress
type SubAddress struct {
	baseAddress
}

// NewSubAddress creates a new subaddress struct
func NewSubAddress(addr string) *SubAddress {
	a := &SubAddress{baseAddress: baseAddress{
		Addr: addr, validNetbytes: subAddrNetBytes, regex: addrRegex},
	}
	err := a.decode()
	if err != nil {
		return nil
	}

	return a
}

// IntegratedAddress is a Monero integrated address.
// A master address integrated with payment id (short one, max 64 bit).
// TODO check length of payment ID
type IntegratedAddress struct {
	baseAddress
}

// NewIntegratedAddress creates a new integrated address struct
func NewIntegratedAddress(addr string) *IntegratedAddress {
	a := &IntegratedAddress{baseAddress: baseAddress{
		Addr: addr, validNetbytes: intAddrNetBytes, regex: iaddrRegex},
	}
	err := a.decode()
	if err != nil {
		return nil
	}

	return a
}

// PaymentID returns the integrated payment id.
func (ia *IntegratedAddress) PaymentID() string {
	l := len(ia.decoded)
	if l == 0 {
		return ""
	}
	hexStr := hex.EncodeToString(ia.decoded[65 : l-4])
	return strings.TrimLeft(hexStr, "0")
}

// BaseAddress returns the base address without payment id.
func (ia *IntegratedAddress) BaseAddress() *Address {
	prefix := masterAddrNetBytes[netIndex[ia.Net()]]

	data := append([]byte{prefix}, ia.decoded[1:65]...)

	h := sha3.NewLegacyKeccak256()
	_, err := h.Write(data)
	if err != nil {
		return nil
	}

	checksum := h.Sum(nil)[:4]

	return NewAddress(base58.Encode(append(data, checksum...)))
}
