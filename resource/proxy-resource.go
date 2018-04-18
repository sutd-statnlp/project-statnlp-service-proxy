package resource

import (
	"../service"
	"../util"
	"github.com/gin-gonic/gin"
)

// InitProxyRoutes .
func InitProxyRoutes(router *gin.Engine) {
	router.POST("/api/proxy/set", func(context *gin.Context) {
		url := context.PostForm("url")
		if len(url) > 0 {
			body := service.SetUrl(url)
			context.JSON(200, body)
		} else {
			context.JSON(400, util.InvalidFormData)
		}
	})

	router.GET("/api/proxy/get/:apiKey", func(context *gin.Context) {
		apiKey := context.Param("apiKey")
		body, err := service.GetUrl(apiKey)
		if err != nil {
			context.JSON(400, err)
		} else {
			context.JSON(200, body)
		}
	})
}
