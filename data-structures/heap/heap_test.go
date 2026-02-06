package heap

import (
	"cmp"
	"errors"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// assertMinHeapProperty verifies the internal heap invariant for a min-heap:
// for every node i, parent(i) <= i.
func assertMinHeapProperty[T cmp.Ordered](t *testing.T, h *Heap[T], less func(a, b T) bool) {
	t.Helper()
	for i := 1; i < len(h.array); i++ {
		p := (i - 1) / 2
		// parent must not be greater than child
		if less(h.array[i], h.array[p]) {
			t.Fatalf("min-heap property violated at index %d (parent %d): parent=%v child=%v heap=%v",
				i, p, h.array[p], h.array[i], h.array)
		}
	}
}

func extractAll[T cmp.Ordered](t *testing.T, h *Heap[T]) []T {
	t.Helper()
	out := make([]T, 0, len(h.array))
	for !h.IsEmpty() {
		v, err := h.ExtractMin()
		if err != nil {
			t.Fatalf("ExtractMin() unexpected error: %v", err)
		}
		out = append(out, v)
	}
	return out
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		wantEmpty bool
	}{
		{name: "empty", input: nil, wantEmpty: true},
		{name: "non_empty", input: []int{1}, wantEmpty: false},
		{name: "non_empty_many", input: []int{3, 2, 1}, wantEmpty: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := BuildHeap[int](tt.input...)
			if got := h.IsEmpty(); got != tt.wantEmpty {
				t.Errorf("IsEmpty() = %v, want %v (heap=%v)", got, tt.wantEmpty, h.array)
			}
		})
	}
}

func TestMin(t *testing.T) {
	t.Run("empty_returns_error", func(t *testing.T) {
		h := BuildHeap[int]()
		got, err := h.Min()
		if !errors.Is(err, ErrEmptyHeap) {
			t.Fatalf("Min() error = %v, want ErrEmptyHeap", err)
		}
		// default value for int should be 0
		if got != 0 {
			t.Errorf("Min() value = %v, want default 0 on error", got)
		}
	})

	tests := []struct {
		name    string
		input   []int
		wantMin int
	}{
		{name: "single", input: []int{42}, wantMin: 42},
		{name: "reverse_sorted", input: []int{9, 8, 7, 6, 5, 4, 3}, wantMin: 3},
		{name: "with_duplicates", input: []int{5, 1, 3, 1, 2, 2, 4}, wantMin: 1},
		{name: "with_negatives", input: []int{0, -10, 7, -3, 2, -10, 5}, wantMin: -10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := BuildHeap[int](tt.input...)
			got, err := h.Min()
			if err != nil {
				t.Fatalf("Min() unexpected error: %v", err)
			}
			if got != tt.wantMin {
				t.Errorf("Min() = %v, want %v (heap=%v)", got, tt.wantMin, h.array)
			}
			assertMinHeapProperty[int](t, h, func(a, b int) bool { return a < b })
		})
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		name    string
		inserts []int
		wantMin int
	}{
		{name: "increasing", inserts: []int{1, 2, 3, 4, 5}, wantMin: 1},
		{name: "decreasing", inserts: []int{5, 4, 3, 2, 1}, wantMin: 1},
		{name: "mixed_with_negatives", inserts: []int{10, -1, 7, -3, 2, 0}, wantMin: -3},
		{name: "duplicates", inserts: []int{2, 2, 2, 2}, wantMin: 2},
		{name: "single", inserts: []int{99}, wantMin: 99},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Heap[int]{}

			curMin := int(1e7)

			for i, v := range tt.inserts {
				h.Insert(v)

				// verify heap property after each insert
				assertMinHeapProperty[int](t, h, func(a, b int) bool { return a < b })

				// verify Min after each insert
				if v < curMin {
					curMin = v
				}
				gotMin, err := h.Min()
				if err != nil {
					t.Fatalf("after insert #%d (%v), Min() unexpected error: %v", i, v, err)
				}
				if gotMin != curMin {
					t.Fatalf("after insert #%d (%v), Min() = %v, want %v (heap=%v)",
						i, v, gotMin, curMin, h.array)
				}
			}

			gotMin, err := h.Min()
			if err != nil {
				t.Fatalf("final Min() unexpected error: %v", err)
			}
			if gotMin != tt.wantMin {
				t.Errorf("final Min() = %v, want %v (heap=%v)", gotMin, tt.wantMin, h.array)
			}
		})
	}
}

func TestExtractMin(t *testing.T) {
	t.Run("empty_returns_error", func(t *testing.T) {
		h := BuildHeap[int]()
		got, err := h.ExtractMin()
		if !errors.Is(err, ErrEmptyHeap) {
			t.Fatalf("ExtractMin() error = %v, want ErrEmptyHeap", err)
		}
		// default value for int should be 0
		if got != 0 {
			t.Errorf("ExtractMin() value = %v, want default 0 on error", got)
		}
	})

	tests := []struct {
		name  string
		input []int
	}{
		{name: "randomish", input: []int{5, 1, 9, 2, 7, 3, 8, 4, 6}},
		{name: "with_duplicates", input: []int{3, 1, 2, 1, 3, 2, 0, 0}},
		{name: "already_sorted", input: []int{1, 2, 3, 4, 5}},
		{name: "reverse_sorted", input: []int{5, 4, 3, 2, 1}},
		{name: "with_negatives", input: []int{-1, -5, 0, 2, -3, 1}},
		{name: "single", input: []int{42}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := BuildHeap[int](tt.input...)

			want := append([]int{}, tt.input...)
			sort.Ints(want)

			got := make([]int, 0, len(tt.input))
			for !h.IsEmpty() {
				v, err := h.ExtractMin()
				if err != nil {
					t.Fatalf("ExtractMin() unexpected error: %v", err)
				}
				got = append(got, v)

				// heap should still be valid after each extract (when non-empty)
				if !h.IsEmpty() {
					assertMinHeapProperty[int](t, h, func(a, b int) bool { return a < b })
				}
			}
			assert.Equal(t, want, got)
		})
	}
}

func TestBuildHeap(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{name: "empty", input: nil},
		{name: "small", input: []int{3, 1, 2}},
		{name: "with_duplicates", input: []int{2, 2, 1, 3, 1}},
		{name: "with_negatives", input: []int{-2, 0, -1, 5}},
		{name: "larger", input: []int{9, 1, 8, 2, 7, 3, 6, 4, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h1 := BuildHeap[int](tt.input...)
			h2 := &Heap[int]{}
			for _, v := range tt.input {
				h2.Insert(v)
			}

			out1 := extractAll[int](t, h1)
			out2 := extractAll[int](t, h2)

			if len(out1) != len(out2) {
				t.Fatalf("length mismatch: out1=%v out2=%v", out1, out2)
			}
			for i := range out1 {
				if out1[i] != out2[i] {
					t.Fatalf("mismatch at %d: out1=%v out2=%v", i, out1, out2)
				}
			}
		})
	}
}

func TestHeap_WorksWithStrings(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{name: "basic", input: []string{"d", "b", "a", "c"}, want: []string{"a", "b", "c", "d"}},
		{name: "with_duplicates", input: []string{"x", "a", "x", "b", "a"}, want: []string{"a", "a", "b", "x", "x"}},
		{name: "single", input: []string{"only"}, want: []string{"only"}},
		{name: "empty", input: []string{}, want: []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := BuildHeap[string](tt.input...)

			// Min() should error on empty; otherwise OK
			if len(tt.input) == 0 {
				_, err := h.Min()
				if !errors.Is(err, ErrEmptyHeap) {
					t.Fatalf("Min() error = %v, want ErrEmptyHeap", err)
				}
			} else {
				_, err := h.Min()
				if err != nil {
					t.Fatalf("Min() unexpected error: %v", err)
				}
			}

			got := extractAll[string](t, h)
			assert.Equal(t, tt.want, got)
		})
	}
}
