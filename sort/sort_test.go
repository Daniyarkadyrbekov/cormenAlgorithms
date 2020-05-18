package sort

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func insertionSort(arr []int) []int {

	for i, _ := range arr {
		for j := 0; j < i; j++ {
			if arr[i] < arr[j] {
				tmp := arr[i]
				for z := i; z > j; z-- {
					arr[z] = arr[z-1]
				}
				arr[j] = tmp
			}
		}
	}

	return arr
}

func recursiveSort(arr []int) []int {

	if len(arr) == 1 {
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

func TestMerge(t *testing.T) {
	arrl := []int{1, 4, 18}
	arrR := []int{2, 3, 5, 6, 7}
	require.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 18}, mergeSortedArr(arrl, arrR))
}

func TestSorts(t *testing.T) {
	arr := []int{1, 2, 2, 14, 1, 17, 19, -2, 0, 18, 120, 15}
	require.Equal(t, []int{-2, 0, 1, 1, 2, 2, 14, 15, 17, 18, 19, 120}, insertionSort(arr))
	require.Equal(t, []int{-2, 0, 1, 1, 2, 2, 14, 15, 17, 18, 19, 120}, recursiveSort(arr))
}
