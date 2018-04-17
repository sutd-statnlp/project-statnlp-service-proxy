package manager

import (
	"sync"

	"../util"
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

// AddApiKey .
func (apiKeyManager ProxyManager) AddApiKey(apiKey string) {
	apiKeyManager.ApiKeyMap[apiKey] = util.GetCurrentTimestamp()
}

// GetTimestamp .
func (apiKeyManager ProxyManager) GetTimestamp(apiKey string) string {
	return apiKeyManager.ApiKeyMap[apiKey]
}

// IsExisted .
func (apiKeyManager ProxyManager) IsExisted(apiKey string) bool {
	timestamp := apiKeyManager.GetTimestamp(apiKey)
	return len(timestamp) > 0
}
