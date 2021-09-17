package cache

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

// RedisManager redis缓存管理
type RedisManager struct {
	pool *redis.Pool
}

// NewRedisManager 获取一个新的redis管理对象
func NewRedisManager(address string) Manager {
	pool := &redis.Pool{
		// 连接方法
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", address)
			if err != nil {
				return nil, err
			}
			c.Do("SELECT", 0)
			return c, nil
		},
		//DialContext:     nil,
		//TestOnBorrow:    nil,
		MaxIdle:     10,                // 最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态。
		MaxActive:   10,                // 最大的激活连接数，表示同时最多有N个连接
		IdleTimeout: 360 * time.Second, // 最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
		//Wait:            false,
		//MaxConnLifetime: 0,
	}
	return &RedisManager{pool}
}

// Set 设置缓存
// key 键
// value 值
func (manager *RedisManager) Set(key string, value interface{}) error {
	conn := manager.pool.Get()
	defer conn.Close()

	_, err := conn.Do("Set", key, value)
	if err != nil {
		return fmt.Errorf("set cache error: %v", err)
	}
	return nil
}

// SetDefault 设置缓存
// key 键
// value 值
// expire 过期时间 (秒)
func (manager *RedisManager) SetDefault(key string, value interface{}, expire int64) error {
	conn := manager.pool.Get()
	defer conn.Close()

	//_, err := conn.Do("set", key, value)
	//if err != nil {
	//	return fmt.Errorf("set default cache error: %v", err)
	//}
	//
	//// 设置过期时间
	//conn.Do("expire", key, expire)
	//return nil

	_, err := conn.Do("set", key, value, "ex", expire)
	if err != nil {
		return fmt.Errorf("set default cache error: %v", err)
	}
	return nil
}

// Get 获取缓存
func (manager *RedisManager) Get(key string) (interface{}, error) {
	conn := manager.pool.Get()
	defer conn.Close()

	// 检查key是否存在
	exit, err := redis.Bool(conn.Do("exists", key))
	if err != nil || !exit {
		return nil, fmt.Errorf("key is not exists")
	}

	value, err := conn.Do("get", key)
	if err != nil {
		return nil, fmt.Errorf("get cache error: %v", err)
	}
	return value, nil
}

// GetString 获取缓存
func (manager *RedisManager) GetString(key string) (string, error) {
	conn := manager.pool.Get()
	defer conn.Close()

	// 检查key是否存在
	exit, err := redis.Bool(conn.Do("exists", key))
	if err != nil || !exit {
		return "", fmt.Errorf("key is not exists")
	}

	value, err := redis.String(conn.Do("get", key))
	if err != nil {
		return "", fmt.Errorf("get cache error: %v", err)
	}
	return value, nil
}

// Exists key是否存在
func (manager *RedisManager) Exists(key string) bool {
	conn := manager.pool.Get()
	defer conn.Close()

	// 检查key是否存在
	exit, err := redis.Bool(conn.Do("expire", key))
	if err != nil {
		return false
	}
	return exit
}

// Delete 删除key
func (manager *RedisManager) Delete(key string) error {
	conn := manager.pool.Get()
	defer conn.Close()

	_, err := conn.Do("del", key)
	if err != nil {
		return fmt.Errorf("删除key：[%s]异常:%v", key, err)
	}
	return nil
}
