package khans

import "testing"

func TestKahns(t *testing.T) {
	adjList := [][]int{
		0: {1, 2},
		1: {4},
		2: {3},
		3: {1},
		4: {},
	}

	ordering := Kahns(adjList)

	t.Log(ordering)
}
