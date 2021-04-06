package service

import (
	"fmt"
	"github.com/OauthSSO/service/common/config"
	"github.com/OauthSSO/service/common/db"
	"github.com/OauthSSO/service/common/lazy/verify"
	"github.com/OauthSSO/service/model"
	"github.com/OauthSSO/service/user/entity"
)

var serviceInfo map[string]interface{}

func init() {
	if serviceInfo == nil {
		config.LoadConfig("server", "config", "yaml")
		serverConfig, _ := config.GetConfig("server")
		serviceInfo = serverConfig.GetStringMap("service")
	}
}

func GetUserInfo(request *entity.UserInfoRequest) *entity.UserInfoResponse {

	url := fmt.Sprintf("http://%s:%d/api/oauth/verify_access_token", serviceInfo["ip"].(string), serviceInfo["port"].(int))
	template := verify.NewTemplate(url)
	_, err := template.GetResource(request.AccessToken.AccessToken, func() verify.Resource {
		return nil
	})
	if err != nil {
		return nil
	}

	userModel := new(model.UserModel)
	userId := template.GetVerify().GetAccessTokenUserId()
	mysqlDB := db.GetDB()
	mysqlDB.Where("id = ?", userId).First(userModel)
	response := new(entity.UserInfoResponse)
	response.Id = userModel.ID
	response.Nickname = userModel.Nickname
	return response
}
