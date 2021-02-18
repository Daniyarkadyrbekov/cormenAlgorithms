package list

func ListPartitions(head *List, partitions int) *List {
	var left, right, last *List
	ptr := head
	for {
		if ptr == nil {
			break
		}
		if ptr.val < partitions {
			if last == nil {
				left = ptr
				last = left
			} else {
				last.next = ptr
				last = ptr
			}
		} else {
			right = right.Add(ptr.val)
		}
		ptr = ptr.next
	}
	if last != nil {
		last.next = right
	}
	return left
}

func ListPartitionsOptimized(head *List, partitions int) *List {

	var left, lastLeft, right *List
	var tmpPtr *List
	ptr := head

	for {
		if ptr == nil {
			break
		}

		tmpPtr = ptr
		ptr = ptr.next

		if tmpPtr.val < partitions {
			if left == nil {
				lastLeft = tmpPtr
			}
			tmpPtr.next = left
			left = tmpPtr
		} else {
			tmpPtr.next = right
			right = tmpPtr
		}
	}

	if lastLeft == nil {
		left = right
	} else {
		lastLeft.next = right
	}

	return left
}
