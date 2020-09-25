package gonero

import (
	"fmt"
)

// AtomicXMR is XMR currency units
type AtomicXMR uint64

// Float converts XMR amount in atomic units
// to float format of XMR
func (a AtomicXMR) Float() FloatXMR {
	return FloatXMR(a) / 1e12
}

// Decimal converts XMR amount in atomic units
// to decimal string
func (a AtomicXMR) Decimal() string {
	s := fmt.Sprintf("%013d", a)
	l := len(s)
	return s[:l-12] + "." + s[l-12:]
}

// FloatXMR is
type FloatXMR float64

// Decimal converts XMR amount in float units
// to decimal string
func (f FloatXMR) Decimal() string {
	return fmt.Sprintf("%.12f", f)
}

// Atomic converts XMR amount in float units
// to atomic units
func (f FloatXMR) Atomic() AtomicXMR {
	return AtomicXMR(f * 1e12)
}
