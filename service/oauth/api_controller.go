package oauth

import (
	"github.com/OauthSSO/service/exception"
	"github.com/OauthSSO/service/help"
	"github.com/OauthSSO/service/oauth/entity"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreatePreAuthCode(ctx *gin.Context) {
	var perAuthCode entity.PreAuthCodeEntity
	if err := ctx.ShouldBind(&perAuthCode); err != nil {
		panic(exception.NewException(exception.ParamErr, perAuthCode.GetError(err.(validator.ValidationErrors))))
	}

	help.Gin200SuccessResponse(ctx, "成功", perAuthCode)
	return
}

//预授权码换取token模式
func PreAuthCodeAccessToken(ctx *gin.Context) {
}

//通过refreshToken刷新token
func RefreshToken(ctx *gin.Context) {

}

//验证AccessToken，看AccessToken是否有权限访问该资源
func VerifyAccessToken(ctx *gin.Context) {

}
