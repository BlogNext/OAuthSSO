package verify

import (
	"encoding/json"
	"github.com/OauthSSO/service/oauth/entity"
	"io"
	"log"
	"net/http"
	"net/url"
)

//响应
type Response struct {
	Code int      `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg"`
}

//验证
type Verify interface {
	Verify(token string) bool
	//获取accessToken的userId
	GetAccessTokenUserId() uint64
}

//鉴权accessToken
type VerifyAccessToken struct {
	//鉴权的api地址
	url               string
	accessTokenUserId uint64
}

//获取accessToken的UserId
func (a *VerifyAccessToken) GetAccessTokenUserId() uint64 {
	return a.accessTokenUserId
}

//鉴权
//return true有权，false无权
func (a *VerifyAccessToken) Verify(token string) bool {
	data := make(url.Values)
	data["access_token"] = []string{token}
	response, err := http.PostForm(a.url, data)
	if err != nil {
		log.Println("请求鉴权服务器失败:", err)
		return false
	}

	defer response.Body.Close()
	body ,err :=io.ReadAll(response.Body)
	if err != nil {
		log.Println("读取数据失败")
		return false
	}

	jsonResponse := new(Response)
	verifyAccessTokenResponse := new(entity.VerifyAccessTokenResponse)
	jsonResponse.Data = verifyAccessTokenResponse
	err = json.Unmarshal(body,jsonResponse)
	if err != nil {
		log.Println("解码失败")
		return false
	}

	if jsonResponse.Data == nil {
		log.Println("认证失败，没有data")
		return false
	}

	if jsonResponse.Data.(*entity.VerifyAccessTokenResponse).IsPower == true {
		//有权执行
		a.accessTokenUserId = jsonResponse.Data.(*entity.VerifyAccessTokenResponse).UserId
		return true
	}

	//没有权限
	return false
}
