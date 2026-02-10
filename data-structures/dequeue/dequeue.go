package dequeue

type node[T any] struct {
	value T
	next  *node[T]
	prev  *node[T]
}

type Dequeue[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

func NewDequeue[T any](values ...T) *Dequeue[T] {
	dequeue := &Dequeue[T]{}
	for _, value := range values {
		dequeue.PushBack(value)
	}
	return dequeue
}

func (d *Dequeue[T]) Size() int {
	return d.size
}

func (d *Dequeue[T]) PushFront(value T) {
	defer func() {
		d.incSize()
	}()
	if d.Size() == 0 {
		d.head = &node[T]{value, nil, nil}
		d.tail = d.head
		return
	}
	newHead := &node[T]{value, d.head, nil}
	d.head.prev = newHead
	d.head = newHead
}

func (d *Dequeue[T]) PushBack(value T) {
	defer func() {
		d.incSize()
	}()
	if d.Size() == 0 {
		d.head = &node[T]{value, nil, nil}
		d.tail = d.head
		return
	}
	newTail := &node[T]{value, nil, d.tail}
	d.tail.next = newTail
	d.tail = newTail
}

func (d *Dequeue[T]) PopFront() (val T, err error) {
	defer func() {
		d.decSize(err)
	}()
	var defaultValue T
	if d.Size() == 0 {
		return defaultValue, ErrEmptyDequeue
	}
	head := d.head
	d.head = d.head.next
	if d.head != nil {
		d.head.prev = nil
	}
	return head.value, nil
}

func (d *Dequeue[T]) PopBack() (val T, err error) {
	defer func() {
		d.decSize(err)
	}()
	var defaultValue T
	if d.Size() == 0 {
		return defaultValue, ErrEmptyDequeue
	}
	tail := d.tail
	d.tail = d.tail.prev
	if d.tail != nil {
		d.tail.next = nil
	}
	return tail.value, nil
}

func (d *Dequeue[T]) incSize() {
	d.size++
}

func (d *Dequeue[T]) decSize(err error) {
	if err == nil {
		d.size--
	}
}
