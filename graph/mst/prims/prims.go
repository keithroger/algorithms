package prims

import (
	"container/heap"
)

type Edge struct {
	From, To int
	Weight   int
}

type minPQ []Edge

func (mpq minPQ) Len() int {
	return len(mpq)
}

func (mpq minPQ) Less(i, j int) bool {
	return (mpq)[i].Weight < (mpq)[j].Weight
}

func (mpq minPQ) Swap(i, j int) {
	(mpq)[i], (mpq)[j] = (mpq)[j], (mpq)[i]
}

func (mpq *minPQ) Push(x any) {
	*mpq = append(*mpq, x.(Edge))
}

func (mpq *minPQ) Pop() any {
	minItem := (*mpq)[(mpq).Len()-1]
	(*mpq) = (*mpq)[:(mpq).Len()-1]

	return minItem
}

func (mpq *minPQ) Empty() bool {
	return len(*mpq) == 0
}

// AddEdges adds all edges for a given vertex to the priority queue.
func (mpq *minPQ) AddEdges(g Graph, vertexIndex int, visited []bool) {
	visited[vertexIndex] = true

	for _, edge := range g[vertexIndex] {
		if !visited[edge.To] {
			heap.Push(mpq, edge)
		}
	}
}

type Graph [][]Edge

func Prims(g Graph) []Edge {
	mst := make([]Edge, 0, len(g)-1)
	mpq := minPQ{}
	visited := make([]bool, len(g))

	// Add first edge.
	mpq.AddEdges(g, 0, visited)

	for !mpq.Empty() {
		edge := heap.Pop(&mpq).(Edge)

		if visited[edge.To] {
			continue
		}

		// Add edges for the next closest vertices.
		mpq.AddEdges(g, edge.To, visited)

		mst = append(mst, edge)
	}

	return mst

}
