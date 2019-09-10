package lfucgo

import (
	"testing"
)

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
