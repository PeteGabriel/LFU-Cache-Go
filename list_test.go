package lfucgo

import (
	"testing"
)

func TestPushNewElementAsList(t *testing.T) {
	internalList := new(List).Init()
	internalList.Push(222)

	list := new(List).Init()
	list.Push(internalList)

	if list.Size() != 1 {
		t.Errorf("Push operation must increase size of list.")
	}

	tmpl := list.Head.Next.Data.(*List)
	if tmpl.Head.Next.Data != 222 {
		t.Errorf("Push operation must insert elements into list.")
	}
}

func TestPushNewElementAsBasicType(t *testing.T) {
	list := new(List).Init()
	list.Push(222)

	if list.Size() != 1 {
		t.Errorf("Push operation must increase size of list.")
	}

	data := list.Head.Next.Data.(int)
	if data != 222 {
		t.Errorf("Push operation must insert elements into list.")
	}
}

func TestPushNewElementMustIncreaseSize(t *testing.T) {
	list := new(List).Init()
	list.Push(222)

	if list.Size() != 1 {
		t.Errorf("Push operation must increase size of list.")
	}
}

func TestPushNewElementsAndAssertBehavior(t *testing.T) {
	list := new(List).Init()
	list.Push(222)

	if list.Head.Next.Data != 222 {
		t.Errorf("Push operation must add to the beginning of the list.")
	}

	list.Push(2)

	if list.Head.Next.Data != 2 {
		t.Errorf("Push operation must add to the beginning of the list.")
	}

	if list.Head.Next.Next.Data != 222 {
		t.Errorf("Push operation must maintain list integrity.")
	}

	if list.Head.Prev.Data != 222 {
		t.Errorf("Push operation must maintain list integrity.")
	}

	if list.Size() != 2 {
		t.Errorf("Push operation must increase size of list.")
	}
}

func TestListWithNoElements(t *testing.T) {
	list := new(List).Init()
	if list.Len != 0 {
		t.Errorf("List with no elements must not contain any element.")
	}
}

func TestAppendOneElementIncreaseLenghtOfList(t *testing.T) {
	list := new(List).Init()
	list.Append(19)

	if list.Size() != 1 {
		t.Errorf("List must support the append operation.")
	}
}

func TestAppendOneElementToTheEndOfList(t *testing.T) {
	list := new(List).Init()
	list.Append(19)

	if list.Head.Next.Data != 19 {
		t.Errorf("First expected element not found.")
	}

	list.Append(23)

	if list.Head.Prev.Data != 23 {
		t.Errorf("Second expected element not found.")
	}

	if list.Size() != 2 {
		t.Errorf("Append operation must increase lenght of list.")
	}

}

func TestAppendVariousElementsAndAssertCorrectSize(t *testing.T) {
	list := new(List).Init()
	list.Append(19)
	list.Append(1)
	list.Append(20)
	list.Append(2)

	if list.Size() != 4 {
		t.Errorf("List must support the append operation.")
	}

	first := list.Head.Next.Data.(int)
	if first != 19 {
		t.Errorf("First expected element not found.")
	}
	second := list.Head.Next.Next.Data.(int)
	if second != 1 {
		t.Errorf("Second expected element not found.")
	}
	third := list.Head.Next.Next.Next.Data.(int)
	if third != 20 {
		t.Errorf("Second expected element not found.")
	}
	fourth := list.Head.Next.Next.Next.Next.Data.(int)
	if fourth != 2 {
		t.Errorf("Second expected element not found.")
	}
}
