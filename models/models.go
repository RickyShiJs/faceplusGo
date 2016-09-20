/**
 * Facepp - Face++  GO SDK
 * Face Detection Api
 * @Doc http://www.faceplusplus.com.cn/detection_detect/
 * @author Ricky Shi <ricky@dorontech.com>
 * @since  2016-09-20
 * @version  1.0
 **/
package models

type DetectRequestBody struct {
	ApiKey    string `url:"api_key"`
	AppSecret string `url:"api_secret"`
	Url       string `url:"url"`
	Attribute string `url:"attribute,omitempty"`
	Mode      string `url:"mode,omitempty"`
	Async     bool   `url:"async,omitempty"`
	Tag       string `url:"tag,omitempty"`
}

type Coordinate struct {
	X float32
	Y float32
}
type Position struct {
	Center     Coordinate
	EyeLeft    Coordinate `json:"eye_left"`
	EyeRight   Coordinate `json:"eye_right"`
	MouthLeft  Coordinate `json:"mouth_left"`
	MouthRight Coordinate `json:"mouth_right"`
	Nose       Coordinate
	Width      float32 `json:"width"`
	Height     float32 `json:"height"`
}
type Age struct {
	Range int
	Value int
}
type Gender struct {
	Confidence float32
	Value      string
}
type Glass struct {
	Confidence float32
	Value      string
}
type PitchAngle struct {
	Value float32
}
type RollAngle struct {
	Value float32
}
type YawAngle struct {
	Value float32
}
type Pose struct {
	PitchAngle PitchAngle `json:"pitch_angle"`
	RollAngle  RollAngle  `json:"roll_angle"`
	YawAngle   YawAngle   `json:"yaw_angle"`
}
type Race struct {
	Confidence float32
	Value      string
}
type Smiling struct {
	Value float32
}
type Attribute struct {
	Age     Age
	Gender  Gender
	Glass   Glass
	Pose    Pose
	Race    Race
	Smiling Smiling
}
type Face struct {
	FaceId    string `json:"face_id"`
	Tag       string `json:"tag"`
	Position  Position
	Attribute Attribute
}
type DetectResponseBody struct {
	SessionId string `json:"session_id"`
	Url       string `json:"url"`
	ImgId     string `json:"img_id"`
	ImgWidth  int    `json:"img_width"`
	ImgHeight int    `json:"img_height"`
	Face      []Face `json:"face"`
}

/* Example for DetectResponseBody
{
    "face": [
        {
            "attribute": {
                "age": {
                    "range": 5,
                    "value": 23
                },
                "gender": {
                    "confidence": 99.9999,
                    "value": "Female"
                },
                "glass": {
                    "confidence": 99.945,
                    "value": "None"
                },
                "pose": {
                    "pitch_angle": {
                        "value": 17
                    },
                    "roll_angle": {
                        "value": 0.735735
                    },
                    "yaw_angle": {
                        "value": -2
                    }
                },
                "race": {
                    "confidence": 99.6121,
                    "value": "Asian"
                },
                "smiling": {
                    "value": 4.86501
                }
            },
            "face_id": "17233b4b1b51ac91e391e5afe130eb78",
            "position": {
                "center": {
                    "x": 49.4,
                    "y": 37.6
                },
                "eye_left": {
                    "x": 43.3692,
                    "y": 30.8192
                },
                "eye_right": {
                    "x": 56.5606,
                    "y": 30.9886
                },
                "height": 26.8,
                "mouth_left": {
                    "x": 46.1326,
                    "y": 44.9468
                },
                "mouth_right": {
                    "x": 54.2592,
                    "y": 44.6282
                },
                "nose": {
                    "x": 49.9404,
                    "y": 38.8484
                },
                "width": 26.8
            },
            "tag": ""
        }
    ],
    "img_height": 500,
    "img_id": "22fd9efc64c87e00224c33dd8718eec7",
    "img_width": 500,
    "session_id": "38047ad0f0b34c7e8c6efb6ba39ed355",
    "url": "http://www.faceplusplus.com.cn/wp-content/themes/faceplusplus/assets/img/demo/1.jpg?v=4"
}
*/
