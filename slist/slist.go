package slist

type node struct {
	data int
	next *node
}

type SinglyLinkedList struct {
	head *node
	tail *node
}
