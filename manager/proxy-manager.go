package manager

import (
	"sync"
)

// ProxyManager .
type ProxyManager struct {
	ApiKeyMap map[string]string
}

var instance *ProxyManager
var once sync.Once

// GetProxyManagerInstance .
func GetProxyManagerInstance() *ProxyManager {
	once.Do(func() {
		instance = &ProxyManager{}
		instance.ApiKeyMap = make(map[string]string)
	})
	return instance
}

// AddProxyUrl .
func (proxyManager ProxyManager) AddProxyUrl(apiKey string, url string) {
	proxyManager.ApiKeyMap[apiKey] = url
}

// GetProxyUrl .
func (proxyManager ProxyManager) GetProxyUrl(apiKey string) string {
	return proxyManager.ApiKeyMap[apiKey]
}

// IsExisted .
func (proxyManager ProxyManager) IsExisted(apiKey string) bool {
	url := proxyManager.GetProxyUrl(apiKey)
	return len(url) > 0
}

// DeleteApiKey .
func (proxyManager ProxyManager) DeleteApiKey(apiKey string) bool {
	delete(proxyManager.ApiKeyMap, apiKey)
	return true
}
