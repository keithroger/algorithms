package graph

// Graph uses Vertices to represent a graph.
type Graph struct {
	Vertices map[string]*Vertex
	Edges    []*Edge
}

// NewGraph creates a new Graph of size n.
func NewGraph(n int) Graph {
	return Graph{
		Vertices: make(map[string]*Vertex, n),
	}
}

// AddVertex adds a vertex to the graph with a key
func (g *Graph) AddVertex(label string) {
	g.Vertices[label] = &Vertex{Label: label}
}

// AddEdge adds an edge to the graph from src to dest with a weight.
func (g *Graph) AddEdge(src, dst string, weight int) {
	srcVertex := g.Vertices[src]
	dstVertex := g.Vertices[dst]
	edge := &Edge{
		Weight: weight,
		From:   srcVertex,
		To:     dstVertex,
	}
	srcVertex.Edges = append(srcVertex.Edges, edge)
	g.Edges = append(g.Edges, edge)
}

// Vertex stores a value and edges that connect to other vertices.
type Vertex struct {
	Label string
	Edges []*Edge
}

// Edge stores weight and a connection to another vertex in the graph.
type Edge struct {
	Weight int
	From   *Vertex
	To     *Vertex
}
