package resource

import (
	"github.com/gin-gonic/gin"
)

// InitProxyRoutes .
func InitProxyRoutes(router *gin.Engine) {

	router.GET("/api/proxy", func(context *gin.Context) {
		body := "hello"
		context.JSON(200, body)
	})
}
