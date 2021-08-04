package jdcsdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

const (
	getQrcode = "http://%s/api/login/qrcode"
)

type Body struct {
	ResultCode string `json:"resultCode"`
	ResultMsg  string `json:"resultMsg"`
	Url        string `json:"url"`
	Img        string `json:"img"`
}

func GetQrcode() (*Body, error) {
	uri := fmt.Sprintf(getQrcode, host)
	result := &struct {
		ResultCode string `json:"resultCode"`
		ResultMsg  string `json:"resultMsg"`
	}{}
	resp, err := client.Net.Get(uri).End(nil, result)
	if err != nil {
		return nil, netError(err)
	}
	if result.ResultCode != respOk {
		return nil, errors.New(result.ResultMsg)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resultqrcode := Body{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return &resultqrcode, nil
}
