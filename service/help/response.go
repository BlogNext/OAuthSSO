package help

import (
	"github.com/OauthSSO/service/exception"
	"github.com/gin-gonic/gin"
	"net/http"
)

//200返回.请求正常
func Gin200SuccessResponse(c *gin.Context, msg string, data interface{}) {
	gin200Response(c, 0, msg, data)
}

//200返回,业务异常,错误的返回
//c gin的上下文
// code 自定义异常
func Gin200ErrorResponse(c *gin.Context, err exception.MyException) {
	gin200Response(c, err.GetErrorCode(), err.Error(), nil)
}

//200返回
func gin200Response(c *gin.Context, code int, msg string, data interface{}) {
	h := gin.H{"code": code, "msg": msg}
	if data != nil {
		h["data"] = data
	}
	c.JSON(http.StatusOK, h)
}
