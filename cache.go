package main

import (
	"time"
	"errors"
)

type cache struct {
	temple map[string]interface{}
}

func NewCache() *cache {
	return &cache{
		temple: make(map[string]interface{}),
	}
}

func (c *cache) Set(key string, value interface{}, ttl time.Duration) {
	c.temple[key] = value
	
	go func() {
		time.Sleep(ttl)
		c.Delete(key)
	}()
}

func (c *cache) Get(key string) (interface{}, error) {
	var err error
	cacheValue := c.temple[key]
	if cacheValue == nil {
		err = errors.New("the key does not exist")
	}
	return cacheValue, err
}

func (c *cache) Delete(key string) {
	delete(c.temple, key)
}
