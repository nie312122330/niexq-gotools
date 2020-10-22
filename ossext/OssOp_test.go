package ossop

import (
	"fmt"
	"github.com/nie312122330/niexq-gotools/filedir"
	"github.com/nie312122330/niexq-gotools/jsonext"
	"testing"
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
	confJsonStr, _ := filedir.ReadFileContent("../main_conf.json")
	jsonext.ToObj(&confJsonStr, &ocf)
	return &ocf
}
