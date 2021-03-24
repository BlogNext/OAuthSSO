package service

import "github.com/OauthSSO/service/oauth/entity"

//auth服务单例
var single *auth

func init() {
	if single == nil {
		single = new(auth)
	}
}

//获取单例
func GetAuthInstall() *auth {
	return single
}

type auth struct {
}

//创建预授权码
//request,请求
func (a *auth) CreatePreAuthCode(request *entity.CreatePreAuthCodeRequest) (response *entity.CreatePreAuthCodeResponse) {

	return nil
}
