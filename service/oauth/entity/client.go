package entity

import (
	"github.com/go-playground/validator/v10"
)

type ClientRequest struct {
	//客户id
	ClientId string `form:"client_id" json:"client_id" binding:"required"`
	//用户授权客户之后，重定向地址到客户的服务器地址，让客户的后台通过预授权码获取accessToken
	RedirectUrl string `form:"redirect_url" json:"redirect_url"`
}

//自定义错误信息
func (u ClientRequest) GetError(validationErrors validator.ValidationErrors) string {
	for _, fieldErr := range validationErrors {

		if fieldErr.Field() == "ClientId" {
			switch fieldErr.Tag() {
			case "required":
				return "client_id必填"
			}
		}

	}

	return validationErrors.Error()
}
