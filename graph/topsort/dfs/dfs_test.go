package dfs

import "testing"

func TestDFS(t *testing.T) {
	adjList := [][]int{
		0: {1, 2},
		1: {4},
		2: {3},
		3: {1},
		4: {},
	}

	ordering := DFS(adjList)

	t.Log(ordering)
}
