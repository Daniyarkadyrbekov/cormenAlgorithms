package max_subarray

func bruteForce(arr []int) (int, int, int) {
	maxSum := 0
	maxRI := 0
	maxLI := 0
	for i, _ := range arr {
		leftI := i
		rightI := 0
		sumI := arr[i]
		for j := i + 1; j < len(arr); j++ {
			if sumI < sumI+arr[j] {
				sumI = sumI + arr[j]
				rightI = j
			}
		}
		if sumI > maxSum {
			maxLI = leftI
			maxRI = rightI
			maxSum = sumI
		}
	}

	return maxLI, maxRI, maxSum
}
