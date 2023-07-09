package main

import (
	"fmt"

	"github.com/nikhilnarayanan623/go-data-strucutres/slist"
)

func main() {

}

// Testing singly linked list print functions manually
func DisplaySinglyLinkedList() {

	s := slist.NewSinglyLinkedList[int]()

	s.InsertAtBeginning(12)
	s.InsertAtEnding(10)
	s.InsertAtPosition(100, 2)

	s.PrintAll()
	s.PrintAllReverse()
}

func CheckSinglyLinkedListExistData() {

	arr := []string{"hii", "hello", "hey"}

	s := slist.SliceToLinkedList(arr)

	fmt.Println("exist", s.IsDataExist("hii"))
}
