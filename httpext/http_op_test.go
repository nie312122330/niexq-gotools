package httpext

import (
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	result, _ := GetText("http://www.baidu.com", time.Second*10)
	t.Log(result)
}

func TestPost(t *testing.T) {
	result, _ := PostJSON("http://www.baidu.com", "", time.Second*10)
	t.Log(result)
}
