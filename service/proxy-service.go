package service

import (
	"errors"

	"../manager"
	"../model"
	"../util"
	"github.com/parnurzeal/gorequest"
	"github.com/rs/xid"
)

// GetUrl .
func GetUrl(apiKey string) (string, error) {
	if manager.GetProxyManagerInstance().IsExisted(apiKey) {
		url := manager.GetProxyManagerInstance().GetProxyUrl(apiKey)
		request := gorequest.New()
		_, content, errs := request.Get(url).End()
		return content, errs[0]
	}
	return "", errors.New(util.InvalidParam)
}

// SetUrl .
func SetUrl(url string) model.ProxyResponse {
	apiKey := xid.New().String()
	manager.GetProxyManagerInstance().AddProxyUrl(apiKey, url)
	return model.ProxyResponse{Url: url, ApiKey: apiKey}
}
