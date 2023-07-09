package slist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveFromBeginning(t *testing.T) {

	testCases := map[string]func(t *testing.T, s *SinglyLinkedList[string]){
		"RemoveFromEmptyShouldReturnError": func(t *testing.T, s *SinglyLinkedList[string]) {
			err := s.RemoveFromBeginning()
			assert.EqualError(t, err, ErrEmptyList.Error())
		},
		"OneInsertThenRemoveListShouldBeEmpty": func(t *testing.T, s *SinglyLinkedList[string]) {
			s.InsertAtBeginning("data")
			err := s.RemoveFromBeginning()
			assert.NoError(t, err)
			assert.True(t, s.IsEmpty())
		},
		"TwoInsertOneRemoveListShouldNotBeEmpty": func(t *testing.T, s *SinglyLinkedList[string]) {
			s.InsertAtBeginning("data1")
			s.InsertAtEnding("data2")
			err := s.RemoveFromBeginning()
			assert.NoError(t, err)
			assert.False(t, s.IsEmpty())
		},
		"TwoInsertOnBeginningAndRemoveThenHeadAndTailShouldBeSameWithFirstVal": func(t *testing.T, s *SinglyLinkedList[string]) {
			firstData := "data1"
			s.InsertAtBeginning(firstData)
			s.InsertAtBeginning("data2")
			err := s.RemoveFromBeginning()
			assert.NoError(t, err)
			assert.Equal(t, s.head, s.tail)
			assert.Equal(t, firstData, s.head.data)
			assert.Equal(t, firstData, s.tail.data)
		},
		"FourInsertOnEndingAndRemoveThenTailShouldBeThirdValue": func(t *testing.T, s *SinglyLinkedList[string]) {
			thirdData := "300"
			s.InsertAtEnding("100")
			s.InsertAtEnding("200")
			s.InsertAtEnding(thirdData)
			s.InsertAtEnding("400")

			err := s.RemoveFromEnding()
			assert.NoError(t, err)
			assert.Equal(t, thirdData, s.tail.data)

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

func TestRemoveFromEnding(t *testing.T) {

	testCases := map[string]func(t *testing.T, s *SinglyLinkedList[string]){
		"RemoveFromEmptyShouldReturnError": func(t *testing.T, s *SinglyLinkedList[string]) {
			err := s.RemoveFromEnding()
			assert.EqualError(t, err, ErrEmptyList.Error())
		},
		"OneInsertThenRemoveListShouldBeEmpty": func(t *testing.T, s *SinglyLinkedList[string]) {
			s.InsertAtBeginning("data")
			s.RemoveFromEnding()
			assert.True(t, s.IsEmpty())
		},
		"TwoInsertOneRemoveListShouldNotBeEmpty": func(t *testing.T, s *SinglyLinkedList[string]) {
			s.InsertAtBeginning("data1")
			s.InsertAtEnding("data2")
			s.RemoveFromEnding()
			assert.False(t, s.IsEmpty())
		},
		"TwoInsertOnBeginningAndRemoveHeadAndTailShouldBeSameWithSecondVal": func(t *testing.T, s *SinglyLinkedList[string]) {
			secondData := "data2"
			s.InsertAtBeginning("data1")
			s.InsertAtBeginning(secondData)
			s.RemoveFromEnding()
			assert.Equal(t, s.head, s.tail)
			assert.Equal(t, secondData, s.head.data)
			assert.Equal(t, secondData, s.tail.data)
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

func TestRemoveByData(t *testing.T) {

	testCases := map[string]func(t *testing.T, s *SinglyLinkedList[int]){
		"RemoveFromEmptyListShouldReturnErrorOfEmptyList": func(t *testing.T, s *SinglyLinkedList[int]) {
			err := s.RemoveByData(123)
			assert.EqualError(t, err, ErrEmptyList.Error())
		},
		"RemoveNotExistDataShouldReturnErrorDataNotExist": func(t *testing.T, s *SinglyLinkedList[int]) {
			s.InsertAtBeginning(12)
			notExistData := 100
			err := s.RemoveByData(notExistData)
			assert.EqualError(t, err, ErrDataNotExist.Error())
		},
		"RemoveFromOneDataListShouldMakeListEmpty": func(t *testing.T, s *SinglyLinkedList[int]) {
			s.InsertAtBeginning(100)
			err := s.RemoveByData(100)
			assert.NoError(t, err)
			assert.True(t, s.IsEmpty())
		},
		"TwoDataListRemoveHeadDataThenHeadAndTailShouldSame": func(t *testing.T, s *SinglyLinkedList[int]) {
			s.InsertAtBeginning(100)
			s.InsertAtBeginning(200)
			err := s.RemoveByData(100)
			assert.NoError(t, err)
			assert.Equal(t, s.head, s.tail)
			assert.Equal(t, 200, s.tail.data)
		},
		"NthDataRemoveTailDataShouldBeN-1": func(t *testing.T, s *SinglyLinkedList[int]) {
			n := 4
			for i := 1; i <= n; i++ { // insert 1 to n data as i
				s.InsertAtEnding(i)
			}
			err := s.RemoveByData(n) // n the data is n and n-1 th data is n-1
			assert.NoError(t, err)
			assert.Equal(t, n-1, s.tail.data)
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

func TestRemoveByPosition(t *testing.T) {
	testCases := map[string]func(t *testing.T, s *SinglyLinkedList[int]){
		"RemoveFromEmptyListShouldReturnErrorOfEmptyList": func(t *testing.T, s *SinglyLinkedList[int]) {
			err := s.RemoveByPosition(12)
			assert.EqualError(t, err, ErrEmptyList.Error())
		},
		"InvalidPositionShouldReturnError": func(t *testing.T, s *SinglyLinkedList[int]) {
			s.InsertAtEnding(10)
			s.InsertAtEnding(20)
			err := s.RemoveByPosition(10)
			assert.EqualError(t, err, ErrInvalidPosition.Error())
		},
		"OneInsertAndRemoveThePositionShouldMakeTheListEmpty": func(t *testing.T, s *SinglyLinkedList[int]) {
			s.InsertAtEnding(10)
			err := s.RemoveByPosition(1)
			assert.NoError(t, err)
			assert.True(t, s.IsEmpty())
		},
		"TwoDataListRemoveHeadPositionThenHeadAndTailShouldBeSame": func(t *testing.T, s *SinglyLinkedList[int]) {
			s.InsertAtEnding(10)
			s.InsertAtEnding(20)
			err := s.RemoveByPosition(1)
			assert.NoError(t, err)
			assert.Equal(t, s.head, s.tail)
			assert.Equal(t, 20, s.head.data)
		},
		"TwoDataListRemoveTailPositionThenHeadAndTailShouldBeSame": func(t *testing.T, s *SinglyLinkedList[int]) {
			s.InsertAtEnding(10)
			s.InsertAtEnding(20)
			err := s.RemoveByPosition(2)
			assert.NoError(t, err)
			assert.Equal(t, s.head, s.tail)
			assert.Equal(t, 10, s.head.data)
		},
		"ThreeDataListRemoveTailPositionThenHeadAndTailShouldDifferentAndDataShouldBeSecond": func(t *testing.T, s *SinglyLinkedList[int]) {
			secondData := 20
			s.InsertAtEnding(10)
			s.InsertAtEnding(secondData)
			s.InsertAtEnding(30)

			err := s.RemoveByPosition(3)
			assert.NoError(t, err)
			assert.NotEqual(t, s.head, s.tail)
			assert.Equal(t, secondData, s.tail.data)
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
