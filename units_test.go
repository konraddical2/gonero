package gonero

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	atomic  AtomicXMR
	float   FloatXMR
	decimal string
}

var tests = []test{
	{atomic: AtomicXMR(34000200000), float: FloatXMR(0.0340002), decimal: "0.034000200000"},
	{atomic: AtomicXMR(15e12), float: FloatXMR(15.), decimal: "15.000000000000"},
	{atomic: AtomicXMR(20000000000), float: FloatXMR(0.02), decimal: "0.020000000000"},
	{atomic: AtomicXMR(314e10), float: FloatXMR(3.14), decimal: "3.140000000000"},
}

func TestAtomicXMR(t *testing.T) {
	for _, tc := range tests {
		assert.Equal(t, tc.decimal, tc.atomic.Decimal(), "Decimal() formats properly")
		assert.Equal(t, tc.float, tc.atomic.Float(), "Float() formats properly")
	}
}

func TestFloatXMR(t *testing.T) {
	for _, tc := range tests {
		assert.Equal(t, tc.decimal, tc.float.Decimal(), "Decimal() formats properly")
		assert.Equal(t, tc.atomic, tc.float.Atomic(), "Atomic() formats properly")
	}
}
