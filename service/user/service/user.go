package service

import (
	"github.com/OauthSSO/service/common/lazy/verify"
	"github.com/OauthSSO/service/user/entity"
)

func GetUserInfo(request *entity.UserInfoRequest) *entity.UserInfoResponse {
	template := verify.NewTemplate("http://127.0.0.1:8080/api/oauth/verify_access_token")
	_, err := template.GetResource(request.AccessToken.AccessToken, func() verify.Resource {
		return nil
	})
	if err != nil {
		return nil
	}

	response := new(entity.UserInfoResponse)
	response.Id = template.GetVerify().GetAccessTokenUserId()
	response.Nickname = "测试"
	return response
}
