package slist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertAtBeginning(t *testing.T) {

	testCases := map[string]func(t *testing.T, s *SinglyLinkedList[int]){

		"AfterInsertHeadAndTailNotNil": func(t *testing.T, s *SinglyLinkedList[int]) {
			s.InsertAtBeginning(12)
			assert.NotNil(t, s.head, s.tail)
		},
		"AfterInsertListShouldNotEmpty": func(t *testing.T, s *SinglyLinkedList[int]) {
			s.InsertAtBeginning(100)
			assert.False(t, s.IsEmpty())
		},
		"OneInsertHeadAndTailShouldBeSame": func(t *testing.T, s *SinglyLinkedList[int]) {
			s.InsertAtBeginning(1)
			assert.Equal(t, s.head, s.tail)
			assert.Equal(t, 1, s.length)
		},

		"TwoInsertHeadAndTailShouldDifferent": func(t *testing.T, s *SinglyLinkedList[int]) {
			s.InsertAtBeginning(20)
			s.InsertAtBeginning(10)
			assert.NotEqual(t, s.head, s.tail)
		},

		"InsertCountAndListLengthShouldSame": func(t *testing.T, s *SinglyLinkedList[int]) {
			count := 12
			for i := 1; i <= count; i++ {
				s.InsertAtBeginning(i)
			}
			assert.Equal(t, count, s.Len())
		},

		"MultipleInsertHeadDataIsInputDataAndTailShouldBeFirstData": func(t *testing.T, s *SinglyLinkedList[int]) {

			firstData := 100
			s.InsertAtBeginning(firstData)
			for i := 1; i <= 5; i++ {
				s.InsertAtBeginning(i)
				assert.Equal(t, i, s.head.data)
				assert.Equal(t, firstData, s.tail.data)
			}
		},
	}

	for name, test := range testCases {

		t.Run(name, func(t *testing.T) {
			test := test
			t.Parallel()

			s := NewSinglyLinkedList[int]()
			test(t, s)
		})
	}
}

func TestInsertAtEnding(t *testing.T) {

	testCases := map[string]func(t *testing.T, s *SinglyLinkedList[string]){

		"AfterInsertHeadAndTailNotNil": func(t *testing.T, s *SinglyLinkedList[string]) {
			s.InsertAtBeginning("data")
			assert.NotNil(t, s.head, s.tail)

		},
		"AfterInsertListShouldNotEmpty": func(t *testing.T, s *SinglyLinkedList[string]) {
			s.InsertAtBeginning("data")
			assert.False(t, s.IsEmpty())
		},
		"OneInsertHeadAndTailShouldBeSame": func(t *testing.T, s *SinglyLinkedList[string]) {
			s.InsertAtEnding("value")
			assert.NotNil(t, s.head, s.tail)
			assert.Equal(t, s.head, s.tail)
		},

		"TwoInsertHeadAndTailShouldDifferent": func(t *testing.T, s *SinglyLinkedList[string]) {
			s.InsertAtEnding("data1")
			s.InsertAtEnding("data2")
			assert.NotEqual(t, s.head, s.tail)
		},

		"InsertCountAndListLengthShouldSame": func(t *testing.T, s *SinglyLinkedList[string]) {

			dataArr := []string{"data1", "data2", "data3"}

			for _, data := range dataArr {
				s.InsertAtEnding(data)
			}

			assert.Equal(t, len(dataArr), s.Len())
		},

		"MultipleInsertHeadDataShouldBeFirstDataTailDataShouldInsertData": func(t *testing.T, s *SinglyLinkedList[string]) {
			firstData := "data"
			s.InsertAtEnding(firstData)

			dataArr := []string{"data1", "data2", "data3"}

			for _, data := range dataArr {
				s.InsertAtEnding(data)
				assert.Equal(t, data, s.tail.data)
				assert.Equal(t, firstData, s.head.data)
			}
		},
	}

	for name, test := range testCases {

		t.Run(name, func(t *testing.T) {
			test := test
			t.Parallel()

			s := NewSinglyLinkedList[string]()
			test(t, s)
		})
	}
}

func TestInsertAtPosition(t *testing.T) {

	testCases := map[string]func(t *testing.T, s *SinglyLinkedList[int]){

		"InvalidPositionShouldReturnError": func(t *testing.T, s *SinglyLinkedList[int]) {
			err := s.InsertAtPosition(100, 10)
			assert.EqualError(t, err, ErrInvalidPosition.Error())
		},
		"InsertAtPositionOneWhenItsEmptyDataShouldBeEqualWithData": func(t *testing.T, s *SinglyLinkedList[int]) {
			err := s.InsertAtPosition(1, 10)
			assert.NoError(t, err)
			assert.Equal(t, s.head, s.tail)
			assert.Equal(t, 10, s.head.data)
		},
		"InsertAtPosition1ShouldBeOnHead": func(t *testing.T, s *SinglyLinkedList[int]) {
			err := s.InsertAtPosition(1, 10)
			assert.NoError(t, err)

			assert.Equal(t, 10, s.head.data)
		},
		"InsertValueAtNth": func(t *testing.T, s *SinglyLinkedList[int]) {
			// insert some data on list
			err := s.InsertAtPosition(1, 10)
			assert.NoError(t, err)
			for i := 1; i <= 5; i++ {
				err = s.InsertAtPosition(i, i)
				assert.NoError(t, err)
			}

			// insert a value at valid nth pos
			n, data := 3, 100
			err = s.InsertAtPosition(n, data)
			assert.NoError(t, err)

			temp := s.head
			for i := 1; i < n; i++ { // loop only n-1(because 1 already taken by head)
				temp = temp.next
			}
			assert.Equal(t, data, temp.data)
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			test := test

			s := NewSinglyLinkedList[int]()
			test(t, s)
		})
	}
}
