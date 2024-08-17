package main

import (
	"fmt"
)

/*
	完成一个缓存，实现 put get putall 方法，其中都使用 O(1) 的时间复杂度
*/

type Cache struct {
	map1 map[string]string
	// 判断是否存在
	map2         map[string]string
	defaultValue string
}

func NewCache() *Cache {
	return &Cache{
		map1: make(map[string]string),
		map2: make(map[string]string),
	}
}

func (c *Cache) put(key string, value string) {
	c.map1[key] = value
}

func (c *Cache) get(key string) string {
	// 先读当前的 有就返回，没有代表使历史
	if v, ok := c.map1[key]; ok {
		return v
	}
	// 历史纪录有，代表要返回 putall
	if _, ok := c.map2[key]; ok {
		return c.defaultValue
	}
	return ""
}

func (c *Cache) putAll(value string) {
	c.defaultValue = value
	// 用于记录历史纪录使用
	c.map2 = c.map1
	// 更新当前读的 map
	c.map1 = make(map[string]string)
}

func main() {
	cache := NewCache()

	cache.put("key1", "a")
	cache.put("key2", "b")
	cache.put("key3", "c")

	fmt.Println(cache.get("key1"))
	fmt.Println(cache.get("key2"))
	fmt.Println(cache.get("key3"))

	cache.putAll("new value")

	fmt.Println(cache.get("key1"))
	fmt.Println(cache.get("key2"))
	fmt.Println(cache.get("key3"))
}
