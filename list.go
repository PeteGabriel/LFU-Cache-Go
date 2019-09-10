package lfucgo

type node struct {
	Data       interface{}
	Next, Prev *node
}

//List represents a set of nodes (DLL)
type List struct {
	Len  int
	Head *node
}

// Init initializes or clears list l.
func (l *List) Init() *List {
	l.Head = new(node)
	l.Head.Next = l.Head
	l.Head.Prev = l.Head
	l.Len = 0
	return l
}

//Push new element into the beginning of the list
func (l *List) Push(newData interface{}) {
	n := &node{newData, new(node), new(node)}
	if l.Len == 0 {
		l.Head.Prev = n
		l.Head.Next = n
		l.Head.Prev.Next = l.Head
		l.Head.Prev.Prev = l.Head
	} else {
		l.Head.Next.Prev = n
		n.Next = l.Head.Next
		n.Prev = l.Head
		l.Head.Next = n
	}
	l.Len++
}

//Append new item to the end of the list
func (l *List) Append(newData interface{}) {
	n := &node{newData, new(node), new(node)}
	if l.Len == 0 {
		l.Head.Prev = n
		l.Head.Next = n
		l.Head.Prev.Next = l.Head
		l.Head.Prev.Prev = l.Head
	} else {
		l.Head.Prev.Next = n
		n.Prev = l.Head.Prev
		l.Head.Prev = n
		l.Head.Prev.Next = l.Head
	}
	l.Len++
}
