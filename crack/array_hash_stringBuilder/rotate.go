package array_hash_stringBuilder

func rotate(arr [][]int) [][]int {

	if len(arr) == 0 {
		return nil
	}

	arrLen := len(arr)
	for _, a := range arr {
		if len(a) != arrLen {
			return nil
		}
	}

	subArr := make([]int, arrLen-1)

	for step := 0; step < arrLen%2; step++ {
		N := arrLen
		//initialize
		for i := step; i < N-step-1; i++ {
			subArr[i] = arr[step][i]
		}

		//top
		for i := step; i < N-step-1; i++ {
			tmp := arr[i+step][N-step-1]
			arr[i+step][N-step-1] = subArr[i-step]
			subArr[i-step] = tmp
		}

		//right
		for i := N - step - 1; i > step; i-- {
			tmp := arr[N-1-step][i]
			arr[N-1-step][i] = subArr[N-i-1]
			subArr[N-i-1] = tmp
		}

		//bottom
		for i := N - 1; i > 0; i-- {
			tmp := arr[i][0]
			arr[i][0] = subArr[N-i-1]
			subArr[N-i-1] = tmp
		}

		//left
		for i := 0; i < N-1; i++ {
			tmp := arr[0][i]
			arr[0][i] = subArr[i]
			subArr[i] = tmp
		}

	}

	return arr
}
