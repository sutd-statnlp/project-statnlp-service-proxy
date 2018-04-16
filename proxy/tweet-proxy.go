package resource

import (
	"github.com/gin-gonic/gin"
)

// InitTweetProxyRoutes .
func InitTweetProxyRoutes(router *gin.Engine) {

	router.GET("/proxy/api/tweets/crawl", func(context *gin.Context) {
		body := "hello"
		context.JSON(200, body)
	})
}
