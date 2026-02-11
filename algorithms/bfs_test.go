package algorithms

import (
	"errors"
	"reflect"
	"testing"
)

// For connected graphs, starting from a valid node should visit every node exactly once.
func TestBFS_ConnectedGraph_VisitsAllNodesExactlyOnce(t *testing.T) {
	// Connected undirected-ish graph (modeled with symmetric edges)
	// 0-1, 0-2, 1-3, 2-4
	adj := [][]int{
		{1, 2}, // 0
		{0, 3}, // 1
		{0, 4}, // 2
		{1},    // 3
		{2},    // 4
	}

	seen := make([]int, 0, len(adj))
	count := make([]int, len(adj))

	b := NewBFS(adj, func(n int) {
		seen = append(seen, n)
		count[n]++
	})

	if err := b.Run(0); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(seen) != len(adj) {
		t.Fatalf("expected to visit %d nodes, got %d: %v", len(adj), len(seen), seen)
	}
	for i := range adj {
		if count[i] != 1 {
			t.Errorf("expected node %d to be visited exactly once, got %d times", i, count[i])
		}
	}
}

// BFS order should be deterministic given adjacency order.
func TestBFS_ConnectedGraph_VisitOrderMatchesAdjacencyOrder(t *testing.T) {
	// This graph + adjacency ordering makes the expected BFS order unambiguous.
	// 0 enqueues 1,2; then 1 enqueues 3; then 2 enqueues 4
	adj := [][]int{
		{1, 2}, // 0
		{3},    // 1
		{4},    // 2
		{},     // 3
		{},     // 4
	}

	var got []int
	b := NewBFS(adj, func(n int) { got = append(got, n) })

	if err := b.Run(0); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := []int{0, 1, 2, 3, 4}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("visit order mismatch\nwant: %v\ngot:  %v", want, got)
	}
}

// Cycles should not cause infinite loops; nodes in a connected cyclic graph should still be visited once.
func TestBFS_ConnectedGraph_WithCycle(t *testing.T) {
	// Cycle: 0 -> 1 -> 2 -> 0 plus tail 2 -> 3 (still connected)
	adj := [][]int{
		{1},    // 0
		{2},    // 1
		{0, 3}, // 2
		{},     // 3
	}

	count := make([]int, len(adj))
	visitedOrder := make([]int, 0, len(adj))

	b := NewBFS(adj, func(n int) {
		count[n]++
		visitedOrder = append(visitedOrder, n)
	})

	if err := b.Run(0); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(visitedOrder) != len(adj) {
		t.Fatalf("expected to visit %d nodes, got %d: %v", len(adj), len(visitedOrder), visitedOrder)
	}
	for i := range adj {
		if count[i] != 1 {
			t.Errorf("expected node %d to be visited exactly once, got %d times", i, count[i])
		}
	}
}

// Self-loop should not break traversal; still visits each node once in a connected graph.
func TestBFS_ConnectedGraph_WithSelfLoop(t *testing.T) {
	// 0 has a self-loop and connects to 1; 1 connects to 2; connected overall
	adj := [][]int{
		{0, 1}, // 0 self-loop then edge to 1
		{2},    // 1
		{},     // 2
	}

	count := make([]int, len(adj))
	b := NewBFS(adj, func(n int) { count[n]++ })

	if err := b.Run(0); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	for i := range adj {
		if count[i] != 1 {
			t.Fatalf("expected node %d to be visited exactly once, got %d times", i, count[i])
		}
	}
}

// Starting from an out-of-range node should return ErrBFSInvalidStartNode and not call visit.
func TestBFS_Run_OutOfRangeStartNode_ReturnsErrorAndDoesNothing(t *testing.T) {
	adj := [][]int{
		{1},
		{0},
	}

	calls := 0
	b := NewBFS(adj, func(int) { calls++ })

	err := b.Run(999)

	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if !errors.Is(err, ErrBFSInvalidStartNode) {
		t.Fatalf("expected ErrBFSInvalidStartNode, got %v", err)
	}
	if calls != 0 {
		t.Fatalf("expected visit not to be called, got %d calls", calls)
	}
}
