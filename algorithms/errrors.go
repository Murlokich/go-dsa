package algorithms

import "errors"

var (
	ErrBFSInvalidStartNode = errors.New("non-existent node provided to bfs")
	ErrDFSInvalidStartNode = errors.New("non-existent node provided to dfs")
)
