package bellmanford

import (
	"errors"
	"math"

	"github.com/keithroger/algorithms/graph"
)

func BellmanFord(g graph.Graph, startingVertex string) (map[string]int, error) {
	// Keep track of the distance to for each vertex.
	distances := make(map[string]int, len(g.Vertices))

	// Set all distances to infinity.
	for label := range g.Vertices {
		distances[label] = math.MaxInt
	}

	// Set distance for the starting node to 0.
	distances[startingVertex] = 0

	// Loop over all edges.
	for i := 0; i < len(g.Vertices)-1; i++ {
		for _, edge := range g.Edges {
			// If current edge gives a shorter path record the distance.
			if distances[edge.From.Label]+edge.Weight < distances[edge.To.Label] {
				distances[edge.To.Label] = distances[edge.From.Label] + edge.Weight
			}
		}
	}

	// Iterate over vertices again to check for negative cycles. (Optional)
	for i := 0; i < len(g.Vertices)-1; i++ {
		for _, edge := range g.Edges {
			// If negative cycle found return error
			if distances[edge.From.Label]+edge.Weight < distances[edge.To.Label] {
				return nil, errors.New("Negative cycle detected.")
			}
		}
	}

	return distances, nil
}
