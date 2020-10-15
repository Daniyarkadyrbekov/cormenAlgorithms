package max_subarray

import "math"

func linearFind(arr []int) (int, int, int) {
	var ri, li, i int
	sum, isum := math.MinInt32, math.MinInt32

	for j := 0; j < len(arr); j++ {
		if isum < 0 {
			i = j
			isum = 0
		}
		isum += arr[j]
		if isum > sum {
			sum = isum
			li = i
			ri = j
		}
	}

	return li, ri, sum
}
