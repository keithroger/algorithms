package kruskals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKruskals(t *testing.T) {
	edges := []Edge{
		{From: 0, To: 1, Weight: 5},
		{From: 0, To: 2, Weight: 1},
		{From: 1, To: 3, Weight: 2},
		{From: 1, To: 4, Weight: 5},
		{From: 2, To: 3, Weight: 4},
	}
	mst, nGroups := Kruskals(edges)

	assert.Equal(t, 1, nGroups)

	t.Log(mst)

}
