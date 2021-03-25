package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/OauthSSO/service/common/cache/lru"
	"github.com/OauthSSO/service/common/db"
	"github.com/OauthSSO/service/exception"
	"github.com/OauthSSO/service/oauth/entity"
	"github.com/OauthSSO/service/oauth/model"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

//auth服务单例
var single *auth

//jwt签名key
var jwtSigningKey []byte

func init() {
	if single == nil {
		single = newAuth()
	}

	if jwtSigningKey == nil {
		jwtSigningKey = []byte("blog_oauth")
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

	//验证授权用户的账号密码是否正确
	mysqlDB := db.GetDB()
	userModel := new(model.UserModel)
	if err := mysqlDB.Where("nickname = ?", request.Nickname).First(userModel).Error; err != nil {
		//没有数据
		panic(exception.NewException(exception.ParamErr, err.Error()))
	}

	h := sha1.New()
	h.Write([]byte(request.Password))
	inputPassword := fmt.Sprintf("%x", h.Sum(nil))
	if strings.Compare(inputPassword, userModel.Password) != 0 {
		panic(exception.NewException(exception.ParamErr, "密码错误"))
	}

	//预授权码的格式： preAuthCode:clientId:userId
	preAuthCodeKey := fmt.Sprintf("preAuthCode:%s:%d", request.ClientId, userModel.ID)
	//到期时间
	expires := 30 * time.Second
	//添加到缓存,30秒缓存
	a.lruCache.Add(preAuthCodeKey, preAuthCodeKey, expires)

	//预授权码用jwt生成
	preAuthCodeJwtClaims := &entity.PreAuthCodeJwt{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "ly",
			ExpiresAt: time.Now().Add(expires).Unix(),
		},
		ClientId: request.ClientId,
		UserId:   userModel.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, preAuthCodeJwtClaims)
	preAuthCode, _ := token.SignedString(jwtSigningKey)

	//数据返回
	response = new(entity.CreatePreAuthCodeResponse)
	response.RedirectUrl = request.RedirectUrl
	response.PreAuthCode = preAuthCode
	return response
}

//preAuthCode换出AccessToken
func (a *auth) PreAuthCodeAccessToken(request *entity.PreAuthCodeAccessTokenRequest) (response *entity.PreAuthCodeAccessTokenResponse) {

	//验证预授权码是否过期
	preAuthCodeJwtClaims := &entity.PreAuthCodeJwt{}
	token, _ := jwt.ParseWithClaims(request.PreAuthCode, preAuthCodeJwtClaims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSigningKey, nil
	})

	if !token.Valid {
		panic(exception.NewException(exception.ParamErr, "pre_auth_code已经失效"))
	}

	//验证授权的clientId账号是否正确
	mysqlDB := db.GetDB()
	oauthClientModel := new(model.OAuthClientModel)
	mysqlDB = mysqlDB.Where("client_id = ?", request.ClientId).Where("client_secret = ?", request.ClientSecret).First(oauthClientModel)
	if mysqlDB.Error != nil {
		//没有数据
		panic(exception.NewException(exception.ParamErr, "client信息不正确"))
	}

	response = new(entity.PreAuthCodeAccessTokenResponse)
	response.AccessToken = "123"
	response.RefreshToken = "123111"

	return response
}
