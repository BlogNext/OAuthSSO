package entity

import (
	"github.com/go-playground/validator/v10"
)

//用户的传输实体
type UserEntity struct {
	//用户的昵称
	Nickname string `form:"nickname" json:"nickname" binding:"required"`
	//密码
	Password string `form:"password" json:"password,omitempty" binding:"required"`
}

//自定义错误信息
func (u UserEntity) GetError(validationErrors validator.ValidationErrors) string {

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
	}

	return validationErrors.Error()
}
