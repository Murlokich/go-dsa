package algorithms

import (
	"reflect"
	"testing"
)

// For connected graphs, starting from a valid node should visit every node exactly once.
func TestDFS_ConnectedGraph_VisitsAllNodesExactlyOnce(t *testing.T) {
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

	d := NewDFS(adj, func(n int) {
		seen = append(seen, n)
		count[n]++
	})

	d.Run(0)

	if len(seen) != len(adj) {
		t.Fatalf("expected to visit %d nodes, got %d: %v", len(adj), len(seen), seen)
	}
	for i := range adj {
		if count[i] != 1 {
			t.Errorf("expected node %d to be visited exactly once, got %d times", i, count[i])
		}
	}
}

// DFS order should be deterministic given adjacency order.
func TestDFS_ConnectedGraph_VisitOrderMatchesAdjacencyOrder(t *testing.T) {
	// This graph + adjacency ordering makes the expected DFS preorder unambiguous.
	// 0 -> 1 -> 3, then back, 0 -> 2 -> 4
	adj := [][]int{
		{1, 2}, // 0
		{3},    // 1
		{4},    // 2
		{},     // 3
		{},     // 4
	}

	var got []int
	d := NewDFS(adj, func(n int) { got = append(got, n) })

	d.Run(0)

	want := []int{0, 1, 3, 2, 4}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("visit order mismatch\nwant: %v\ngot:  %v", want, got)
	}
}

// Cycles should not cause infinite recursion; nodes in a connected cyclic graph should still be visited once.
func TestDFS_ConnectedGraph_WithCycle(t *testing.T) {
	// Cycle: 0 -> 1 -> 2 -> 0 plus tail 2 -> 3 (still connected)
	adj := [][]int{
		{1},    // 0
		{2},    // 1
		{0, 3}, // 2
		{},     // 3
	}

	count := make([]int, len(adj))
	visitedOrder := make([]int, 0, len(adj))

	d := NewDFS(adj, func(n int) {
		count[n]++
		visitedOrder = append(visitedOrder, n)
	})

	d.Run(0)

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
func TestDFS_ConnectedGraph_WithSelfLoop(t *testing.T) {
	// 0 has a self-loop and connects to 1; 1 connects to 2; connected overall
	adj := [][]int{
		{0, 1}, // 0 self-loop then edge to 1
		{2},    // 1
		{},     // 2
	}

	count := make([]int, len(adj))
	d := NewDFS(adj, func(n int) { count[n]++ })

	d.Run(0)

	for i := range adj {
		if count[i] != 1 {
			t.Fatalf("expected node %d to be visited exactly once, got %d times", i, count[i])
		}
	}
}

// Starting from an out-of-range node should do nothing (guard clause in Run).
func TestDFS_Run_OutOfRangeStartNode_DoesNothing(t *testing.T) {
	adj := [][]int{
		{1},
		{0},
	}

	calls := 0
	d := NewDFS(adj, func(int) { calls++ })

	d.Run(999)

	if calls != 0 {
		t.Fatalf("expected visit not to be called, got %d calls", calls)
	}
}
