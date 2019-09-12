package lfucgo

type Node struct {
	Data       interface{}
	Next, Prev *Node
	Parent interface{}
}

//List represents a set of nodes (DLL)
type List struct {
	len  int
	Head *Node
}

// Init initializes or clears list l.
func (l *List) Init() *List {
	l.Head = new(Node)
	l.Head.Next = l.Head
	l.Head.Prev = l.Head
	l.len = 0
	return l
}

//Size returns the amount of elements inside the list
func (l *List) Size() int {
	return l.len
}

//Push new element into the beginning of the list
func (l *List) Push(newData interface{}, parentInfo interface{}) *Node {
	n := &Node{newData, new(Node), new(Node), parentInfo}
	if l.len == 0 {
		l.Head.Prev = n
		l.Head.Next = n
		l.Head.Prev.Next = l.Head
		l.Head.Prev.Prev = l.Head
		l.len++
		return n
	}

	l.Head.Next.Prev = n
	n.Next = l.Head.Next
	n.Prev = l.Head
	l.Head.Next = n

	l.len++
	return n
}

//Append new item to the end of the list
func (l *List) Append(newData interface{}, parentInfo interface{}) {
	n := &Node{newData, new(Node), new(Node), parentInfo}
	if l.len == 0 {
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
	l.len++
}
