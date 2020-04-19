package middleware

import "github.com/gin-gonic/gin"

// CORS 同源策略 中间件
func CORSMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.Writer.Header().Set("Access-Control-Max-Age", "86400")
		context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		context.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		context.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(200)
			return
		}
		context.Next()
	}
}
