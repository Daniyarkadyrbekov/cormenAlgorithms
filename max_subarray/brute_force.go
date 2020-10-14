package max_subarray

import "math"

func bruteForce(arr []int) (left, right, sum int) {
	sum = math.MinInt32

	for i := 0; i < len(arr); i++ {
		lsum := 0
		for j := i; j < len(arr); j++ {
			lsum += arr[j]
			if lsum > sum {
				sum = lsum
				left = i
				right = j
			}
		}
	}

	return
}
