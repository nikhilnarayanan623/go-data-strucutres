package slist

// To remove a value from the beginning of the singly linked list
func (c *SinglyLinkedList[T]) RemoveFromBeginning() error {
	if c.head == nil {
		return ErrEmptyList
	}

	// if the list have only one value left then remove the value in the tail
	if c.head == c.tail {
		c.tail = nil
	}
	c.head = c.head.next
	c.length--
	return nil
}

// To remove a value from the end of the singly linked list
func (c *SinglyLinkedList[T]) RemoveFromEnding() error {
	if c.head == nil {
		return ErrEmptyList
	}

	if c.head == c.tail {
		c.head = nil
		c.tail = nil
		c.length = 0
		return nil
	}

	current := c.head

	// find the previous node of tail
	for current.next != c.tail {
		current = current.next
	}

	current.next = nil
	c.tail = current
	c.length--

	return nil
}

// To remove a data if its exist else remove error
func (s *SinglyLinkedList[T]) RemoveByData(dataToRemove T) error {

	if s.head == nil {
		return ErrEmptyList
	}

	if s.head.data == dataToRemove {
		// check this is the only value in list then remove the tail
		if s.head == s.tail {
			s.tail = nil
		}
		// if its last value head will be head's next(nil) else just remove the head
		s.head = s.head.next
		s.length--
		return nil
	}

	current := s.head
	// find previous node of the data to remove
	for current.next != nil && current.next.data != dataToRemove {
		current = current.next
	}

	// if current node is tail means data not found on any next node
	if current == s.tail {
		return ErrDataNotExist
	}

	// if the next node is tail then assign current as tail
	if current.next == s.tail {
		s.tail = current
	}
	current.next = current.next.next // just remove the next node that hold the data to remove
	s.length--
	return nil
}

func (s *SinglyLinkedList[T]) RemoveByPosition(position int) error {

	if s.head == nil {
		return ErrEmptyList
	}
	// check given position in the range of 1 to the length
	if position < 1 || position > s.length {
		return ErrInvalidPosition
	}

	// check the position is head
	if position == 1 {
		// if list contain only on data then remove tail also
		if s.head == s.tail {
			s.tail = nil
		}
		// if list have only one data then head will be nil otherwise just remove then head
		s.head = s.head.next
		s.length--
		return nil
	}

	current := s.head
	// find the previous node of the position
	for i := 2; i < position; i++ { // current already on head so we can start from two (list: 1 2 3; pos: 2; no need to loop(and loop will not run also))
		current = current.next
	}

	// check the next node(remove node) is tail then assign current node as tail
	if current.next == s.tail {
		s.tail = current
	}
	// current node's next is tail then current's next will be assign nill; otherwise it will just remove then next node
	current.next = current.next.next
	s.length--
	return nil
}

// To remove all data in the linked list
func (s *SinglyLinkedList[T]) RemoveAll() {

	s.head = nil
	s.tail = nil
	s.length = 0
}
