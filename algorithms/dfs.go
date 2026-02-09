package algorithms

type DFS struct {
	adjList [][]int
	visit   func(node int)
	visited []bool
}

func NewDFS(adjList [][]int, visit func(node int)) *DFS {
	return &DFS{
		adjList: adjList,
		visit:   visit,
		visited: make([]bool, len(adjList)),
	}
}

func (d *DFS) Run(node int) {
	if node >= len(d.visited) {
		return
	}
	if d.visited[node] {
		return
	}
	d.visited[node] = true
	d.visit(node)
	for _, adjNode := range d.adjList[node] {
		d.Run(adjNode)
	}
}
