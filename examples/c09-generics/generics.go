package main

import "fmt"

type Cache[Key comparable, Value any] struct {
	entries map[Key]Value
}

func newCache[Key comparable, Value any]() *Cache[Key, Value] {
	return &Cache[Key, Value]{
		entries: make(map[Key]Value),
	}
}

func (cache *Cache[Key, Value]) set(key Key, value Value) {
	cache.entries[key] = value
}

func (cache *Cache[Key, Value]) get(key Key) (Value, bool) {
	value, ok := cache.entries[key]
	return value, ok
}

func main() {
	ports := newCache[string, int]()
	ports.set("http", 80)
	ports.set("ssh", 22)

	port, ok := ports.get("ssh")
	fmt.Println("ssh in cache?:", ok, "port:", port)

	port, ok = ports.get("smtp")
	fmt.Println("smtp in cache?:", ok, "port:", port)
}
