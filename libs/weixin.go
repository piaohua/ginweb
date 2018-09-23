package libs

import (
	"fmt"
)

// OAUTH2PAGE oauth2鉴权
const (
	Code2PAGE = "https://api.weixin.qq.com/sns/jscode2session?appid=%v&secret=%v&js_code=%v&grant_type=authorization_code"
)

// Jscode2Session code换session
func Jscode2Session(code, appid, secret string) (wxs *WxSession, err error) {
	url := fmt.Sprintf(Code2PAGE, appid, secret, code)
	wxs = new(WxSession)
	err = GetJson(url, wxs)
	if wxs.Error() != nil {
		err = wxs.Error()
	}
	return
}

//WxSession 登录获取微信JSON数据
type WxSession struct {
	WxErr
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
}

//WxErr 通用错误
type WxErr struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (w *WxErr) Error() error {
	if w.ErrCode != 0 {
		return fmt.Errorf("err: errcode=%v , errmsg=%v", w.ErrCode, w.ErrMsg)
	}
	return nil
}

//ParseUserInfo 解析微信用户数据
func ParseUserInfo(b []byte, wxUserInfo interface{}) error {
	return json.Unmarshal(b, wxUserInfo)
}