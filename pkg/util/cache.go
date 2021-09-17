package util

// 简单缓存 适合小数据量

import (
	"sync"
	"sync/atomic"
	"time"
)

var globalMap sync.Map
var cacheLen int64

func SetCache(key string, data interface{}, timeout int) {
	globalMap.Store(key, data)
	atomic.AddInt64(&cacheLen, 1)
	time.AfterFunc(time.Second*time.Duration(timeout), func() {
		atomic.AddInt64(&cacheLen, -1)
		globalMap.Delete(key)
	})
}

func GetCache(key string) (interface{}, bool) {
	return globalMap.Load(key)
}

func DeleteCache(key string) {
	globalMap.Delete(key)
}

//func LenCache() int {
//	return int(cacheLen)
//}
