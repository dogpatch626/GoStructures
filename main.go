package main

import (
	LinkedList "DataStructures/structures"
	"fmt"
)

func main() {
	meow := LinkedList.CreateLinkedList(1)
	meow.Insert(2)
	fmt.Println(meow)
}
