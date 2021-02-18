package list

func deleteMiddleNode(head *List) *List {

	if head == nil || head.next == nil {
		return nil
	}

	left := head
	right := head.next

	if right.next == nil {
		return right
	}

	for {
		if right.next == nil || right.next.next == nil || right.next.next.next == nil {
			break
		}
		right = right.next.next
		left = left.next
	}

	left.next = left.next.next

	return head
}
