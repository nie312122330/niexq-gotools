package fileext

import (
	"testing"
)

func Test_CheckFileIsExist(t *testing.T) {
	CheckFileIsExist("D:/install.ini")
}

func Test_WriteFileContent(t *testing.T) {
	_, er := WriteFileContent("D:/aaa/aaa/aaa/aaa.txt", "2222", false)
	if nil != er {
		t.Error(er)
	}
}

func TestReadFileContent(t *testing.T) {
	str, err := ReadFileContent("D:/aaa/aaa/aaa/aaa.txt")
	if nil != err {
		t.Log(err.Error())
	} else {
		t.Log(str)
	}
}

func TestPathSplit(t *testing.T) {
	t.Log(PathDir("d:/ss/sss/sss.exe"))
	t.Log(PathFileName("d:/ss/sss/sss.exe"))
	t.Log(PathFileNameWithSuffix("d:/ss/sss/sss.exe"))
	t.Log(PathFileSuffix("d:/ss/sss/sss.exe"))

	t.Log(JoinPath("./aaa", "bbbb"))
	t.Log(JoinPath("", "/ccc/", "tes.txe"))
	t.Log(JoinPath("./aaa/", "bbbb/", "/dddd/exex/a.txt"))

}
