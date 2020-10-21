package filedir

import (
	"testing"
)

func Test_CheckFileIsExist(t *testing.T) {
	CheckFileIsExist("D:/install.ini")
}

func Test_WriteFileContent(t *testing.T) {
	_,er:=WriteFileContent("D:/aaa/aaa/aaa/aaa.txt","2222",false)
	if nil!=er{
		t.Error(er)
	}
}