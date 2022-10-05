package main

import "sync"

type Cache struct {
	cache map[string]*entry
	sync.Mutex
}

type entry struct {
	res   result
	ready chan struct{}
}

type result struct {
	value []byte
	err   error
}

type Func func() ([]byte, error)

func NewCache() *Cache {
	return &Cache{
		cache: make(map[string]*entry),
	}
}

func (c *Cache) Get(key string, f Func) ([]byte, error) {
	c.Lock()
	e := c.cache[key]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		c.cache[key] = e
		c.Unlock()

		e.res.value, e.res.err = f()
		close(e.ready)
	} else {
		<-e.ready
	}
}
