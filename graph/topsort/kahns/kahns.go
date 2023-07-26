package khans

type Queue []int

func (this *Queue) Push(x int) {
	*this = append(*this, x)
}

func (this *Queue) Pop() int {
	item := (*this)[0]
	*this = (*this)[1:]

	return item
}

func (this *Queue) Empty() bool {
	return len(*this) == 0
}

// Kahns does returns a topological sort for the given adjacency list.
func Kahns(adjList [][]int) []int {
	queue := Queue{}
	ordering := make([]int, 0, len(adjList))

	// Get dependency counts into deps.
	deps := make([]int, len(adjList))
	for i := range adjList {
		for _, j := range adjList[i] {
			deps[j]++
		}
	}

	// Add vertices with no dependencies.
	for i := range deps {
		if deps[i] == 0 {
			queue.Push(i)
		}
	}

	// Keep removing vertices with no dependencies.
	for !queue.Empty() {
		currVertex := queue.Pop()
		ordering = append(ordering, currVertex)

		// Decrement dependencies from neighboring vertices.
		for _, neighbor := range adjList[currVertex] {
			deps[neighbor]--

			// If the vertex now has no dependencies, add it to the queue.
			if deps[neighbor] == 0 {
				queue.Push(neighbor)
			}
		}
	}

	return ordering
}
