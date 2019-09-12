package lfucgo


type Cache struct {
  elems    map[int]interface{}
  freq *List
}

// Init initializes or clears list l.
func (c *Cache) Init() *Cache {
  c.elems = make(map[int]interface{})
  c.freq = new(List).Init()

  p:= &Pair{1, new(List).Init()}
  c.freq.Push(p, p)
  return c
}


//Access element by key without removing it.
//This operation costs O(1).
func (c *Cache) Access(key int) (interface{}, int) {

	_, exists := c.elems[key]
	if exists {
		elem:= c.elems[key].(*Node)
		freq:= elem.Parent.(*Pair)
		freq.v++
		return elem, freq.v
	}

  return nil, 0

}

//Insert element by key
//This operation costs O(1).
func (c *Cache) Insert(key int, elem interface{}) bool {
  _, exists := c.elems[key]
	if exists {
		return false
	}


	if c.freq.Size() == 0 {
    p:= &Pair{1, elem}
    n := c.freq.Push(p,p)
    c.elems[key] = n

		return true
	}

  freqOneNode := c.freq.Head.Next.Data.(*Pair)
  if freqOneNode.v == 1 {
    l:= freqOneNode.d.(*List)
		n:= l.Push(elem, c.freq.Head.Next.Data)
		c.elems[key] = n
		return true
	}

  return false
}

//Remove element by key
//This operation costs O(1).
func (c *Cache) Remove(key int) {

  delete(c.elems, key)
}

//Size of cache
func (c *Cache) Size() int {
 return len(c.elems)
}
