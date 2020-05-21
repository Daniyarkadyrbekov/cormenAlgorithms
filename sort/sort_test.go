package sort

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func insertionSort(arr []int) []int {

	for i := 1; i < len(arr); i++ {
		tmp := arr[i]
		j := 0
		for j = i; j > 0; j-- {
			if tmp < arr[j-1] {
				arr[j] = arr[j-1]
			} else {
				break
			}
		}
		arr[j] = tmp
	}

	return arr
}

func recursiveSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}
	splitter := len(arr) / 2
	return mergeSortedArr(recursiveSort(arr[:splitter]), recursiveSort(arr[splitter:]))
}

func mergeSortedArr(left, right []int) []int {

	result := make([]int, 0, len(left)+len(right))

	i, j := 0, 0
	for z := 0; z < len(left)+len(right); z++ {
		if left[i] < right[j] {
			result = append(result, left[i])
			if i == len(left)-1 {
				result = append(result, right[j:]...)
				break
			} else {
				i++
			}
		} else {
			result = append(result, right[j])
			if j == len(right)-1 {
				result = append(result, left[i:]...)
				break
			}
			if j < len(right)-1 {
				j++
			}
		}
	}

	return result
}

func selectSortExtraArr(arr []int) []int {

	sorted := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		min := math.MaxInt32
		for j := i; j < len(arr); j++ {
			if arr[j] < min {
				min = arr[j]
			}
		}
		sorted[i] = min
	}

	return sorted
}

func selectSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		min := math.MaxInt32
		minI := min
		for j := i; j < len(arr); j++ {
			if arr[j] < min {
				min = arr[j]
				minI = j
			}
		}

		arr[minI] = arr[i]
		arr[i] = min
	}

	return arr
}

func TestMerge(t *testing.T) {
	arrl := []int{1, 4, 18}
	arrR := []int{2, 3, 5, 6, 7}
	require.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 18}, mergeSortedArr(arrl, arrR))
}

func TestSorts(t *testing.T) {

	for i, arr := range []struct {
		unsorted []int
		sorted   []int
	}{
		{
			unsorted: []int{},
			sorted:   []int{},
		},
		{
			unsorted: []int{1},
			sorted:   []int{1},
		},
		{
			unsorted: []int{-2, 0, 1, 1, 2, 2, 14, 15, 17, 18, 19, 120},
			sorted:   []int{-2, 0, 1, 1, 2, 2, 14, 15, 17, 18, 19, 120},
		},
		{
			unsorted: []int{120, 19, 18, 17, 15, 14, 2, 2, 1, 1, 0, -2},
			sorted:   []int{-2, 0, 1, 1, 2, 2, 14, 15, 17, 18, 19, 120},
		},
		{
			unsorted: []int{1, 2, 2, 14, 1, 17, 19, -2, 0, 18, 120, 15},
			sorted:   []int{-2, 0, 1, 1, 2, 2, 14, 15, 17, 18, 19, 120},
		},
	} {
		fmt.Printf("i = %d\n", i)
		require.Equal(t, arr.sorted, insertionSort(arr.unsorted))
		require.Equal(t, arr.sorted, recursiveSort(arr.unsorted))
		require.Equal(t, arr.sorted, selectSortExtraArr(arr.unsorted))
		require.Equal(t, arr.sorted, selectSort(arr.unsorted))
	}
}
