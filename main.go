package cache

import (
	"fmt"
)

type Cache struct {
	items map[int]Item
}

type Item struct {
	Value interface{}
}

func New() *Cache {

	items := make(map[int]Item)

	cache := Cache{
		items: items,
	}
	return &cache
}

func (c *Cache) Set(key int, value interface{}) {
	c.items[key] = Item{
		Value: value,
	}
}

func (c *Cache) Get(key int) interface{} {
	item := c.items[key]
	return item.Value
}

func (c *Cache) Delete(key int) interface{} {
	item := c.items[key]
	delete(c.items, key)
	fmt.Println(item)
	return item
}
