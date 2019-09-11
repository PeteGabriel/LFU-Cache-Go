package lfucgo

//Cache definition
type Cache struct {
  freq *List
	elem map[int]interface{}
}

//New cache
func New() *Cache {
	return new(Cache).init()
}

func (c *Cache) init() *Cache {
  c.elem = make(map[int]interface{})
	c.freq = new(List).Init()
	return c
}