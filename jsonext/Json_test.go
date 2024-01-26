package jsonext

import (
	"fmt"
	"testing"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestToString(t *testing.T) {
	person := Person{Name: "小强", Age: 18}
	result, _ := ToStr(&person)
	t.Log("JsonStr:" + result)
	if result != "{\"name\":\"小强\",\"age\":18}" {
		t.Error("解析失败")
	}
}

func TestToObj(t *testing.T) {
	jsonStr := "{\"name\":\"小强\",\"age\":18}"
	var p Person
	ToObj(&jsonStr, &p)

	t.Log("字符串转对象通过", p.Name, p.Age)

	jsonStr = "[{\"name\":\"小强\",\"age\":18}]"
	var ps []Person
	ToObj(&jsonStr, &ps)

	t.Log("字符串转数据通过,长度[", len(ps), "]", ps[0].Age, ps[0].Name)

}

func TestToMap(t *testing.T) {
	jsonStr := "{\"name\":\"小强\",\"age\":18}"
	r, _ := ToMap(&jsonStr)
	t.Log("字符串转数据通过,长度[", len(r), "]")
}

func TestSnoic(t *testing.T) {
	jsonStr := "{\"name\":\"小强\",\"age\":18}"
	var p Person
	Sonic2Obj(&jsonStr, &p)
	fmt.Println(p)

	jsonStr = "[{\"name\":\"小强\",\"age\":18}]"
	var ps []Person
	Sonic2Obj(&jsonStr, &ps)
	fmt.Println(ps)

	jsonStr = ToStrOk(ps)
	fmt.Println(jsonStr)

	root, _ := Sonic2AstNode(jsonStr)
	nodeArr, _ := root.ArrayUseNode()

	fmt.Println(nodeArr[0].GetByPath("name").String())
	fmt.Println(root.GetByPath(0, "name").String())

}
