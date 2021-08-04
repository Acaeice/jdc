package jdcsdk

import (
	"encoding/json"
	"fmt"
	"testing"
)

const (
	testhost = "127.0.0.1:8080"
)

func Test_GetQrcode(t *testing.T) {
	Init(testhost)
	res, err := GetQrcode()
	result := &Body{}
	if err != nil {
		fmt.Println("++++++++++++++++++++++++++++++++++++++++")
		fmt.Println("err:", err)
		// return
	}
	if err := json.Unmarshal(res, &result); err != nil {
		return
	}
	fmt.Println("-----------------------------------------")
	fmt.Println("res:", result)
}
