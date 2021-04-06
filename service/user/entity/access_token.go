package entity

import "github.com/go-playground/validator/v10"

type AccessToken struct {
	AccessToken string `form:"access_token" json:"access_token"  binding:"required"`
}

type UserInfoRequest struct {
	AccessToken AccessToken
}

func (UserInfoRequest) GetError(validationErrors validator.ValidationErrors) string {
	for _, fieldErr := range validationErrors {
		if fieldErr.Field() == "AccessToken" {
			switch fieldErr.Tag() {
			case "required":
				return "access_token必填"
			}
		}

	}
	return validationErrors.Error()
}

type UserInfoResponse struct {
	Id       uint64 `json:"id"`
	Nickname string `json:"nickname"`
}
