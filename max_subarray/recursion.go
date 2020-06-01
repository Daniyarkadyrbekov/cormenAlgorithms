package max_subarray

func recursionFind(arr []int) (int, int, int) {
	if len(arr) == 0 {
		return 0, 0, 0
	}
	if len(arr) == 1 {
		return 0, 0, arr[0]
	}

	mid := len(arr) / 2
	lli, lri, leftSum := recursionFind(arr[0:mid])
	rli, rri, rightSum := recursionFind(arr[mid : len(arr)-1])
	mli, mri, midSum := midSumFind(arr, mid)

	if leftSum > rightSum && leftSum > midSum {
		return lli, lri, leftSum
	}
	if rightSum > leftSum && rightSum > midSum {
		return rli, rri, rightSum
	}
	return mli, mri, midSum
}

func midSumFind(arr []int, mid int) (int, int, int) {
	lSum := 0
	lI := 0
	{
		sumI := 0
		for i := mid; i > 0; i-- {
			sumI += arr[i]
			if lSum < sumI {
				lSum = sumI
				lI = i
			}
		}
	}

	rSum := 0
	rI := 0
	{
		sumI := 0
		for i := mid + 1; i < len(arr); i++ {
			sumI += arr[i]
			if lSum < sumI {
				rSum = sumI
				rI = i
			}
		}
	}

	return lI, rI, lSum + rSum
}
