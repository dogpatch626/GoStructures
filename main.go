package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head Node
}

func (n *LinkedList) insert(data int) {
	idk := Node{value: data, next: nil}
	n.head.next = &idk

}

func CreateLinkedList(data int) LinkedList {

}

func main() {
	meow := LinkedList{head: Node{value: 1, next: nil}}
	meow.insert(2)
	fmt.Println(meow.head.next)
}
