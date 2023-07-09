package slist

// To insert a new data into the beginning of singly linked list
func (c *SinglyLinkedList[T]) InsertAtBeginning(data T) {

	newNode := &node[T]{
		data: data,
	}

	if c.head == nil {
		c.tail = newNode
	}

	newNode.next = c.head
	c.head = newNode
	c.length++
}

// To insert a new data into the ending of singly linked list
func (c *SinglyLinkedList[T]) InsertAtEnding(data T) {

	newNode := &node[T]{
		data: data,
	}

	if c.head == nil {
		c.head = newNode
	} else {
		c.tail.next = newNode
	}
	c.tail = newNode
	c.length++
}

// To insert a data at a specific position in singly linked list
func (c *SinglyLinkedList[T]) InsertAtPosition(position int, data T) error {
	// if position less than 1 or greater then the length (except position 1)
	if position < 1 || position > c.length && position != 1 {

		return ErrInvalidPosition
	}

	newNode := &node[T]{
		data: data,
	}

	if position == 1 {
		newNode.next = c.head
		c.head = newNode
		if c.tail == nil {
			c.tail = newNode
		}
		c.length++
		return nil
	}

	current := c.head
	// find value node of position -1 (start from 2 because start from current(head) which already took 1 step)
	for i := 2; i < position; i++ {
		current = current.next
	}

	newNode.next = current.next
	current.next = newNode
	c.length++

	return nil
}
