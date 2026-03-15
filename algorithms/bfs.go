package algorithms

import (
	linked_list "github.com/Murlokich/go-dsa/data-structures/linked-list"
	"github.com/pkg/errors"
)

type BFS struct {
	adjList [][]int
	visit   func(node int)
}

func NewBFS(adjList [][]int, visit func(node int)) *BFS {
	return &BFS{
		adjList: adjList,
		visit:   visit,
	}
}

// Run
// Time Complexity: O(V + E) - we iterate across all the nodes and all the edges
// Space complexity: O(V) - queue
func (d *BFS) Run(startNode int) error {
	if startNode >= len(d.adjList) || startNode < 0 {
		return ErrBFSInvalidStartNode
	}
	visited := make([]bool, len(d.adjList))
	queue := linked_list.NewDoublyLinkedList(startNode)
	visited[startNode] = true
	for !queue.IsEmpty() {
		value, err := queue.GetHead()
		if err != nil {
			return errors.Wrap(err, "queue get head")
		}
		if err = queue.DeleteHead(); err != nil {
			return errors.Wrap(err, "queue delete head")
		}
		d.visit(value)
		for _, neighbour := range d.adjList[value] {
			if !visited[neighbour] {
				queue.InsertAtTail(neighbour)
				visited[neighbour] = true
			}
		}
	}
	return nil
}
