package LinkedList

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head Node
}

func (n *LinkedList) Insert(data int) {
	idk := Node{value: data, next: nil}
	n.head.next = &idk

}

func CreateLinkedList(data int) LinkedList {
	ToBeReturned := LinkedList{head: Node{value: data, next: nil}}
	return ToBeReturned
}
