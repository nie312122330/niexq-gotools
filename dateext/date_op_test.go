package dateext

import "testing"

func Test_CurDateStr(t *testing.T) {
	str, _ := Now().Format("yyyyMMdd")
	t.Log(str)
}
