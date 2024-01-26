package jsonext

import (
	"encoding/json"

	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
)

func init() {

}

func Sonic2JsonBytes(t interface{}) (*[]byte, error) {
	bytes, err := sonic.Marshal(t)
	return &bytes, err
}

func Sonic2Str(t interface{}) (string, error) {
	bytes, err := Sonic2JsonBytes(t)
	if nil != err {
		return "", err
	}
	return string(*bytes), err
}

func Sonic2StrOk(t interface{}) string {
	str, _ := Sonic2Str(t)
	return str
}

func Sonic2Obj(str *string, t interface{}) error {
	err := sonic.UnmarshalString(*str, &t)
	if err != nil {
		return err
	}
	return nil
}

func Sonic2AstNode(str string) (ast.Node, error) {
	root, err := sonic.GetFromString(str)
	return root, err
}

// ToStrOk 对象转换为Str
// t 任意对象，注意取地址传入
func ToStrOk(t interface{}) string {
	result, _ := ToStr(t)
	return result
}

// ToStr 对象转换为Str
//
//	t 任意对象，注意取地址传入
func ToStr(t interface{}) (string, error) {
	jsonBytes, err := ToJSONBytes(t)
	if err != nil {
		return "", err
	}
	return string(*jsonBytes), nil
}

//	 ToJSONBytes 对象转换为[]byte
//
//		t 任意对象，注意取地址传入
func ToJSONBytes(t interface{}) (*[]byte, error) {
	jsonBytes, err := json.Marshal(&t)
	if err != nil {
		return nil, err
	}
	return &jsonBytes, nil
}

// ToObj Str转换为对象
//
//	str 字符串的引用
//	t  需要转换到的对象的引用
func ToObj(str *string, t interface{}) error {
	err := json.Unmarshal([]byte(*str), &t)
	if err != nil {
		return err
	}
	return nil
}

// ToMap 对象转换为map[string]interface{}
// t 任意对象，注意取地址传入
func ToMap(str *string) (r map[string]interface{}, e error) {
	jsonMap := make(map[string]interface{}, 0)
	err := json.Unmarshal([]byte(*str), &jsonMap)
	return jsonMap, err
}
