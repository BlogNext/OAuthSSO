package entity

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

//预授权码的传输实体
type PreAuthCodeEntity struct {
	User   UserEntity
	Client ClientEntity
	//预授权码
	PreAuthCode string `form:"pre_auth_code" json:"pre_auth_code"`
}

func (pace PreAuthCodeEntity) GetError(validationErrors validator.ValidationErrors) string {
	for _, fieldErr := range validationErrors {
		
		if strings.Contains(fieldErr.StructNamespace(), "Client") {
			return pace.Client.GetError(validationErrors)
		}

		if strings.Contains(fieldErr.StructNamespace(), "User") {
			return pace.User.GetError(validationErrors)
		}
	}

	return validationErrors.Error()
}
