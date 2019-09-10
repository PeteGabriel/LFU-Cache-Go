package lfucgo

type node struct {
	Data       int
	Next, Prev *node
}

//List represents a set of nodes (DLL)
type List struct {
	Len  int
	Head *node
}

//Append new item to the end of the list
func (l *List) Append(newData int) {
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
