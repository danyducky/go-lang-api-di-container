package app

import "github.com/gin-gonic/gin"

// Application http context wrapper.
type Context struct {
	Gin      *gin.Engine
	ApiGroup *gin.RouterGroup
}

// Creates application http context instance.
func NewRequestHandler() Context {
	engine := gin.Default()

	apiGroup := engine.Group("/api")

	return Context{
		Gin:      engine,
		ApiGroup: apiGroup,
	}
}
