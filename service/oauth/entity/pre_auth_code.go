package entity

import "github.com/go-playground/validator/v10"

//预授权码的传输实体
type PreAuthCodeEntity struct {
	User   UserEntity
	Client ClientEntity
	//预授权码
	PreAuthCode string `form:"pre_auth_code" json:"pre_auth_code"`
}

func (pace PreAuthCodeEntity) GetError(validationErrors validator.ValidationErrors) string {
	errMsg := pace.User.GetError(validationErrors)
	if errMsg != "" {
		return errMsg
	}

	return validationErrors.Error()
}
