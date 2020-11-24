package weightconv

import "fmt"

type Kilogram float64
type Pound float64

const (
	transferUnit float64 = 0.453592
)

func KToP(k Kilogram) Pound {
	return Pound(k / Kilogram(transferUnit))
}

func PToK(p Pound) Kilogram {
	return Kilogram(p * Pound(transferUnit))
}

func (k Kilogram) String() string {
	return fmt.Sprintf("%g kilogram", k)
}

func (p Pound) String() string {
	return fmt.Sprintf("%g Pound", p)
}
