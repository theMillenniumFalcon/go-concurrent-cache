package main

import "sync"

type Cache struct {
	cache map[string]*result
	sync.Mutex
}

type result struct {
	value []byte
	err   error
}

type Func func() ([]byte, error)

func NewCache() *Cache {
	return &Cache{
		cache: make(map[string]*result),
	}
}

func (c *Cache) Get(key string, f Func) ([]byte, error) {
	c.Lock()
	defer c.Unlock()
	res, ok := c.cache[key]
	if !ok {
		res = &result{}
		res.value, res.err = f()
		c.cache[key] = res
	}

	return res.value, res.err
}
