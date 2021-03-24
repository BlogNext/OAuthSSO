package oauth

import "github.com/gin-gonic/gin"


//创建预授权码
func CreatePreAuthCode(ctx *gin.Context){

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
