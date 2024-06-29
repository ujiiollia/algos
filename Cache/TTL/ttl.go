package main

import (
	"fmt"
	"sync"
	"time"
)

// CacheItem представляет элемент в кэше.
type CacheItem struct {
	Value      interface{}
	Expiration int64
}

// TTLCache представляет кэш с поддержкой TTL.
type TTLCache struct {
	items    map[string]*CacheItem
	mutex    sync.Mutex
	ttl      time.Duration
	gcTicker *time.Ticker
	stopGcCh chan bool
}

// NewTTLCache создает новый TTL кэш с заданным временем жизни элементов.
func NewTTLCache(ttl time.Duration, gcInterval time.Duration) *TTLCache {
	cache := &TTLCache{
		items:    make(map[string]*CacheItem),
		ttl:      ttl,
		gcTicker: time.NewTicker(gcInterval),
		stopGcCh: make(chan bool),
	}

	// Запуск сборщика мусора в отдельной горутине
	go cache.gc()

	return cache
}

// Set добавляет элемент в кэш.
func (c *TTLCache) Set(key string, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	expiration := time.Now().Add(c.ttl).UnixNano()
	c.items[key] = &CacheItem{
		Value:      value,
		Expiration: expiration,
	}
}

// Get возвращает значение из кэша по ключу.
func (c *TTLCache) Get(key string) (interface{}, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	item, found := c.items[key]
	if !found {
		return nil, false
	}

	if time.Now().UnixNano() > item.Expiration {
		// Удаляем просроченный элемент из кэша
		delete(c.items, key)
		return nil, false
	}
	return item.Value, true
}

// Delete удаляет элемент из кэша.
func (c *TTLCache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.items, key)
}

// gc выполняет сборку мусора каждые gcInterval.
func (c *TTLCache) gc() {
	for {
		select {
		case <-c.gcTicker.C:
			c.removeExpiredItems()
		case <-c.stopGcCh:
			c.gcTicker.Stop()
			return
		}
	}
}

// removeExpiredItems удаляет просроченные элементы из кэша.
func (c *TTLCache) removeExpiredItems() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	now := time.Now().UnixNano()
	for key, item := range c.items {
		if now > item.Expiration {
			delete(c.items, key)
		}
	}
}

// Close останавливает сборщик мусора.
func (c *TTLCache) Close() {
	c.stopGcCh <- true
}

func main() {
	cache := NewTTLCache(5*time.Second, 1*time.Second)
	defer cache.Close()

	cache.Set("a", 1)
	cache.Set("b", 2)

	time.Sleep(3 * time.Second)
	v, ok := cache.Get("a")
	if ok {
		fmt.Println("a:", v)
	} else {
		fmt.Println("a: expired")
	}

	time.Sleep(3 * time.Second)
	v, ok = cache.Get("b")
	if ok {
		fmt.Println("b:", v)
	} else {
		fmt.Println("b: expired")
	}
}
