package prims

import "testing"

func TestPrims(t *testing.T) {
	graph := Graph{
		0: []Edge{
			{To: 1, From: 0, Weight: 3},
			{To: 2, From: 0, Weight: 1},
		},
		1: []Edge{
			{To: 3, From: 1, Weight: 2},
			{To: 4, From: 1, Weight: 5},
		},
		2: []Edge{
			{To: 0, From: 2, Weight: 1},
			{To: 3, From: 2, Weight: 4},
		},
		3: []Edge{
			{To: 1, From: 3, Weight: 2},
			{To: 2, From: 3, Weight: 4},
		},
		4: []Edge{
			{To: 1, From: 4, Weight: 5},
		},
	}

	t.Log(graph)

	t.Log(Prims(graph))

}
