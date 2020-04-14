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

// Response func 200 Success
func Success(context *gin.Context, data *gin.H, msg string) {
	Response(context, http.StatusOK, 200, data, msg)
}

// Response func 400 Fail
func Fail(context *gin.Context, data *gin.H, msg string) {
	Response(context, http.StatusOK, 400, data, msg)
}

// Response func 422 Unprocessable Entity
// 服务器理解请求实体的内容类型，并且请求实体的语法是正确的，但是服务器无法处理所包含的指令
func UnprocessableEntity(context *gin.Context, data *gin.H, msg string) {
	Response(context, http.StatusUnprocessableEntity, 422, data, msg)
}

// Response func 201 Created
// 代表成功的应答状态码，表示请求已经被成功处理，并且创建了新的资源。
func Created(context *gin.Context, data *gin.H, msg string) {
	Response(context, http.StatusCreated, 201, data, msg)
}

// Response func 500 ServerError
func ServerError(context *gin.Context, data *gin.H, msg string) {
	Response(context, http.StatusInternalServerError, 500, data, msg)
}
