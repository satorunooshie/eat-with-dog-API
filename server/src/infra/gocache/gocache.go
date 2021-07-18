package gocache

import (
	"encoding/gob"
	"reflect"
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	masterCache *cache.Cache
	nativeVersionCache *cache.Cache
)

func init() {
	gob.Register(map[string]interface{}{})
	gob.Register(map[string]interface{}{})
	gob.Register(map[int64]float64{})
}

func InitCache() {
	initMasterCache()
	initNativeVersionCache()
}

func initMasterCache() {
	ttl := 1 * time.Hour
	masterCache = cache.New(ttl, ttl+3*time.Minute)
}

func initNativeVersionCache() {
	ttl := 10 * time.Minute
	nativeVersionCache = cache.New(ttl, ttl+3*time.Minute)
}

func GetMasterCache(key string, ptr interface{}) bool {
	if key == "" {
		return false
	}
	if masterCache == nil {
		return false
	}
	v, ok := masterCache.Get(key)
	if !ok {
		return false
	}
	return copyValue(ptr, v)
}

func GetNativeVersionCache(key string, ptr interface{}) bool {
	if key == "" {
		return false
	}
	if nativeVersionCache == nil {
		return false
	}
	v, ok := nativeVersionCache.Get(key)
	if !ok {
		return false
	}
	return copyValue(ptr, v)
}

func SetMasterCache(key string, value interface{}) {
	if key == "" {
		return
	}
	if masterCache == nil {
		return
	}
	masterCache.Set(key, value, cache.DefaultExpiration)
}

func SetNativeVersionCache(key string, value interface{}) {
	if key == "" {
		return
	}
	if nativeVersionCache == nil {
		return
	}
	nativeVersionCache.Set(key, value, cache.DefaultExpiration)
}

func copyValue(dst, src interface{}) bool {
	vvDst := reflect.ValueOf(dst)
	switch {
	case vvDst.Kind() != reflect.Ptr:
		// cannot assign value for non-pointer
		return false
	case vvDst.IsNil():
		return false
	}

	vvDst = vvDst.Elem()
	if !vvDst.CanSet() {
		return false
	}

	vvSrc := reflect.ValueOf(src)
	if vvSrc.Kind() == reflect.Ptr {
		vvSrc = vvSrc.Elem()
	}

	// type check
	switch {
	case vvSrc.Kind() != vvDst.Kind():
		return false
	case vvSrc.Type() != vvDst.Type():
		return false
	}

	vvDst.Set(vvSrc)
	return true
}
