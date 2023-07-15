package floydwarshall

// FloydMarshall finds the shortest path between all vertices.
// A value for matrix[i][j] contains the weight of the edge from i to j.
func FloydMarshall(matrix [][]float64) [][]float64 {
	n := len(matrix)

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if matrix[i][k]+matrix[k][j] < matrix[i][j] {
					matrix[i][j] = matrix[i][k] + matrix[k][j]
				}
			}
		}
	}

	return matrix
}
