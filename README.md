# faceplusGo
face++ detection sdk for Golang 人脸识别api
## Dependencies
###### Download and install
    go get github.com/google/go-querystring/query
## Quick Start
###### Download the sdk
 * go get github.com/RickyShiJs/faceplusGo

###### Input your  api_key and api_secret in api/faceApi.go
``` go
var (
	apiKey    = "<Api_Key>"
	apiSecret = "<Api_Secret>"
  ...
```

###### Run test demo by following command

``` bash
go build
./faceplusGo
```

## How to Use

``` go
package main

import (
	"faceplusGo/api"
	"faceplusGo/models"
	"fmt"
)

import "encoding/json"

func main() {
	var f api.FaceApi
	var params models.DetectRequestBody
	params.Url = "http://tpic.home.news.cn/xhCloudNewsPic/xhpic1501/M06/23/24/wKhTlVfGjmuEW0rbAAAAADcO2D8123.jpg"
	params.Attribute = "gender,age,race,smiling,glass,pose"
	params.Mode = "oneface"
	// More params supportted
	res := f.Execute("/detection/detect", params)
	var responseBody models.DetectResponseBody
	json.Unmarshal(res, &responseBody)
	fmt.Printf("%+v\n", responseBody)
	/*
		Print:
		{SessionId:6fa8dcd3aa0b43ffb1fada4286c44bfa Url:http://tpic.home.news.cn/xhCloudNewsPic/xhpic1501/M06/23/24/wKhTlVfGjmuEW0rbAAAAADcO2D8123.jpg ImgId:4771e7e44ff19973d89176eedc2a3765 ImgWidth:380 ImgHeight:455 Face:[{FaceId:3dd9648e191ae01315c448697af42c2e Tag: Position:{Center:{X:59
.210526 Y:19.56044} EyeLeft:{X:53.334736 Y:18.272549} EyeRight:{X:61.272106 Y:15.208923} MouthLeft:{X:56.951054 Y:24.726593} MouthRight:{X:63.446053 Y:22.08879} Nose:{X:60.967632 Y:20.624834} Width:18.421053 Height:15.384615} Attribute:{Age:{Range:5 Value:29} Gender:{Confidence:99.9
999 Value:Female} Glass:{Confidence:99.9344 Value:None} Pose:{PitchAngle:{Value:2e-06} RollAngle:{Value:-24.8042} YawAngle:{Value:21.697386}} Race:{Confidence:82.8659 Value:Asian} Smiling:{Value:70.6813}}}]}

	*/
}

```

## Files
###### models/models.go
It contains the go structs for DetectRequestBody and DetectResponseBody, it help us to get and set value by . method easily

###### api/faceApi.go
The main SDK

###### main.go
Sample for new user

## Limitation

* It is not whole face++ go api, only /detection/detect
* Only implement the /detection/detect with url mode, will refine to support image upload later
