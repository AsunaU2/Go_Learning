package weightconv_test

import (
	"ch13/weightconv"
	"testing"
)

func TestWeightConv(t *testing.T) {
	t.Log(weightconv.KToP(100))
	t.Log(weightconv.PToK(100))
}
