package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/OauthSSO/service/common/cache/lru"
	"github.com/OauthSSO/service/common/db"
	"github.com/OauthSSO/service/exception"
	"github.com/OauthSSO/service/oauth/entity"
	"github.com/OauthSSO/service/oauth/model"
	"strings"
	"time"
)

//auth服务单例
var single *auth

func init() {
	if single == nil {
		single = newAuth()
	}
}

//获取单例
func GetAuthInstall() *auth {
	return single
}

type auth struct {
	//缓存
	lruCache    *lru.LruCache
	OAuthClient OAuthClient
}

func newAuth() *auth {
	return &auth{lruCache: lru.NewLruCache(30)}
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

	//加密
	h := sha1.New()
	h.Write([]byte(request.Password))
	inputPassword := fmt.Sprintf("%x", h.Sum(nil))
	if strings.Compare(inputPassword, userModel.Password) != 0 {
		panic(exception.NewException(exception.ParamErr, "密码错误"))
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

//preAuthCode换出AccessToken
func (a *auth) PreAuthCodeAccessToken(request *entity.PreAuthCodeAccessTokenRequest) (response *entity.PreAuthCodeAccessTokenResponse) {
	return nil
}
