package stringext

import (
	"fmt"
	"testing"
)

func Test_CutString(t *testing.T) {

	x0 := CutString("1234567890", 9)

	fmt.Println(x0)
}
