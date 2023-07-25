package unionfind

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnionFind(t *testing.T) {
	adjacencyList := [][]int{
		// Group 1
		0: {1, 2},
		1: {0, 2},
		2: {0, 1},

		// Group 2
		3: {4, 5},
		4: {3, 5},
		5: {3, 4},
	}

	uf := NewUnionFind(6)

	for fromVertex, list := range adjacencyList {
		for _, toVertex := range list {
			uf.Union(fromVertex, toVertex)
		}
	}

	want := 2
	got := uf.GroupCount()

	if uf.GroupCount() != want {
		assert.Equal(t, want, got)
	}
}
