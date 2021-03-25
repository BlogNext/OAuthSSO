package entity

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
)

//预授权码换取accessToken的请求
type PreAuthCodeAccessTokenRequest struct {
	//客户id
	ClientId     string `form:"client_id" json:"client_id" binding:"required"`
	ClientSecret string `form:"client_secret" json:"client_secret" binding:"required"`
	//预授权码
	PreAuthCode string `form:"pre_auth_code" json:"pre_auth_code"`
}

//创建预授权码
func (c PreAuthCodeAccessTokenRequest) GetError(validationErrors validator.ValidationErrors) string {

	for _, fieldErr := range validationErrors {

		if fieldErr.Field() == "ClientId" {
			switch fieldErr.Tag() {
			case "required":
				return "client_id必填"
			}
		}

		if fieldErr.Field() == "ClientSecret" {
			switch fieldErr.Tag() {
			case "required":
				return "client_secret必填"
			}
		}

		if fieldErr.Field() == "PreAuthCode" {
			switch fieldErr.Tag() {
			case "required":
				return "pre_auth_code必填"
			}
		}
	}

	return validationErrors.Error()
}

//预授权码换取accessToken的响应
type PreAuthCodeAccessTokenResponse struct {
	//预授权码
	//给客户的前端用的
	AccessToken string `form:"access_token" json:"access_token"`
	//给客户的后端用的，不能暴露
	RefreshToken string `form:"refresh_token" json:"refresh_token"`
}

//accessToken
type AccessTokenJwt struct {
	jwt.StandardClaims
	//预授权码
	ClientId string `json:"client_id"`
	UserId   uint64 `json:"user_id"`
}

//refresh_token
type RefreshTokenJwt struct {
	jwt.StandardClaims
	//预授权码
	ClientId string `json:"client_id"`
	UserId   uint64 `json:"user_id"`
}
