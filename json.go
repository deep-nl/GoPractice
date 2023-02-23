package main

import (
	"encoding/json"
	"fmt"
	"go-base/logging"
	"io"
	"strings"
)

type Stu struct {
	Name  string `json:"name"`
	Age   int
	HIgh  bool
	sex   string
	Class *Class `json:"class"`
}

type Class struct {
	Name  string
	Grade int
}

func main() {

	//实例化一个数据结构，用于生成json字符串
	stu := Stu{
		Name: "张三",
		Age:  18,
		HIgh: true,
		sex:  "男",
	}

	//指针变量
	cla := new(Class)
	cla.Name = "1班"
	cla.Grade = 3
	stu.Class = cla

	fmt.Println(stu)
	//Marshal失败时err!=nil
	jsonStu, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("生成json字符串错误")
	}

	//jsonStu是[]byte类型，转化成string类型便于查看
	fmt.Println(string(jsonStu))
	ExampleDecoder()
}

/*
上面都是使用的UnMarshall解析的JSON数据，如果JSON数据的载体是打开的文件或者HTTP请求体这种数据流（他们都是io.Reader的实现），
我们不必把JSON数据读取出来后再去调用encode/json包的UnMarshall方法，包提供的Decode方法可以完成读取数据流并解析JSON数据最后填充变量的操作。
*/

func ExampleDecoder() {
	const jsonStream = `
    {"Name": "Ed", "Text": "Knock knock."}
    {"Name": "Sam", "Text": "Who's there?"}
    {"Name": "Ed", "Text": "Go fmt."}
    {"Name": "Sam", "Text": "Go fmt who?"}
    {"Name": "Ed", "Text": "Go fmt yourself!"}
`
	type Message struct {
		Name, Text string
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			logging.Logger.Errorln(err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Text)
	}
}
