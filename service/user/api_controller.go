package user

import (
	"github.com/OauthSSO/service/exception"
	"github.com/OauthSSO/service/help"
	"github.com/OauthSSO/service/user/entity"
	"github.com/OauthSSO/service/user/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

//获取用户信息
func UserInfo(ctx *gin.Context){

	var request entity.UserInfoRequest

	if err := ctx.ShouldBind(&request); err != nil {
		panic(exception.NewException(exception.ParamErr, request.GetError(err.(validator.ValidationErrors))))
	}

	response := service.GetUserInfo(&request)
	help.Gin200SuccessResponse(ctx, "成功", response)
	return
}