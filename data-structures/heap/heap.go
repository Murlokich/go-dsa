package heap

import "cmp"

type Heap[T cmp.Ordered] struct {
	array []T
}

func BuildHeap[T cmp.Ordered](array ...T) *Heap[T] {
	heap := &Heap[T]{
		array: array,
	}
	for i := len(array) - 1; i >= 0; i-- {
		heap.moveDown(i)
	}
	return heap
}

func (h *Heap[T]) IsEmpty() bool {
	return len(h.array) == 0
}

func (h *Heap[T]) Min() (T, error) {
	var defaultValue T
	if h.IsEmpty() {
		return defaultValue, ErrEmptyHeap
	}
	return h.array[0], nil
}

func (h *Heap[T]) Insert(value T) {
	h.array = append(h.array, value)
	h.moveUp()
}

func (h *Heap[T]) ExtractMin() (T, error) {
	minValue, err := h.Min()
	if err != nil {
		return minValue, err
	}
	h.array[0], h.array[len(h.array)-1] = h.array[len(h.array)-1], h.array[0]
	// assign last value to default value to prevent potential data leakage
	var defaultValue T
	h.array[len(h.array)-1] = defaultValue
	h.array = h.array[:len(h.array)-1]
	h.moveDown(0)
	return minValue, nil
}

func (h *Heap[T]) moveUp() {
	curIndex := len(h.array) - 1
	parentIndex := h.getParentIndex(curIndex)
	for parentIndex >= 0 && h.array[curIndex] < h.array[parentIndex] {
		// swap parent and child
		h.array[curIndex], h.array[parentIndex] = h.array[parentIndex], h.array[curIndex]
		// move upwards
		curIndex = parentIndex
		parentIndex = h.getParentIndex(curIndex)
	}
}

func (h *Heap[T]) moveDown(index int) {
	curIndex := index
	leftChildIndex := h.getLeftChildIndex(curIndex)
	for leftChildIndex != -1 {
		minChildValue := h.array[leftChildIndex]
		minChildIndex := leftChildIndex
		rightChildIndex := h.getRightChildIndex(curIndex)
		if rightChildIndex != -1 && minChildValue > h.array[rightChildIndex] {
			minChildValue = h.array[rightChildIndex]
			minChildIndex = rightChildIndex
		}
		if h.array[curIndex] <= minChildValue {
			return
		}
		// swap with smaller child
		h.array[curIndex], h.array[minChildIndex] = h.array[minChildIndex], h.array[curIndex]
		// recalculate values for next iteration
		curIndex = minChildIndex
		leftChildIndex = h.getLeftChildIndex(curIndex)
	}
}

func (h *Heap[T]) getLeftChildIndex(index int) int {
	leftChildIndex := index*2 + 1
	if leftChildIndex >= len(h.array) {
		return -1
	}
	return leftChildIndex
}

func (h *Heap[T]) getRightChildIndex(index int) int {
	rightChildIndex := index*2 + 2
	if rightChildIndex >= len(h.array) {
		return -1
	}
	return rightChildIndex
}

func (h *Heap[T]) getParentIndex(index int) int {
	if index <= 0 {
		return -1
	}
	return (index - 1) / 2
}
