package main

import (
	"github.com/OauthSSO/service/oauth"
	"github.com/gin-gonic/gin"
)


func main(){
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//oauth路由组
	oauthRouterGroup := r.Group("/oauth")
	{
		//创建预授权码
		oauthRouterGroup.POST("create_pre_auth_code",oauth.CreatePreAuthCode)
		//预授权码获取accessToken
		oauthRouterGroup.POST("pre_auth_code_access_token",oauth.PreAuthCodeAccessToken)
		//通过refresh_token刷新accessToken
		oauthRouterGroup.POST("refresh_token",oauth.RefreshToken)
		//验证accessToken，判断是否与权限访问资源
		oauthRouterGroup.POST("verify_access_token",oauth.VerifyAccessToken)
	}


	r.Run()
}
