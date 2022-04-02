package middlewares

import (
	"github.com/danyducky/social/app"
	"github.com/gin-gonic/gin"
)

// This middleware rewrite every response 'Content-Type' header.
type JsonMiddleware struct {
	ctx app.Context
}

// Creates json middleware instance.
func NewJsonMiddleware(ctx app.Context) JsonMiddleware {
	return JsonMiddleware{
		ctx: ctx,
	}
}

// Setup this middleware.
func (m JsonMiddleware) Setup() {
	middleware := func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Next()
	}

	m.ctx.ApiGroup.Use(middleware)
}
