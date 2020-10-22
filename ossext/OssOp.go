package ossop

import (
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go.uber.org/zap"
	"io"
	"strings"
)

var logger zap.Logger

func init() {

}

//OssConf 配置文件
type OssConf struct {
	BucketName      string `json:"bucketName"`
	Endpoint        string `json:"endpoint"`
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	OssKeyPrefix    string `json:"ossKeyPrefix"`
}

//CreateOssClient 创建OssClient
func CreateOssClient(conf *OssConf) *oss.Client {
	client, err := oss.New(conf.Endpoint, conf.AccessKeyId, conf.AccessKeySecret)
	if err != nil {
		panic(err)
	}
	return client
}

//CreateOssClient 创建OssBucket
func CreateOssBucket(conf *OssConf) *oss.Bucket {
	client := CreateOssClient(conf)
	bucket, err := client.Bucket(conf.BucketName)
	if err != nil {
		panic(err)
	}
	return bucket
}

//PubObjectFromFile 上传文件
//  dir 目录，不能以/开始，必须以/结尾
func PubObjectFromFile(bucket *oss.Bucket, objKey string, filePath string) error {
	err := bucket.PutObjectFromFile(objKey, filePath)
	if nil != err {
		return err
	}
	return nil
}

//PubObjectFromIoReader 上传文件
//  dir 目录，不能以/开始，必须以/结尾
func PubObjectFromIoReader(bucket *oss.Bucket, objKey string, ioReader *io.Reader) error {
	err := bucket.PutObject(objKey, *ioReader)
	if nil != err {
		return err
	}
	return nil
}

//DelObject 删除单个Object
func DelObject(bucket *oss.Bucket, key string) error {
	//如果为/结尾则需要检查目录是否为空
	if strings.HasSuffix(key, "/") {
		empty, err := DirIsEmpty(bucket, key)
		if nil != err {
			return err
		}
		if !empty {
			return errors.New("目录非空")
		}
	}
	err := bucket.DeleteObject(key)
	if nil != err {
		return err
	}
	return nil
}

//DirIsEmpty 目录是否非空
//  dir 目录，不能以/开始，必须以/结尾
func DirIsEmpty(bucket *oss.Bucket, dir string) (bool, error) {
	lsRes, err := bucket.ListObjects(oss.Prefix(dir), oss.MaxKeys(1))
	if nil != err {
		return false, err
	}
	if len(lsRes.Objects) == 0 {
		return true, nil
	}
	return false, nil
}

//ListDirs 指定目录下的所有对象（文件|目录）
//  bucket
//  prefix 指定前缀，不能以/开头,根目录为空字符串
func ListObjects(bucket *oss.Bucket, prefix string) ([]oss.ObjectProperties, error) {
	var result []oss.ObjectProperties
	marker := oss.Marker("")
	for {
		lsRes, err := bucket.ListObjects(oss.Prefix(prefix), marker, oss.MaxKeys(200))
		if err != nil {
			return result, err
		}
		for _, object := range lsRes.Objects {
			result = append(result, object)
		}
		marker = oss.Marker(lsRes.NextMarker)
		//如果已全部返回则中断
		if !lsRes.IsTruncated {
			break
		}
	}
	return result, nil
}

//ListDirs 指定目录下的所有目录
//  bucket
//  prefix 指定前缀，不能以/开头,根目录为空字符串
func ListDirs(bucket *oss.Bucket, prefix string) ([]string, error) {
	var result []string
	marker := oss.Marker("")
	delimiter := oss.Delimiter("/")
	for {
		lsRes, err := bucket.ListObjects(oss.Prefix(prefix), marker, oss.MaxKeys(200), delimiter)
		if err != nil {
			return result, err
		}
		for _, object := range lsRes.CommonPrefixes {
			result = append(result, object)
		}
		marker = oss.Marker(lsRes.NextMarker)
		//如果已全部返回则中断
		if !lsRes.IsTruncated {
			break
		}
	}
	return result, nil
}
