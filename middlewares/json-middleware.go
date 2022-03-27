package middlewares

import "github.com/gin-gonic/gin"

// This middleware rewrite every response 'Content-Type' header.
func JsonMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Next()
	}
}
