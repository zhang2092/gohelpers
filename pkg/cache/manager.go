package cache

// Manager 管理缓存的接口定义
type Manager interface {
	// Set 设置缓存
	Set(key string, value interface{}) error
	// SetDefault 设置缓存
	SetDefault(key string, value interface{}, expire int64) error
	// Get 获取缓存
	Get(key string) (interface{}, error)
	// GetString 获取缓存
	GetString(key string) (string, error)
	// Exists key是否存在
	Exists(key string) bool
	// Delete 删除key
	Delete(key string) error
}
