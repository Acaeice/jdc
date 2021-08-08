package jdcsdk

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	getQrcode = "http://%s/api/login/qrcode"
	getQuery  = "http://%s/api/login/query"
)

type Body struct {
	ResultCode string `json:"resultCode"`
	ResultMsg  string `json:"resultMsg"`
	Url        string `json:"url"`
	Img        string `json:"img"`
}

func GetQrcode() ([]byte, error) {
	uri := fmt.Sprintf(getQrcode, host)
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	// defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func GetQuery() ([]byte, error) {
	uri := fmt.Sprintf(getQuery, host)
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	// defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
