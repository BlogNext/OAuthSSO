package entity

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

//创建预授权码请求
type CreatePreAuthCodeRequest struct {
	User   *UserRequest   `json:"user"`
	Client *ClientRequest `json:"client"`
}

//创建预授权码
func (c CreatePreAuthCodeRequest) GetError(validationErrors validator.ValidationErrors) string {
	for _, fieldErr := range validationErrors {

		if strings.Contains(fieldErr.StructNamespace(), "Client") {
			return c.Client.GetError(validationErrors)
		}

		if strings.Contains(fieldErr.StructNamespace(), "User") {
			return c.User.GetError(validationErrors)
		}
	}

	return validationErrors.Error()
}

//创建预授权码响应
type CreatePreAuthCodeResponse struct {
	//预授权码
	PreAuthCode string `form:"pre_auth_code" json:"pre_auth_code"`
}
