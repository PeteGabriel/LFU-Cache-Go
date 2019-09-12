package lfucgo

import "testing"

func TestNewCacheIsEmpty(t *testing.T){
  cache:= new(Cache).Init()
	if cache.GetElement(-1) != nil {
		t.Error("Recent created cache must not contain any elements")
	}
}

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
	cache.Insert(key, elem2)
	if cache.Size() != 1 {
		t.Error("Cache cannot insert elements with same key.")
	}
}
