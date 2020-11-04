package array_hash_stringBuilder

type ArrayList struct {
	factor int
	curI   int
	arr    []int
}

func NewArrList(size, factor int) ArrayList {
	return ArrayList{
		curI:   -1,
		factor: factor,
		arr:    make([]int, size),
	}
}

func (a *ArrayList) Add(val int) {
	if len(a.arr) == a.curI+1 {
		a.arr = raise(a.arr, a.factor)
	}
	a.curI++
	a.arr[a.curI] = val
}

func (a *ArrayList) Pop() int {
	if len(a.arr) == 0 {
		return -1
	}

	val := a.arr[a.curI]
	a.curI -= 1
	return val
}

func raise(arr []int, factor int) []int {
	newSize := len(arr) * factor
	if newSize == 0 {
		newSize = factor
	}
	newArr := make([]int, newSize)
	for i, v := range arr {
		newArr[i] = v
	}

	return newArr
}
