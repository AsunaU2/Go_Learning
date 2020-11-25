package weightconv

import "fmt"

// Kilogram 公斤
type Kilogram float64

// Pound 磅
type Pound float64

const (
	transferUnit float64 = 0.453592
)

// KToP 公斤转磅
func KToP(k Kilogram) Pound {
	return Pound(k / Kilogram(transferUnit))
}

// PToK 磅转公斤
func PToK(p Pound) Kilogram {
	return Kilogram(p * Pound(transferUnit))
}

func (k Kilogram) String() string {
	return fmt.Sprintf("%g kilogram", k)
}

func (p Pound) String() string {
	return fmt.Sprintf("%g Pound", p)
}
