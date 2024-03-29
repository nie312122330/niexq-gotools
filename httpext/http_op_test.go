package httpext

import (
	"net/url"
	"testing"
	"time"

	"github.com/nie312122330/niexq-gotools/fileext"
)

func TestGet(t *testing.T) {

	result, _ := GetText("http://www.baidu.com", time.Second*10)
	t.Log(result)
}

func TestPost(t *testing.T) {
	result, _ := PostJSON("http://www.baidu.com", "", time.Second*10)
	t.Log(result)
}

func TestProxy(t *testing.T) {
	SetProxy("http://127.0.0.1:19180")
	result, _ := GetText("http://www.google.com/", time.Second*10)
	t.Log(result)
}

func TestPostForm(t *testing.T) {
	data := url.Values{
		"path": {"asasdfds/2022/05/18/82319253CCA446BB9FB44DBE0E22BBF6.txt"},
	}
	str, err := PostForm("http://192.168.0.253:10081/pub/fileExist.do", data, time.Second*10)
	if nil != err {
		t.Log(err.Error())
	} else {
		t.Log(str)
	}
}

func TestPostFile(t *testing.T) {
	fileBytes, err := fileext.ReadFileByte("http_op.go")
	extData := url.Values{"filepath": {"test.txt"}}
	url := "http://127.0.0.1:29833/upfile"

	str, err := PostFile(url, &fileBytes, "file", "file.txt", extData, time.Second*10)
	if nil != err {
		t.Log(err.Error())
	} else {
		t.Log(str)
	}
}
