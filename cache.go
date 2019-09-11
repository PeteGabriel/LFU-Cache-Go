package lfucgo

type Cache struct {
  elems    map[int]interface{}
  freq *List
}

// Init initializes or clears list l.
func (c *Cache) Init() *Cache {
  c.elems = make(map[int]interface{})
  c.freq = new(List).Init()

  c.freq.Push(&Pair{1, new(List).Init()})
  return c
}


//Get element by key without removing it.
//This operation costs O(1).
func (c *Cache) GetElement(key int) interface{} {
  return nil
}

//Insert element by key
//This operation costs O(1).
func (c *Cache) Insert(key int, elem interface{}) bool {
  c.elems[key] = elem
  c.freq.Push
  return true
}

//Remove element by key
//This operation costs O(1).
func (c *Cache) Remove(key int) {

  delete(c.elements, key)
}
