package slist

// To return a new singly list which consist the given array values
func SliceToLinkedList[T dataType](arr []T) *SinglyLinkedList[T] {

	sList := NewSinglyLinkedList[T]()

	for _, data := range arr {
		sList.InsertAtEnding(data)
	}
	return sList
}

// To add given slice of value to the existing singly linked list
func (s *SinglyLinkedList[T]) AddSliceToTheList(arr []T) {
	for _, data := range arr {
		s.InsertAtEnding(data)
	}
}
