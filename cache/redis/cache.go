package redis

import (
	"encoding/json"
	"sync"
	"time"
)

var mu sync.RWMutex

func GetCache(cacheKey string) ([]byte, error) {
	mu.RLock()
	defer mu.RUnlock()
	data, err := Get(cacheKey)
	if err != nil {
		return nil, err
	}
	return []byte(data), nil
}

func SetCache(cacheKey string, value interface{}, expiration time.Duration) error {
	mu.RLock()
	oldData, _ := GetCache(cacheKey)
	mu.RUnlock()

	if oldData != nil {
		mu.Lock()
		Set("old"+cacheKey, oldData, expiration)
		defer Delete("old" + cacheKey)
		mu.Unlock()
	}

	mu.Lock()
	defer mu.Unlock()
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return Set(cacheKey, data, expiration)
}
