package lfucgo

import "testing"

func TestCacheCanSaveNewElement(t *testing.T){
  cache:= new(Cache).Init()
	key:= 1
	elem:= "NewElement"
	inserted := cache.Insert(key, elem)
  if cache.Size() != 1 || !inserted {
		t.Error("Cache must insert new elements.")
	}
}

func TestCacheCannotSaveElementWithExistentKey(t *testing.T){
  cache:= new(Cache).Init()
	key:= 1
	elem:= "NewElement"
	elem2:= "NewElement2"
	cache.Insert(key, elem)
	insertedSecond := cache.Insert(key, elem2)

  if insertedSecond {
		t.Error("Cache cannot insert elements with same key.")
	}

	if cache.Size() != 1 {
		t.Error("Cache cannot insert elements with same key.")
	}
}


func TestCacheCanSaveMoreThanOneElement(t *testing.T){
  cache:= new(Cache).Init()

	elem:= "NewElement"
	elem2:= "NewElement2"
	elem3:= "NewElement3"
	cache.Insert(1, elem)
	cache.Insert(2, elem2)
	cache.Insert(3, elem3)

  if cache.Size() != 3 {
		t.Error("Cache cannot insert elements with same key.")
	}
}

func TestCacheShouldReturnNilWhenElementNotFound(t *testing.T){
  cache:= new(Cache).Init()
	key:= 1
	elem:= "NewElement"
	cache.Insert(key, elem)
	cachedElem, _ := cache.Access(89)

	if cache.Size() != 1 {
		t.Error("Cache cannot insert elements with same key.")
	}

 if cachedElem != nil {
		t.Error("Cache must be able to access element by key.")
	}

}

func TestCacheShouldProvideAccessToElementsByKey(t *testing.T){
  cache:= new(Cache).Init()
	key:= 1
	elem:= "NewElement"
	cache.Insert(key, elem)
	cachedElem, _ := cache.Access(key)

	if cache.Size() != 1 {
		t.Error("Cache cannot insert elements with same key.")
	}

 if cachedElem == nil {
		t.Error("Cache must be able to access element by key.")
	}

}

func TestCacheShouldIncreaseFrequencyToAccessedElement(t *testing.T){
  cache:= new(Cache).Init()
	key:= 1
	elem:= "NewElement"
	cache.Insert(key, elem)
	cache.Access(key)
  cache.Access(key)
  cachedElem, freq := cache.Access(key)

	if cache.Size() != 1 {
		t.Errorf("Cache cannot insert elements with same key. Current Size is %d", cache.Size())
	}

	if cachedElem == nil {
		t.Error("Cache must be able to access element by key.")
	}

	if freq != 4 {
		t.Errorf("Cache must increase frequency of elements accessed. Current freq: %d", freq)
	}
}