package slist

// To check given data is exist on singly linked list
func (s *SinglyLinkedList[T]) IsDataExist(dataToCheck T) bool {

	if s.head == nil {
		return false
	}

	current := s.head

	for current != nil {
		if current.data == dataToCheck {
			return true
		}
		current = current.next
	}
	return false
}
