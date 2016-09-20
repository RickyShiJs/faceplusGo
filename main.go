/**
 * Facepp - Face++  GO SDK
 * Face Detection Api
 * @Doc http://www.faceplusplus.com.cn/detection_detect/
 * @author Ricky Shi <ricky@dorontech.com>
 * @since  2016-09-20
 * @version  1.0
 **/
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
	res := f.Execute("/detection/detect", params)
	var responseBody models.DetectResponseBody
	json.Unmarshal(res, &responseBody)
	fmt.Printf("%+v\n", responseBody)
	/*
		Print:
			{SessionId:c864ecfaab664aaf9a26c0f5325e62ce Url:http://qntp01.resource.bo123.tv/img/aHR0cDovL3N0YXRpYy5tLnl5LmNvbS9ncm91cDE2L00wMC9GMS8yNC90ejBNOVZlYm1NWUFBQUFBQUFCTXMxS2pObjQ3MjYuanBn-avatar330 ImgId:0eda8f01840ba30dc7ef827431304095 ImgWidth:330 ImgHeight:330 Face:[{FaceId:0334592e
		bb510b883e28221be48ebd44 Tag: Position:{Center:{X:63.939396 Y:37.878788} EyeLeft:{X:59.550907 Y:30.938787} EyeRight:{X:71.61788 Y:33.253635} MouthLeft:{X:59.178486 Y:44.858486} MouthRight:{X:66.96424 Y:46.071213} Nose:{X:64.1297 Y:40.433334} Width:22.424242 Height:22.424242} Attribu
		te:{Age:{Range:5 Value:24} Gender:{Confidence:56.0255 Value:Female} Glass:{Confidence:98.6067 Value:None} Pose:{PitchAngle:{Value:-0.082646} RollAngle:{Value:10.8593} YawAngle:{Value:-0.384357}} Race:{Confidence:96.52 Value:Asian} Smiling:{Value:0.318674}}}]}

	*/

}
