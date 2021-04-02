package verify

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

//鉴权accessToken
type AccessToken struct {
	//鉴权的api地址
	url string
}

//鉴权
//return true有权，false无权
func(a *AccessToken) Verify(token string) bool{
	data := make(url.Values)
	data["access_token"] = []string{token}
	response ,err := http.PostForm(a.url,data)
	if err != nil {
		log.Println("请求鉴权服务器失败:",err)
		return false
	}
	defer response.Body.Close()

	result := make(map[string]interface{})
	jsonDecode := json.NewDecoder(response.Body)
	err = jsonDecode.Decode(result)
	if err != nil {
		log.Println("解码失败")
		return false
	}

	//是有有权限

	return true
}
