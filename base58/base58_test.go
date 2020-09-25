package base58

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHexToBin(t *testing.T) {
	type test struct {
		input    string
		expected []byte
	}

	tests := []test{
		{input: "dead", expected: []byte{222, 173}},
		{input: "deadbeef", expected: []byte{222, 173, 190, 239}},
	}

	for _, tc := range tests {
		out, _ := hexToBin(tc.input)
		assert.Equal(t, tc.expected, out)
	}

	type errorTest struct {
		input string
	}

	etests := []errorTest{
		{input: "zzz"},
	}

	for _, tc := range etests {
		_, err := hexToBin(tc.input)
		assert.Error(t, err)
	}
}

func TestBinToHex(t *testing.T) {
	type test struct {
		input    []byte
		expected string
	}

	tests := []test{
		{input: []byte{222, 173}, expected: "dead"},
		{input: []byte{222, 173, 190, 239}, expected: "deadbeef"},
	}

	for _, tc := range tests {
		out := binToHex(tc.input)
		assert.Equal(t, tc.expected, out)
	}
}

func TestUint8beTo64(t *testing.T) {
	type test struct {
		input    []byte
		expected uint64
	}

	tests := []test{
		{input: []byte{222, 173}, expected: 57005},
		{input: []byte{222, 173, 190, 239}, expected: 3735928559},
		{input: []byte{222, 173, 222, 173}, expected: 3735936685},
	}

	for _, tc := range tests {
		out, _ := uint8beTo64(tc.input)
		assert.Equal(t, tc.expected, out)
	}
}

func TestUint64To8be(t *testing.T) {
	type test struct {
		input    uint64
		size     uint
		expected []byte
	}

	tests := []test{
		{input: 57005, size: 2, expected: []byte{222, 173}},
		{input: 3735928559, size: 4, expected: []byte{222, 173, 190, 239}},
		{input: 3735936685, size: 4, expected: []byte{222, 173, 222, 173}},
	}

	for _, tc := range tests {
		out, _ := uint64To8be(tc.input, tc.size)
		assert.Equal(t, tc.expected, out)
	}
}

func TestEncode(t *testing.T) {
	type test struct {
		input    string
		expected string
	}

	tests := []test{
		{input: "dead", expected: "Hwr"},
		{input: "beef", expected: "FXk"},
		{input: "deadbeef", expected: "6h8cQN"},
		{input: "1865fd9e473d77c9b0894118f312b7f94931079108a12018720840696fcd1f9bb7c97cdd4e0b7985be9e436a6f360947c795cb4e890bd89ca7e620a5960efd886465eea013",
			expected: "55hKAMnUWXaWXd4hUQ9jzQDF3sL7Yu7fML5FpWMFMouUXjxxK2y6oKJYtEf91Vf3ZGaPE6cHccRnbV5q96uC3ChmCKyGuXY"},
	}

	for _, tc := range tests {
		data, _ := hex.DecodeString(tc.input)
		out := Encode(data)
		assert.Equal(t, tc.expected, out)
	}
}

func TestDecode(t *testing.T) {
	type test struct {
		input    string
		expected string
	}

	tests := []test{
		{input: "Hwr", expected: "dead"},
		{input: "FXk", expected: "beef"},
		{input: "6h8cQN", expected: "deadbeef"},
		{input: "55hKAMnUWXaWXd4hUQ9jzQDF3sL7Yu7fML5FpWMFMouUXjxxK2y6oKJYtEf91Vf3ZGaPE6cHccRnbV5q96uC3ChmCKyGuXY",
			expected: "1865fd9e473d77c9b0894118f312b7f94931079108a12018720840696fcd1f9bb7c97cdd4e0b7985be9e436a6f360947c795cb4e890bd89ca7e620a5960efd886465eea013"},
	}

	for _, tc := range tests {
		out := Decode(tc.input)
		expected, _ := hex.DecodeString(tc.expected)
		assert.Equal(t, expected, out)
	}
}
