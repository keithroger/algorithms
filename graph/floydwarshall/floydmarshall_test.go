package floydwarshall

import (
	"math"
	"testing"
)

var INF = math.Inf(1)

func TestFloydMarshall(t *testing.T) {
	// Test input
	matrix := [][]float64{
		{0, 3, INF, 5},
		{2, 0, INF, 4},
		{INF, 1, 0, INF},
		{INF, INF, 2, 0},
	}

	want := [][]float64{
		{0, 3, 7, 5},
		{2, 0, 6, 4},
		{3, 1, 0, 5},
		{5, 3, 2, 0},
	}

	result := FloydMarshall(matrix)

	if !equal(want, result) {
		t.Fatalf("got: %v want: %v", result, want)
	}
}

func equal(a, b [][]float64) bool {
	equalityThreshold := 1e-9
	n := len(a)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if math.Abs(a[i][j]-b[i][j]) > equalityThreshold {

				return false
			}
		}
	}

	return true
}
