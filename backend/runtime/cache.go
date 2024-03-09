package runtime

import (
	"fmt"
	"grimlang/backend/asm"
)

type Cache struct {
	Items map[string]CacheItem
}

func NewCache() *Cache {
	return &Cache{
		Items: map[string]CacheItem{},
	}
}

func (ch *Cache) Set(key string, val CacheItem) {
	ch.Items[key] = val
}

func (ch *Cache) Get(key string) (CacheItem, error) {
	item, ok := ch.Items[key]
	if !ok {
		return CacheItem{}, fmt.Errorf("can't find item '%s'", key)
	}
	return item, nil
}

type CacheItem struct {
	Return []asm.Value
}
