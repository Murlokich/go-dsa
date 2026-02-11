package algorithms

type DFS struct {
	adjList [][]int
	visit   func(node int)
}

func NewDFS(adjList [][]int, visit func(node int)) *DFS {
	return &DFS{
		adjList: adjList,
		visit:   visit,
	}
}

func (d *DFS) Run(node int) error {
	visited := make([]bool, len(d.adjList))
	if node >= len(visited) || node < 0 {
		return ErrDFSInvalidStartNode
	}
	if visited[node] {
		return nil
	}
	visited[node] = true
	d.visit(node)
	for _, adjNode := range d.adjList[node] {
		if err := d.Run(adjNode); err != nil {
			return err
		}
	}
	return nil
}
