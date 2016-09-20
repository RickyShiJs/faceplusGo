/**
 * Facepp - Face++  GO SDK
 * Face Detection Api
 * @Doc http://www.faceplusplus.com.cn/detection_detect/
 * @author Ricky Shi <ricky@dorontech.com>
 * @since  2016-09-20
 * @version  1.0
 **/
package api

import (
	"faceplusGo/models"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/go-querystring/query"
)

// FaceApi class
type FaceApi struct {
	server    string
	apiKey    string
	apiSecret string
}

var (
	apiKey    = "<Api_Key>"
	apiSecret = "<Api_Secret>"
	server    = "http://apicn.faceplusplus.com/v2"
)

func (f *FaceApi) Execute(method string, body models.DetectRequestBody) []byte {
	body.ApiKey = apiKey
	body.AppSecret = apiSecret
	res := f.Request(server+method, body)
	return res
}

func (f *FaceApi) Request(requestUrl string, body models.DetectRequestBody) []byte {
	v, _ := query.Values(body)
	url := requestUrl + "?" + v.Encode()
	resp, err := http.Get(url)
	if err != nil {
		panic("call api fail")
	}

	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("the response could not be read into memory")
	}
	return resBody
}
