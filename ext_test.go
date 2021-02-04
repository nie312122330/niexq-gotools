package niexqext

import (
	"fmt"
	"testing"
)

func TestUuidToString(t *testing.T) {
	x0 := UUID()
	fmt.Println(x0)
	x0 = UUIDUperStr()
	fmt.Println(x0)
}
