package dfs

// DFS uses an adjacency list to create a topological ordering.
func DFS(adjList [][]int) []int {
	visited := make([]bool, len(adjList))
	ordering := make([]int, len(adjList))
	orderIdx := new(int)
	*orderIdx = len(adjList) - 1

	// Iterate through all vertices in the graph.
	for v := range adjList {
		dfs(v, adjList, &visited, &ordering, orderIdx)
	}

	return ordering
}

// dfs is used to run though the neighbors of node v.
func dfs(v int, adjList [][]int, visited *[]bool, ordering *[]int, orderIdx *int) {
	// Return early if already visited.
	if (*visited)[v] {
		return
	}
	(*visited)[v] = true

	// Visit all neighboring vertices with dfs.
	for _, neighbor := range adjList[v] {
		dfs(neighbor, adjList, visited, ordering, orderIdx)

	}

	// Add vertices to ordering as dfs unwinds.
	(*ordering)[*orderIdx] = v
	*orderIdx--
}
