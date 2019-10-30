package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func getRequest() {
	client := http.Client{}
	resp, err := client.Get("http://192.168.8.94:8080/api/actuator/health")
	if nil != err {
		fmt.Println("http 请求失败")
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	personMap := make(map[string]interface{})
	err1 := json.Unmarshal(buf.Bytes(), &personMap)
	for key, value := range personMap {
		if key == "status" {
			fmt.Println("系统状态运行好")
		}
		if key == "details" {
		}

		fmt.Println("index : ", key, " value : ", value)
	}
	fmt.Errorf(err1.Error())
	fmt.Println(resp.StatusCode)
}
