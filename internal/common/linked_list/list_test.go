package linked_list

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAppend_whenEmptyList_SingletonList(t *testing.T) {
	list := NonValidatedList[int]{}
	list.Append(1)
	require.Equal(t, 1, list.size)
	require.Equal(t, 1, list.head.Value)
	require.Equal(t, 1, list.tail.Value)
	require.Nil(t, list.head.Prev)
	require.Nil(t, list.head.Next)
	require.Nil(t, list.tail.Prev)
	require.Nil(t, list.tail.Next)
}

func TestAppend_whenNonEmpty_AppendValue(t *testing.T) {
	list := NonValidatedList[int]{}
	var prv *Node[int]
	for i := 1; i <= 10; i++ {
		list.Append(i)

		if prv == nil {
			prv = list.Head()
			require.Nil(t, prv.Prev)
			require.Nil(t, prv.Next)
		} else {
			require.Equal(t, i, prv.Next.Value)
			require.Equal(t, prv, prv.Next.Prev)
			prv = prv.Next
		}
	}
}

func TestPrepend_whenEmptyList_SingletonList(t *testing.T) {
	list := NonValidatedList[int]{}
	list.Prepend(1)
	require.Equal(t, 1, list.size)
	require.Equal(t, 1, list.head.Value)
	require.Equal(t, 1, list.tail.Value)
	require.Nil(t, list.head.Prev)
	require.Nil(t, list.head.Next)
	require.Nil(t, list.tail.Prev)
	require.Nil(t, list.tail.Next)
}

func TestPrepend_whenNonEmptyList_PrependValue(t *testing.T) {
	list := NonValidatedList[int]{}
	list.Prepend(1)
	list.Prepend(2)
	require.Equal(t, 2, list.size)
	require.Equal(t, 2, list.head.Value)
	require.Equal(t, 1, list.tail.Value)
	require.Nil(t, list.head.Prev)
	require.Equal(t, list.tail, list.head.Next)
	require.Equal(t, list.head, list.tail.Prev)
	require.Nil(t, list.tail.Next)
}

func TestLength_whenCalled_ReturnLength(t *testing.T) {
	list := NonValidatedList[int]{
		size: 0,
	}
	require.Equal(t, 0, list.Length())
	list.size = 1000
	require.Equal(t, 1000, list.Length())
}

func TestHead_whenCalled_ReturnHead(t *testing.T) {
	list := NonValidatedList[int]{
		head: nil,
	}
	require.Nil(t, list.Head())

	node := &Node[int]{Value: 1}
	list.head = node
	require.Equal(t, node, list.Head())
}

func TestTail_whenCalled_ReturnTail(t *testing.T) {
	list := NonValidatedList[int]{
		tail: nil,
	}
	require.Nil(t, list.Tail())

	node := &Node[int]{Value: 1}
	list.tail = node
	require.Equal(t, node, list.Tail())
}

func TestValidate_whenCalled_Panic(t *testing.T) {
	list := NonValidatedList[int]{}
	require.Panics(t, func() { _ = list.Validate() })
}

func TestNewList_whenCalled_ReturnList(t *testing.T) {
	list := NewList[int]()
	require.NotNil(t, list)
	require.Equal(t, 0, list.Length())
	require.Nil(t, list.Head())
	require.Nil(t, list.Tail())
}

func TestRemove_whenNonExistingIndex_ReturnError(t *testing.T) {
	list := NonValidatedList[int]{
		size: 0,
	}
	err := list.Remove(-1)
	require.Error(t, err)

	err = list.Remove(0)
	require.Error(t, err)
}

func TestRemove_whenSingleton_MakesEmpty(t *testing.T) {
	list := NonValidatedList[int]{
		size: 1,
		head: &Node[int]{Value: 1},
		tail: &Node[int]{Value: 1},
	}
	err := list.Remove(0)
	require.NoError(t, err)
	require.Equal(t, 0, list.size)
	require.Nil(t, list.head)
	require.Nil(t, list.tail)
}

func TestRemove_whenTail_ReplacesTail(t *testing.T) {
	list := NonValidatedList[int]{}
	list.Append(1)
	list.Append(2)
	err := list.Remove(1)
	require.NoError(t, err)
	require.Equal(t, 1, list.size)
	require.Equal(t, 1, list.head.Value)
	require.Equal(t, 1, list.tail.Value)
	require.Nil(t, list.head.Prev)
	require.Nil(t, list.head.Next)
	require.Nil(t, list.tail.Prev)
	require.Nil(t, list.tail.Next)
}

func TestRemove_whenHead_ReplacesHead(t *testing.T) {
	list := NonValidatedList[int]{}
	list.Append(1)
	list.Append(2)
	err := list.Remove(0)
	require.NoError(t, err)
	require.Equal(t, 1, list.size)
	require.Equal(t, 2, list.head.Value)
	require.Equal(t, 2, list.tail.Value)
	require.Nil(t, list.head.Prev)
	require.Nil(t, list.head.Next)
	require.Nil(t, list.tail.Prev)
	require.Nil(t, list.tail.Next)
}

func TestRemove_whenMiddle_RemovesMiddle(t *testing.T) {
	list := NonValidatedList[int]{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	// _middle := list.head.Next

	err := list.Remove(1)
	require.NoError(t, err)
	require.Equal(t, 2, list.size)
	require.Equal(t, 1, list.head.Value)
	require.Equal(t, 3, list.tail.Value)
	require.Nil(t, list.head.Prev)
	require.Equal(t, list.tail, list.head.Next)

}

func TestFind_whenEmptyList_ReturnsNil(t *testing.T) {
	list := NonValidatedList[int]{}
	_, err := list.Find(1)
	require.Error(t, err)
}

func TestFind_whenValueExists_ReturnsIndex(t *testing.T) {
	list := NonValidatedList[int]{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	idx, err := list.Find(2)
	require.NoError(t, err)
	require.Equal(t, 1, idx)
}
