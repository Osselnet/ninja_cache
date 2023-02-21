package main

import (
	"errors"
	"sync"
	"time"
)

type cache struct {
	temple map[string]interface{}
	mx sync.RWMutex
}

func NewCache() *cache {
	return &cache{
		temple: make(map[string]interface{}),
	}
}

func (c *cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mx.Lock()
	c.temple[key] = value
	c.mx.Unlock()

	go func() {
		time.Sleep(ttl)
		c.Delete(key)
	}()
}

func (c *cache) Get(key string) (cacheValue interface{}, err error) {
	c.mx.RLock()
    defer c.mx.RUnlock()

	cacheValue = c.temple[key]
	if cacheValue == nil {
		err = errors.New("the key does not exist")
	}
	return cacheValue, err
}

func (c *cache) Delete(key string) {
	c.mx.Lock()
	defer c.mx.Unlock()

	delete(c.temple, key)
}
