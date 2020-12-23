package list

func KthFromEnd(list *List, k int) (res int, exists bool) {

	if list == nil {
		return
	}

	leftPtr := list
	rightPtr := list

	for i := 0; i < k; i++ {
		if rightPtr == nil {
			return
		}
		rightPtr = rightPtr.next
	}

	exists = true
	for {
		if rightPtr.next == nil {
			res = leftPtr.val
			return
		}
		rightPtr = rightPtr.next
		leftPtr = leftPtr.next
	}
}
