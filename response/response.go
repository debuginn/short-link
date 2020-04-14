package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response func
func Response(context *gin.Context, httpStatus int, code int, data *gin.H, msg string) {
	context.JSON(httpStatus, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

// Response func Success
func Success(context *gin.Context, data *gin.H, msg string) {
	Response(context, http.StatusOK, 200, data, msg)
}

// Response func Fail
func Fail(context *gin.Context, data *gin.H, msg string) {
	Response(context, http.StatusOK, 400, data, msg)
}
