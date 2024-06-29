package main

import (
	"errors"
	"fmt"
)

type dll struct {
	next  *dll
	prev  *dll
	key   int
	value int
}

type lruCache struct {
	cap   int
	cache map[int]*dll
	head  *dll
	tail  *dll
}

func newLRUCache(cap int) lruCache {
	lru := lruCache{
		cap:   cap,
		cache: make(map[int]*dll, cap),
		head:  &dll{},
		tail:  &dll{},
	}
	lru.head.next = lru.tail
	lru.head.prev = lru.head
	return lru
}

func (current *lruCache) getLRU(key int) (int, error) {
	if dllV, ok := current.cache[key]; ok {

		dllV.prev.next = dllV.next
		dllV.next.prev = dllV.prev

		dllV.next = current.head.next
		dllV.prev = current.tail.prev

		current.tail.next = dllV
		current.head.next.prev = dllV
		return dllV.value, nil
	}

	return 0, errors.New("no such key")
}
func (current *lruCache) setLRU(key, value int) error {
	if dllV, ok := current.cache[key]; ok {
		dllV.value = value
		_, err := current.getLRU(key)
		if err != nil {
			return err
		}
	} else {
		if len(current.cache) == current.cap {
			// rewrite
			p := current.tail.prev
			p.prev.next = current.tail
			current.tail.prev = p.prev
			delete(current.cache, p.key)
		}
		newNode := &dll{
			next:  current.head,
			prev:  current.head.next,
			key:   key,
			value: value,
		}
		current.cache[key] = newNode
		current.head.next.prev = newNode
		current.head.next = newNode
	}

	return nil
}

func main() {
	c := newLRUCache(3)
	c.setLRU(1, 1)
	c.setLRU(2, 2)
	c.setLRU(3, 3)
	c.setLRU(4, 4)
	c.getLRU(2)
	c.setLRU(5, 5)
	fmt.Printf("%#v", c)
}
