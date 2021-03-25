package entity

import (
	"github.com/go-playground/validator/v10"
)

//创建预授权码请求
type CreatePreAuthCodeRequest struct {
	//用户的昵称
	Nickname string `form:"nickname" json:"nickname" binding:"required"`
	//密码
	Password string `form:"password" json:"password" binding:"required"`
	//客户id
	ClientId string `form:"client_id" json:"client_id" binding:"required"`
	//获取预授权码之后，重定向到原来的页面
	RedirectUrl string `form:"redirect_url" json:"redirect_url" binding:"required"`
}

//创建预授权码
func (c CreatePreAuthCodeRequest) GetError(validationErrors validator.ValidationErrors) string {

	for _, fieldErr := range validationErrors {

		if fieldErr.Field() == "Nickname" {
			switch fieldErr.Tag() {
			case "required":
				return "用户昵称nickname必填"
			}
		}

		if fieldErr.Field() == "Password" {
			switch fieldErr.Tag() {
			case "required":
				return "密码password必填"
			}
		}

		if fieldErr.Field() == "ClientId" {
			switch fieldErr.Tag() {
			case "required":
				return "client_id必填"
			}
		}

		if fieldErr.Field() == "RedirectUrl" {
			switch fieldErr.Tag() {
			case "required":
				return "redirect_url必填"
			}
		}
	}

	return validationErrors.Error()
}

//创建预授权码响应
type CreatePreAuthCodeResponse struct {
	//预授权码
	PreAuthCode string `form:"pre_auth_code" json:"pre_auth_code"`
	//用户授权客户之后，重定向地址到客户的服务器地址，让客户的后台通过预授权码获取accessToken
	RedirectUrl string `form:"redirect_url" json:"redirect_url"`
}
