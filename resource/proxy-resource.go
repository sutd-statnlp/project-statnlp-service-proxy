package resource

import (
	"../manager"
	"../util"
	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
	"github.com/rs/xid"
)

// ProxyResponse .
type ProxyResponse struct {
	Url    string
	ApiKey string
}

// InitProxyRoutes .
func InitProxyRoutes(router *gin.Engine) {
	router.POST("/api/proxy/set", func(context *gin.Context) {
		url := context.PostForm("url")
		if len(url) > 0 {
			apiKey := xid.New().String()
			body := ProxyResponse{Url: url, ApiKey: apiKey}
			manager.GetProxyManagerInstance().AddProxyUrl(apiKey, url)
			context.JSON(200, body)
		} else {
			context.JSON(400, util.InvalidFormData)
		}
	})

	router.GET("/api/proxy/get/:apiKey", func(context *gin.Context) {
		apiKey := context.Param("apiKey")
		if !manager.GetProxyManagerInstance().IsExisted(apiKey) {
			context.JSON(400, util.InvalidParam)
			return
		}
		url := manager.GetProxyManagerInstance().GetProxyUrl(apiKey)
		request := gorequest.New()
		_, body, errs := request.Get(url).End()
		if len(errs) > 0 {
			context.JSON(400, errs)
		} else {
			context.JSON(200, body)
		}
	})
}
