package gocache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitCache(t *testing.T) {
	masterCache = nil
	nativeVersionCache = nil
	InitCache()
	assert.NotNil(t, masterCache)
	assert.NotNil(t, nativeVersionCache)
}

type CacheStruct struct {
	Field1 string `json:"field_1"`
	Field2 int64  `json:"field_2"`
	Field3 struct {
		Field1 bool `json:"field_1"`
	} `json:"field_3"`
}

func TestMasterCache(t *testing.T) {
	masterCache = nil
	initMasterCache()

	var ent1 *CacheStruct
	ok1 := GetMasterCache("key1", ent1)
	assert.False(t, ok1)
	assert.Nil(t, ent1)

	ent2 := CacheStruct{
		Field1: "field",
		Field2: 1,
		Field3: struct {
			Field1 bool `json:"field_1"`
		}{
			Field1: true,
		},
	}
	SetMasterCache("key2", &ent2)

	var ent3 CacheStruct
	ok2 := GetMasterCache("key2", &ent3)
	assert.True(t, ok2)
	assert.NotNil(t, ent3)
	assert.Equal(t, ent2, ent3)
}

func TestNativeVersionCache(t *testing.T) {
	nativeVersionCache = nil
	initNativeVersionCache()

	var ent1 *CacheStruct
	ok1 := GetNativeVersionCache("key1", ent1)
	assert.False(t, ok1)
	assert.Nil(t, ent1)

	ent2 := CacheStruct{
		Field1: "field",
		Field2: 1,
		Field3: struct {
			Field1 bool `json:"field_1"`
		}{
			Field1: true,
		},
	}
	SetNativeVersionCache("key2", &ent2)

	var ent3 CacheStruct
	ok2 := GetNativeVersionCache("key2", &ent3)
	assert.True(t, ok2)
	assert.NotNil(t, ent3)
	assert.Equal(t, ent2, ent3)
}
