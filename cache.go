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
  _, exists := c.elems[key]
	if exists {
		return false
	}

	c.elems[key] = elem

	if c.freq.Size() == 0 {
    c.freq.Push(&Pair{1, elem})
		return true
	}

  freqOneNode := &c.freq.Head.Next.Data.(Pair)
  if freqOneNode.v == 1 {
    l:= &freqOneNode.d.(List)
		l.Push(elem)
		return true
	}

  return false
}

//Remove element by key
//This operation costs O(1).
func (c *Cache) Remove(key int) {

  delete(c.elems, key)
}

func (c *Cache) Size() int {
 return len(c.elems)
}
