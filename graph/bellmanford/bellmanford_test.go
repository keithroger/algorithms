package bellmanford

import (
	"reflect"
	"testing"

	"github.com/keithroger/algorithms/graph"
	"github.com/stretchr/testify/assert"
)

func TestBellmanFord(t *testing.T) {
	want := map[string]float64{
		"A": 0,
		"B": 3,
		"C": -2,
		"D": -1,
		"E": 2,
	}

	graph := graph.NewGraph(5)
	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddVertex("C")
	graph.AddVertex("D")
	graph.AddVertex("E")
	graph.AddEdge("A", "B", 3)
	graph.AddEdge("A", "C", -2)
	graph.AddEdge("B", "D", 1)
	graph.AddEdge("C", "D", 1)
	graph.AddEdge("D", "E", 3)

	got, err := BellmanFord(graph, "A")
	if err != nil {
		t.Fatal(err)
	}

	if reflect.DeepEqual(want, got) {
		t.Fatalf("Got: %v not equal to want: %v", got, want)
	}
}

func TestBellmanFordNegative(t *testing.T) {
	graph := graph.NewGraph(5)
	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddVertex("C")
	graph.AddVertex("D")
	graph.AddVertex("E")
	graph.AddEdge("A", "B", 3)
	graph.AddEdge("A", "C", -2)
	graph.AddEdge("B", "D", 1)
	graph.AddEdge("C", "D", 1)
	graph.AddEdge("D", "E", 3)
	graph.AddEdge("C", "A", -1)

	got, err := BellmanFord(graph, "A")
	if err != nil {
		assert.EqualError(t, err, "Negative cycle detected.")
	}

	if got != nil {
		t.Fatal("Got not nil")
	}
}
