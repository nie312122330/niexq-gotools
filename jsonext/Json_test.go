package jsonext

import (
	"testing"
)

type Person struct {
	Name string `json:"name"`
	Age  int  `json:"age"`
}

func TestToString(t *testing.T) {
	person := Person{Name: "小强", Age: 18}
	result, _ := ToStr(&person)
	t.Log("JsonStr:"+result)
	if result!="{\"name\":\"小强\",\"age\":18}" {
		t.Error("解析失败")
	}
}

func TestToObj(t *testing.T) {
	jsonStr:="{\"name\":\"小强\",\"age\":18}"
	var p Person
	ToObj(&jsonStr,&p)

	t.Log("字符串转对象通过",p.Name,p.Age)

	jsonStr="[{\"name\":\"小强\",\"age\":18}]"
	var ps []Person
	ToObj(&jsonStr,&ps)

	t.Log("字符串转数据通过,长度[",len(ps),"]",ps[0].Age,ps[0].Name)

}