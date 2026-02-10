package dequeue

import (
	"errors"
	"testing"
)

func assertSize[T any](t *testing.T, d *Dequeue[T], want int) {
	t.Helper()
	if got := d.Size(); got != want {
		t.Fatalf("size mismatch: want %d, got %d", want, got)
	}
}

func assertErrIs(t *testing.T, err error, want error) {
	t.Helper()
	if !errors.Is(err, want) {
		t.Fatalf("error mismatch: want %v, got %v", want, err)
	}
}

func TestNewDequeue_Empty(t *testing.T) {
	d := NewDequeue[int]()
	assertSize(t, d, 0)
}

func TestNewDequeue_WithValues_PreservesOrderOnPopFront(t *testing.T) {
	d := NewDequeue[int](1, 2, 3, 4)
	assertSize(t, d, 4)

	for _, want := range []int{1, 2, 3, 4} {
		got, err := d.PopFront()
		if err != nil {
			t.Fatalf("PopFront unexpected error: %v", err)
		}
		if got != want {
			t.Errorf("PopFront value mismatch: want %d, got %d", want, got)
		}
	}
	assertSize(t, d, 0)
}

func TestNewDequeue_WithValues_PreservesOrderOnPopBack(t *testing.T) {
	d := NewDequeue[int](1, 2, 3, 4)
	assertSize(t, d, 4)

	for _, want := range []int{4, 3, 2, 1} {
		got, err := d.PopBack()
		if err != nil {
			t.Fatalf("PopBack unexpected error: %v", err)
		}
		if got != want {
			t.Errorf("PopBack value mismatch: want %d, got %d", want, got)
		}
	}
	assertSize(t, d, 0)
}

func TestPushFront_ThenPopFront_LIFOForFront(t *testing.T) {
	var d Dequeue[int]
	d.PushFront(1)
	d.PushFront(2)
	d.PushFront(3)
	assertSize(t, &d, 3)

	for _, want := range []int{3, 2, 1} {
		got, err := d.PopFront()
		if err != nil {
			t.Fatalf("PopFront unexpected error: %v", err)
		}
		if got != want {
			t.Errorf("PopFront value mismatch: want %d, got %d", want, got)
		}
	}
	assertSize(t, &d, 0)
}

func TestPushBack_ThenPopBack_LIFOForBack(t *testing.T) {
	var d Dequeue[int]
	d.PushBack(1)
	d.PushBack(2)
	d.PushBack(3)
	assertSize(t, &d, 3)

	for _, want := range []int{3, 2, 1} {
		got, err := d.PopBack()
		if err != nil {
			t.Fatalf("PopBack unexpected error: %v", err)
		}
		if got != want {
			t.Errorf("PopBack value mismatch: want %d, got %d", want, got)
		}
	}
	assertSize(t, &d, 0)
}

func TestPushFront_ThenPopBack_FIFOFromFrontToBack(t *testing.T) {
	// PushFront builds [3,2,1]; PopBack returns 1 then 2 then 3.
	var d Dequeue[int]
	d.PushFront(1)
	d.PushFront(2)
	d.PushFront(3)
	assertSize(t, &d, 3)

	for _, want := range []int{1, 2, 3} {
		got, err := d.PopBack()
		if err != nil {
			t.Fatalf("PopBack unexpected error: %v", err)
		}
		if got != want {
			t.Errorf("PopBack value mismatch: want %d, got %d", want, got)
		}
	}
	assertSize(t, &d, 0)
}

func TestPushBack_ThenPopFront_FIFOFromBackToFront(t *testing.T) {
	// PushBack builds [1,2,3]; PopFront returns 1 then 2 then 3.
	var d Dequeue[int]
	d.PushBack(1)
	d.PushBack(2)
	d.PushBack(3)
	assertSize(t, &d, 3)

	for _, want := range []int{1, 2, 3} {
		got, err := d.PopFront()
		if err != nil {
			t.Fatalf("PopFront unexpected error: %v", err)
		}
		if got != want {
			t.Errorf("PopFront value mismatch: want %d, got %d", want, got)
		}
	}
	assertSize(t, &d, 0)
}

func TestMixedOperations_Sequence(t *testing.T) {
	// Build by mixing:
	// start []
	// PushBack(2)   -> [2]
	// PushFront(1)  -> [1,2]
	// PushBack(3)   -> [1,2,3]
	// PushFront(0)  -> [0,1,2,3]
	// PopFront -> 0 -> [1,2,3]
	// PopBack  -> 3 -> [1,2]
	// PopFront -> 1 -> [2]
	// PopBack  -> 2 -> []
	var d Dequeue[int]

	d.PushBack(2)
	d.PushFront(1)
	d.PushBack(3)
	d.PushFront(0)
	assertSize(t, &d, 4)

	v, err := d.PopFront()
	if err != nil {
		t.Fatalf("PopFront unexpected error: %v", err)
	}
	if v != 0 {
		t.Errorf("PopFront mismatch: want %d, got %d", 0, v)
	}
	assertSize(t, &d, 3)

	v, err = d.PopBack()
	if err != nil {
		t.Fatalf("PopBack unexpected error: %v", err)
	}
	if v != 3 {
		t.Errorf("PopBack mismatch: want %d, got %d", 3, v)
	}
	assertSize(t, &d, 2)

	v, err = d.PopFront()
	if err != nil {
		t.Fatalf("PopFront unexpected error: %v", err)
	}
	if v != 1 {
		t.Errorf("PopFront mismatch: want %d, got %d", 1, v)
	}
	assertSize(t, &d, 1)

	v, err = d.PopBack()
	if err != nil {
		t.Fatalf("PopBack unexpected error: %v", err)
	}
	if v != 2 {
		t.Errorf("PopBack mismatch: want %d, got %d", 2, v)
	}
	assertSize(t, &d, 0)
}

func TestPopFront_Empty_ReturnsErrAndDoesNotChangeSize(t *testing.T) {
	d := NewDequeue[int]()
	assertSize(t, d, 0)

	_, err := d.PopFront()
	assertErrIs(t, err, ErrEmptyDequeue)
	assertSize(t, d, 0)
}

func TestPopBack_Empty_ReturnsErrAndDoesNotChangeSize(t *testing.T) {
	d := NewDequeue[int]()
	assertSize(t, d, 0)

	_, err := d.PopBack()
	assertErrIs(t, err, ErrEmptyDequeue)
	assertSize(t, d, 0)
}

func TestSize_TracksOperations(t *testing.T) {
	var d Dequeue[int]
	assertSize(t, &d, 0)

	d.PushFront(10)
	assertSize(t, &d, 1)

	d.PushBack(20)
	assertSize(t, &d, 2)

	_, err := d.PopFront()
	if err != nil {
		t.Fatalf("PopFront unexpected error: %v", err)
	}
	assertSize(t, &d, 1)

	_, err = d.PopBack()
	if err != nil {
		t.Fatalf("PopBack unexpected error: %v", err)
	}
	assertSize(t, &d, 0)
}

func TestGenericType_String(t *testing.T) {
	d := NewDequeue[string]("a", "b")
	assertSize(t, d, 2)

	d.PushFront("z") // [z,a,b]
	assertSize(t, d, 3)

	v, err := d.PopBack() // b
	if err != nil {
		t.Fatalf("PopBack unexpected error: %v", err)
	}
	if v != "b" {
		t.Fatalf("PopBack mismatch: want %q, got %q", "b", v)
	}
	assertSize(t, d, 2)

	v, err = d.PopFront() // z
	if err != nil {
		t.Fatalf("PopFront unexpected error: %v", err)
	}
	if v != "z" {
		t.Fatalf("PopFront mismatch: want %q, got %q", "z", v)
	}
	assertSize(t, d, 1)
}
