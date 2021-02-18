package list

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {

	var list *List
	require.Equal(t, "", list.String())

	list = list.Add(1)
	require.Equal(t, "1 ", list.String())
	list = list.Add(2)
	require.Equal(t, "2 1 ", list.String())
	list = list.Add(3)
	require.Equal(t, "3 2 1 ", list.String())

	list = list.Delete(2)
	require.Equal(t, "3 1 ", list.String())

	list = list.Add(3)
	list = list.Add(3)
	require.Equal(t, "3 3 3 1 ", list.String())

	list = list.Delete(3)
	require.Equal(t, "1 ", list.String())
}

func TestRemoveDuplicatesWithTempMemory(t *testing.T) {
	var list *List

	list = RemoveDuplicatesWithTempMemory(list)
	require.Equal(t, "", list.String())

	list = list.Add(1)
	list = list.Add(2)
	list = list.Add(3)
	list = list.Add(4)
	list = list.Add(1)

	list = RemoveDuplicatesWithTempMemory(list)
	require.Equal(t, "1 4 3 2 ", list.String())

}

func TestKthFromEnd(t *testing.T) {
	var list *List

	{
		res, exists := KthFromEnd(list, 0)
		require.False(t, exists)
		require.Equal(t, 0, res)
	}
	{
		res, exists := KthFromEnd(list, -1)
		require.False(t, exists)
		require.Equal(t, 0, res)
	}
	{
		res, exists := KthFromEnd(list, 100)
		require.False(t, exists)
		require.Equal(t, 0, res)
	}

	list = list.Add(0)
	list = list.Add(1)
	list = list.Add(2)
	list = list.Add(3)
	list = list.Add(4)
	list = list.Add(5)
	list = list.Add(6)

	{
		res, exists := KthFromEnd(list, 100)
		require.False(t, exists)
		require.Equal(t, 0, res)
	}
	{
		res, exists := KthFromEnd(list, 0)
		require.True(t, exists)
		require.Equal(t, 0, res)
	}
	{
		res, exists := KthFromEnd(list, 4)
		require.True(t, exists)
		require.Equal(t, 4, res)
	}
}

func TestDeleteMiddleNode(t *testing.T) {

	{
		var list *List
		list = deleteMiddleNode(list)
		require.Equal(t, "", list.String())
	}

	{
		var list *List
		list = list.Add(0)

		list = deleteMiddleNode(list)
		require.Equal(t, "", list.String())
	}

	{
		var list *List
		list = list.Add(0)
		list = list.Add(1)

		list = deleteMiddleNode(list)
		require.Equal(t, "0 ", list.String())
	}

	{
		var list *List
		list = list.Add(0)
		list = list.Add(1)
		list = list.Add(2)
		list = list.Add(3)
		list = list.Add(4)
		list = list.Add(5)

		list = deleteMiddleNode(list)
		require.Equal(t, "5 4 2 1 0 ", list.String())
	}

	{
		var list *List
		list = list.Add(0)
		list = list.Add(1)
		list = list.Add(2)
		list = list.Add(3)
		list = list.Add(4)
		list = list.Add(5)
		list = list.Add(6)

		list = deleteMiddleNode(list)
		require.Equal(t, "6 5 4 2 1 0 ", list.String())
	}

}

func TestListPartitions(t *testing.T) {

	{
		var list *List

		list = ListPartitions(list, 5)
		require.Equal(t, "", list.String())
	}

	{
		var list *List
		list = list.Add(12)
		list = list.Add(1)
		list = list.Add(20)
		list = list.Add(3)
		list = list.Add(4)
		list = list.Add(10)

		list = ListPartitions(list, 5)
		require.Equal(t, "4 3 1 12 20 10 ", list.String())
	}
}

func TestListPartitionsOptimized(t *testing.T) {

	{
		var list *List

		list = ListPartitionsOptimized(list, 5)
		require.Equal(t, "", list.String())
	}

	{
		var list *List
		list = list.Add(12)
		list = list.Add(1)
		list = list.Add(20)
		list = list.Add(3)
		list = list.Add(4)
		list = list.Add(10)
		list = list.Add(5)
		list = list.Add(5)
		list = list.Add(-7)
		list = list.Add(32)

		list = ListPartitionsOptimized(list, 5)
		require.Equal(t, "1 3 4 -7 12 20 10 5 5 32 ", list.String())
	}
}
