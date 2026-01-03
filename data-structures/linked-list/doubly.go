package linked_list

type doublyLinkedNode[T comparable] struct {
	value T
	next  *doublyLinkedNode[T]
	prev  *doublyLinkedNode[T]
}

type DoublyLinkedList[T comparable] struct {
	head *doublyLinkedNode[T]
	tail *doublyLinkedNode[T]
}

func NewDoublyLinkedList[T comparable](values ...T) *DoublyLinkedList[T] {
	doublyLinkedList := &DoublyLinkedList[T]{}
	for _, value := range values {
		doublyLinkedList.InsertAtTail(value)
	}
	return doublyLinkedList
}

func (l *DoublyLinkedList[T]) InsertAtHead(value T) {
	newHead := &doublyLinkedNode[T]{value: value, next: l.head}
	if l.IsEmpty() {
		l.tail = newHead
	} else {
		l.head.prev = newHead
	}
	l.head = newHead
}

func (l *DoublyLinkedList[T]) InsertAtTail(value T) {
	newTail := &doublyLinkedNode[T]{value: value, prev: l.tail}
	if l.IsEmpty() {
		l.tail = newTail
		l.head = newTail
		return
	}
	l.tail.next = newTail
	l.tail = newTail
}

func (l *DoublyLinkedList[T]) DeleteHead() (err error) {
	if l.IsEmpty() {
		return ErrEmptyList
	}
	l.head = l.head.next
	if l.head == nil {
		l.tail = nil
	} else {
		l.head.prev = nil
	}
	return nil
}

func (l *DoublyLinkedList[T]) DeleteTail() (err error) {
	if l.IsEmpty() {
		return ErrEmptyList
	}
	l.tail = l.tail.prev
	if l.tail == nil {
		l.head = nil
	} else {
		l.tail.next = nil
	}
	return nil
}

func (l *DoublyLinkedList[T]) DeleteValue(value T) (err error) {
	if l.IsEmpty() {
		return ErrNoSuchValue
	}
	if l.head.value == value {
		l.head = l.head.next
		if l.head == nil {
			l.tail = nil
		} else {
			l.head.prev = nil
		}
		return nil
	}
	var prev *doublyLinkedNode[T]
	for prev = l.head; prev.next != nil && prev.next.value != value; prev = prev.next {
	}
	if prev.next == nil {
		return ErrNoSuchValue
	}
	deleted := prev.next
	prev.next = deleted.next
	if prev.next != nil {
		prev.next.prev = prev
	}
	if deleted == l.tail {
		l.tail = prev
	}
	return nil
}

func (l *DoublyLinkedList[T]) IsEmpty() bool {
	return l.head == nil
}

func (l *DoublyLinkedList[T]) GetTail() (T, error) {
	var defaultVal T
	if l.IsEmpty() {
		return defaultVal, ErrEmptyList
	}
	return l.tail.value, nil
}

func (l *DoublyLinkedList[T]) GetHead() (T, error) {
	var defaultVal T
	if l.IsEmpty() {
		return defaultVal, ErrEmptyList
	}
	return l.head.value, nil
}
