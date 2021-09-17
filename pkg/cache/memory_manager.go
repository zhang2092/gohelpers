package cache

import (
	"errors"
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

// MemoryManager memory缓存管理
type MemoryManager struct {
	cache *cache.Cache
}

// NewMemoryManager 获取一个新的内存管理对象
// defaultExpiration 缓存过期时间
// cleanupInterval 缓存清理时间
func NewMemoryManager(defaultExpiration, cleanupInterval time.Duration) Manager {
	return &MemoryManager{cache: cache.New(defaultExpiration, cleanupInterval)}
}

// Set 设置缓存
// key 键
// value 值
func (manager *MemoryManager) Set(key string, value interface{}) error {
	manager.cache.Set(key, value, time.Hour*24*365*100) // 10年
	return nil
}

// SetDefault 设置缓存
// key 键
// value 值
// expire 过期时间 (秒)
func (manager *MemoryManager) SetDefault(key string, value interface{}, expire int64) error {
	manager.cache.Set(key, value, time.Second*time.Duration(expire))
	return nil
}

// Get 获取缓存
func (manager *MemoryManager) Get(key string) (interface{}, error) {
	value, found := manager.cache.Get(key)
	if found {
		return value, nil
	}
	return nil, errors.New("no found data with key")
}

// GetString 获取缓存
func (manager *MemoryManager) GetString(key string) (string, error) {
	value, found := manager.cache.Get(key)
	if found {
		return fmt.Sprintf("%v", value), nil
	}
	return "", errors.New("no found data with key")
}

// Exists key是否存在
func (manager *MemoryManager) Exists(key string) bool {
	if _, found := manager.cache.Get(key); found {
		return true
	}
	return false
}

// Delete 删除key
func (manager *MemoryManager) Delete(key string) error {
	manager.cache.Delete(key)
	return nil
}
