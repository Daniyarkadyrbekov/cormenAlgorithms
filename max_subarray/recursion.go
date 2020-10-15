package max_subarray

import "math"

func recursionFind(arr []int) (int, int, int) {

	if len(arr) == 0 {
		return 0, 0, math.MinInt32
	}

	if len(arr) == 1 {
		return 0, 0, arr[0]
	}
	mid := len(arr) / 2

	left1, right1, sum1 := recursionFind(arr[0:mid])

	left2, right2, sum2 := recursionFind(arr[mid:])
	left2 += mid
	right2 += mid

	left3, right3, sum3 := findMaxThroughMid(arr)

	if sum3 >= sum1 && sum3 >= sum2 {
		return left3, right3, sum3
	}

	if sum2 >= sum1 && sum2 >= sum3 {
		return left2, right2, sum2
	}

	return left1, right1, sum1

}

func findMaxThroughMid(arr []int) (int, int, int) {
	mid := len(arr) / 2

	sumL, sumR := math.MinInt32, math.MinInt32
	left, right := mid, mid

	{
		//left
		var localSum int
		for i := mid; i >= 0; i-- {
			localSum += arr[i]
			if sumL < localSum {
				sumL = localSum
				left = i
			}
		}
	}

	{
		//right
		var localSum int
		for i := mid + 1; i < len(arr); i++ {
			localSum += arr[i]
			if sumR < localSum {
				sumR = localSum
				right = i
			}
		}
	}

	return left, right, sumL + sumR
}
