package tempconv_test

import (
	"ch12/tempconv"
	"fmt"
	"testing"
)

func TestTempConv(t *testing.T) {
	fmt.Println(tempconv.CToK(tempconv.BoilingC))
}
