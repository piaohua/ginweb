package libs

import (
	"net/http"

	"github.com/json-iterator/go"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

// GetJson 发送GET请求解析json
func GetJson(uri string, v interface{}) error {
	r, err := http.Get(uri)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}
