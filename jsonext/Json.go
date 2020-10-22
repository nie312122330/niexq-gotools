package jsonext

import "encoding/json"

func init() {

}

/*ToStr 对象转换为Str
 * t 任意对象，注意取地址传入
 */
func ToStr(t interface{}) (string, error) {
	jsonBytes, err := json.Marshal(&t)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

/*ToObj Str转换为对象
 * str 字符串的引用
 * t  需要转换到的对象的引用
 */
func ToObj(str *string,t interface{}) error{
	err:=json.Unmarshal([]byte(*str),&t)
	if err != nil {
		return  err
	}
	return  nil
}
