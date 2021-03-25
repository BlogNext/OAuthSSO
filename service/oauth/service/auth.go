package service

import (
	"fmt"
	"github.com/OauthSSO/service/common/cache/lru"
	"github.com/OauthSSO/service/common/db"
	"github.com/OauthSSO/service/exception"
	"github.com/OauthSSO/service/oauth/entity"
	"github.com/OauthSSO/service/oauth/model"
	"time"
)

//auth服务单例
var single *auth

func init() {
	if single == nil {
		single = new(auth)
		single.lruCache = lru.New(30)
	}
}

//获取单例
func GetAuthInstall() *auth {
	return single
}

type auth struct {
	//缓存
	lruCache *lru.LruCache
}

//创建预授权码
//request,请求
func (a *auth) CreatePreAuthCode(request *entity.CreatePreAuthCodeRequest) (response *entity.CreatePreAuthCodeResponse) {

	mysqlDB := db.GetDB()

	userModel := new(model.UserModel)
	if err := mysqlDB.Where("nickname = ?", request.Nickname).First(userModel).Error; err != nil {
		//没有数据
		panic(exception.NewException(exception.ParamErr, err.Error()))
	}

	//预授权码的格式： preAuthCode:clientId:userId
	preAuthCodeKey := fmt.Sprintf("preAuthCode:%s:%d", request.ClientId, userModel.ID)
	//添加到缓存,30秒缓存
	a.lruCache.Add(preAuthCodeKey, preAuthCodeKey, 30*time.Second)

	response = new(entity.CreatePreAuthCodeResponse)
	response.RedirectUrl = request.RedirectUrl
	response.PreAuthCode = preAuthCodeKey
	return response
}
