package list

import "strconv"

type List struct {
	val  int
	next *List
}

func (l *List) Add(val int) *List {

	return &List{val: val, next: l}
}

func (l *List) Delete(val int) *List {
	head := l

	if head == nil {
		return nil
	}

	ptr := head
	for {
		if ptr.next == nil {
			break
		}
		if ptr.next.val == val {
			ptr.next = ptr.next.next
			continue
		}
		ptr = ptr.next
	}

	if head.val == val {
		head = head.next
	}

	return head
}

func (l *List) String() string {
	var res string
	ptr := l
	for {
		if ptr == nil {
			return res
		}
		res += strconv.Itoa(ptr.val) + " "
		ptr = ptr.next
	}
}
