package ossext

import (
	"fmt"
	"testing"

	"github.com/nie312122330/niexq-gotools/fileext"
	"github.com/nie312122330/niexq-gotools/jsonext"
)

func TestCreateOssClient(t *testing.T) {
	ocf := getOssConf()
	client := CreateOssClient(ocf)
	fmt.Printf("CreateOssClient:%v", client)
}
func TestCreateOssBucket(t *testing.T) {
	ocf := getOssConf()
	bucket := CreateOssBucket(ocf)
	fmt.Printf("CreateOssBucket:%v", bucket)
}

func TestListObjects(t *testing.T) {
	ocf := getOssConf()
	bucket := CreateOssBucket(ocf)
	objects, _ := ListObjects(bucket, "")
	t.Log(objects)
}

func TestListDir(t *testing.T) {
	ocf := getOssConf()
	bucket := CreateOssBucket(ocf)
	objects, _ := ListDirs(bucket, "")
	t.Log(objects)
}

func getOssConf() *OssConf {
	var ocf OssConf
	confJSONStr, _ := fileext.ReadFileContent("../main_conf.json")
	jsonext.ToObj(&confJSONStr, &ocf)
	return &ocf
}
