package max_subarray

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

type checkType struct {
	arr []int
	lI  int
	rI  int
	sum int
}

func TestMaxSubArray(t *testing.T) {
	for i, arr := range []checkType{
		{
			arr: []int{},
			rI:  0,
			lI:  0,
			sum: math.MinInt32,
		},
		{
			arr: []int{-2, -5, -15, -1},
			rI:  3,
			lI:  3,
			sum: -1,
		},
		{
			arr: []int{2, 5, 15, 1},
			rI:  3,
			lI:  0,
			sum: 23,
		},
		{
			arr: []int{2, 5, -15, 1, 3, 7},
			rI:  5,
			lI:  3,
			sum: 11,
		},
		{
			arr: []int{2, 5, -2, 1, 3},
			rI:  4,
			lI:  0,
			sum: 9,
		},
		{
			arr: []int{-2, 5, -2, 1, 3},
			rI:  4,
			lI:  1,
			sum: 7,
		},
	} {
		fmt.Printf("i = %d\n", i)
		checkMethod(t, arr, bruteForce)
		checkMethod(t, arr, recursionFind)
	}
}

func checkMethod(t *testing.T, arr checkType, method func([]int) (int, int, int)) {
	lI, rI, sum := method(arr.arr)
	require.Equal(t, arr, checkType{
		arr: arr.arr,
		lI:  lI,
		rI:  rI,
		sum: sum})
}
