package middleware

import (
	"net/http"
	"strings"

	"github.com/debuginn/GoShortLink/common"
	"github.com/debuginn/GoShortLink/model"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware 用户校验中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 获取 authorization header
		tokenString := context.GetHeader("Authorization")

		// 验证 token 是否有效
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			context.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			context.Abort()
			return
		}
		// 获取通过认证后的 userId
		userId := claims.UserId
		user, err := model.GetAUserById(userId)

		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			context.Abort()
			return
		}
		// 用户不存在
		if user.ID == 0 {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			context.Abort()
			return
		}
		// 用户存在，将 user 写入上下文
		context.Set("user", user)
		context.Next()
	}
}
