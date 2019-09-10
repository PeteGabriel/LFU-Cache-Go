package lfucgo

import (
	"testing"
)

func TestPushNewElementMustIncreaseSize(t *testing.T) {
	list := List{0, new(node)}
	list.Push(222)

	if list.Len != 1 {
		t.Errorf("Push operation must increase size of list.")
	}
}

func TestPushNewElementsAndAssertBehavior(t *testing.T) {
	list := List{0, new(node)}
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

	if list.Len != 2 {
		t.Errorf("Push operation must increase size of list.")
	}
}

func TestListWithNoElements(t *testing.T) {
	list := new(List)
	if list.Len != 0 {
		t.Errorf("List with no elements must not contain any element.")
	}
}

func TestAppendOneElementIncreaseLenghtOfList(t *testing.T) {
	list := List{0, new(node)}
	list.Append(19)

	if list.Len != 1 {
		t.Errorf("List must support the append operation.")
	}
}

func TestAppendOneElementToTheEndOfList(t *testing.T) {
	list := List{0, new(node)}
	list.Append(19)

	if list.Head.Next.Data != 19 {
		t.Errorf("First expected element not found.")
	}

	list.Append(23)

	if list.Head.Prev.Data != 23 {
		t.Errorf("Second expected element not found.")
	}

	if list.Len != 2 {
		t.Errorf("Append operation must increase lenght of list.")
	}

}

func TestAppendVariousElementsAndAssertCorrectSize(t *testing.T) {
	list := List{0, new(node)}
	list.Append(19)
	list.Append(1)
	list.Append(20)
	list.Append(2)

	if list.Len != 4 {
		t.Errorf("List must support the append operation.")
	}

	if list.Head.Next.Data != 19 {
		t.Errorf("First expected element not found.")
	}
	if list.Head.Next.Next.Data != 1 {
		t.Errorf("Second expected element not found.")
	}
	if list.Head.Next.Next.Next.Data != 20 {
		t.Errorf("Second expected element not found.")
	}
	if list.Head.Next.Next.Next.Next.Data != 2 {
		t.Errorf("Second expected element not found.")
	}

}
