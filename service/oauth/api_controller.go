package oauth

import (
	"github.com/OauthSSO/service/exception"
	"github.com/OauthSSO/service/help"
	"github.com/OauthSSO/service/oauth/entity"
	"github.com/OauthSSO/service/oauth/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreatePreAuthCode(ctx *gin.Context) {
	var request entity.CreatePreAuthCodeRequest

	if err := ctx.ShouldBind(&request); err != nil {
		panic(exception.NewException(exception.ParamErr, request.GetError(err.(validator.ValidationErrors))))
	}

	//创建预授权码
	auth := service.GetAuthInstall()
	response := auth.CreatePreAuthCode(&request)

	help.Gin200SuccessResponse(ctx, "成功", response)

	return
}

//预授权码换取token模式
func PreAuthCodeAccessToken(ctx *gin.Context) {
	var request entity.PreAuthCodeAccessTokenRequest

	if err := ctx.ShouldBind(&request); err != nil {
		panic(exception.NewException(exception.ParamErr, request.GetError(err.(validator.ValidationErrors))))
	}

	//预授权码换取AccessToken
	auth := service.GetAuthInstall()
	response := auth.PreAuthCodeAccessToken(&request)

	help.Gin200SuccessResponse(ctx, "成功", response)

	return
}

//通过refreshToken刷新token
func RefreshToken(ctx *gin.Context) {

}

//验证AccessToken，看AccessToken是否有权限访问该资源
func VerifyAccessToken(ctx *gin.Context) {

}
