package jdcsdk

import (
	"fmt"
	"testing"
)

const (
	testhost = "localhost:8080"
)

func Test_GetQrcode(t *testing.T) {
	Init(testhost)
	res, err := GetQrcode()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("-----------------------------------------")
	fmt.Println("res:", res.Img)
}
