package LinkedList

type Node struct {
	value int
	next  *Node
}
type LinkedList struct {
	head *Node
}

func (i LinkedList) Insert(x int) {
	nex := &Node{value: x, next: nil}
	i.head.next = nex
}

func (LinkedList) Create(data int) LinkedList {
	list := Node{value: data, next: nil}
	return LinkedList{head: &list}
}
