package lru

type Queue struct {
	head *QueueElem
	tail *QueueElem
}

type QueueElem struct {
	next  *QueueElem
	prev  *QueueElem
	value string
}

func (q *Queue) add(value string) {
	elem := &QueueElem{value: value}

	if q.tail == nil {
		q.head = elem
		q.tail = elem
	} else {
		q.head.prev = elem
		elem.next = q.head
		q.head = elem
	}
}

func (q *Queue) remove(value string) {
	if q.tail == nil {
		return
	}

	for elem := q.head; elem != nil; elem = elem.next {
		if elem.value != value {
			elem = elem.next
			continue
		}

		parent := elem.prev
		child := elem.next

		if parent != nil {
			parent.next = child
		} else {
			q.head = child
		}

		if child != nil {
			child.prev = parent
		} else {
			q.tail = parent
		}

		break
	}
}

func (q *Queue) removeFirst() *string {
	if q.head == nil {
		return nil
	}

	val := q.head.value
	child := q.head.next
	if child != nil {
		child.next = nil
		q.head = child
	}

	return &val
}

func (q *Queue) removeLast() *string {
	if q.tail == nil {
		return nil
	}

	val := q.tail.value
	parent := q.tail.prev
	if parent != nil {
		parent.next = nil
		q.tail = parent
	}

	return &val
}

func (q *Queue) makeHeadIfExist(value string) {
	for elem := q.head; elem != nil; elem = elem.next {
		if elem.value != value {
			elem = elem.next
			continue
		}

		parent := elem.prev
		child := elem.next

		if parent != nil {
			parent.next = child
		} else {
			q.head = child
		}

		if child != nil {
			child.prev = parent
		} else {
			q.tail = parent
		}

		elem.next = q.head
		q.head = elem

		break
	}
}
