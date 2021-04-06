package main

import (
	"fmt"
	"github.com/OauthSSO/service/common/config"
	_ "github.com/OauthSSO/service/common/db"
	"github.com/OauthSSO/service/exception"
	"github.com/OauthSSO/service/help"
	"github.com/OauthSSO/service/oauth"
	"github.com/OauthSSO/service/user"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func init(){
	err := config.LoadConfig("server", "config", "yaml")
	if err != nil {
		log.Fatal("加载服务器配置失败",err)
	}
}

func main() {

	engine := gin.Default()

	//api路由
	api := engine.Group("/api")
	apiRouter(api)

	//web路由

	//启动web服务器
	serverConfig,_ := config.GetConfig("server")
	serviceInfo := serverConfig.GetStringMap("service")
	http.ListenAndServe(fmt.Sprintf("%s:%d",serviceInfo["ip"].(string),serviceInfo["port"].(int)),engine)
}

//api路由
func apiRouter(api *gin.RouterGroup) {
	//全局中间件注册
	//api统一异常处理中间件
	api.Use(func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch errLevel := err.(type) {
				case exception.MyException:
					//自定义的错误
					help.Gin200ErrorResponse(ctx, errLevel)
				default:
					//语言级别的，向上抛
					panic(errLevel)
				}
				//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
				ctx.Abort()
			}
		}()

		ctx.Next()
	})

	//路由
	api.GET("/ping", func(c *gin.Context) {
		panic(exception.NewException(1, "测试自定义异常"))
		//panic("系统级别")
	})

	//oauth授权功能模块
	oauthRouterGroup := api.Group("/oauth")
	{
		//预授权码办法token（有登录作用）
		oauthRouterGroup.POST("create_pre_auth_code", oauth.CreatePreAuthCode)
		//预授权码获取accessToken
		oauthRouterGroup.POST("pre_auth_code_access_token", oauth.PreAuthCodeAccessToken)
		//通过refresh_token刷新accessToken
		oauthRouterGroup.POST("refresh_token", oauth.RefreshToken)
		//验证accessToken，判断是否与权限访问资源
		oauthRouterGroup.POST("verify_access_token", oauth.VerifyAccessToken)
	}

	//资源模块
	resourceGroup := api.Group("/resource")
	{
		//用户资源
		userResourceGroup := resourceGroup.Group("/user")
		{
			//获取用户信息
			userResourceGroup.POST("/user_info",user.UserInfo)
		}
	}
}
