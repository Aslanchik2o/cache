package cache

import (
	"fmt"
)

type Cache struct {
	items map[string]Item
}

type Item struct {
	Value interface{}
}

func New() *Cache {

	items := make(map[string]Item)

	cache := Cache{
		items: items,
	}
	return &cache
}

func (c *Cache) Set(key string, value interface{}) {
	c.items[key] = Item{
		Value: value,
	}
}

func (c *Cache) Get(key string) interface{} {
	item := c.items[key]
	return item.Value
}

func (c *Cache) Delete(key string) interface{} {
	item := c.items[key]
	delete(c.items, key)
	fmt.Println(item)
	return item
}
