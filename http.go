package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"unsafe"
)

type JsonPostSample struct {
}

func (this *JsonPostSample) SamplePost() {
	song := make(map[string]interface{})
	song["name"] = "李白"
	song["timelength"] = 128
	song["author"] = "李荣浩"
	bytesData, err := json.Marshal(song)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	reader := bytes.NewReader(bytesData)
	url := "http://localhost/echo.php"
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//byte数组直接转成string,优化内存
	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println(*str)
}
