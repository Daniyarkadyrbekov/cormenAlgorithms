package list

func RemoveDuplicatesWithTempMemory(list *List) *List {
	if list == nil {
		return nil
	}

	ptr := list.next
	prev := list
	duplicates := make(map[int]bool)
	duplicates[list.val] = true
	for {
		if ptr == nil {
			return list
		}
		if duplicates[ptr.val] {
			prev.next = ptr.next
		} else {
			duplicates[ptr.val] = true
			prev = ptr
		}
		ptr = ptr.next
	}
}
