package slist

import "fmt"

// To print value in the singly linked list in forward
func (c *SinglyLinkedList[T]) PrintAll() {
	if c.head == nil {
		fmt.Println("singly linked list is empty")
		return
	}

	current := c.head
	fmt.Println("values in the singly linked list")
	for current != nil {
		fmt.Print(current.data, " ")
		current = current.next
	}
	fmt.Println()
}

func (c *SinglyLinkedList[T]) PrintAllReverse() {
	if c.head == nil {
		fmt.Println("singly linked list is empty")
		return
	}
	fmt.Println("values in the singly linked list in reverse")
	c.printAllReverseHelper(c.head)
	fmt.Println()
}

func (c *SinglyLinkedList[T]) printAllReverseHelper(node *node[T]) {
	if node == nil {
		return
	}
	c.printAllReverseHelper(node.next)
	fmt.Print(node.data, " ")
}
