package list

func deleteMiddleNode(list *List) {

	if list == nil {
		return
	}

	left := list
	right := list

	for {
		if right.next == nil || right.next.next == nil {
			break
		}
		right = right.next.next
		left = left.next
	}

	left.next = left.next.next
}
