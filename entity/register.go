package entity

// RegisterArg register arg
type RegisterArg struct {
	Code      string `form:"code"  json:"code"`
	ShareType int    `form:"shareType" json:"shareType,omitempty"`
	Scene     int    `form:"scene" json:"scene,omitempty"`
	Cover     int    `form:"cover" json:"cover,omitempty"`
}

// WXUserInfoArg wxGetUserInfo arg
type WXUserInfoArg struct {
	RawData       string `form:"rawData"  json:"rawData"`
	Signature     string `form:"signature"  json:"signature"`
	EncryptedData string `form:"encryptedData"  json:"encryptedData"`
	Iv            string `form:"iv"  json:"iv"`
}

// WxUserInfo 微信用户数据
type WxUserInfo struct {
	OpenId    string    `json:"openId"`
	NickName  string    `json:"nickName"`
	Gender    int32     `json:"gender"`
	City      string    `json:"city"`
	Province  string    `json:"province"`
	Country   string    `json:"country"`
	AvatarUrl string    `json:"avatarUrl"`
	UnionId   string    `json:"unionId"`
	Watermark Watermark `json:"watermark"`
}

// Watermark watermark
type Watermark struct {
	Appid     string `json:"appid"`
	Timestamp int64  `json:"timestamp"`
}
