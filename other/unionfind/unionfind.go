package unionfind

type UnionFind struct {
	// parent node of each vertex
	parent []int

	// size of a vertex
	size []int

	// nGroups counts how many groups
	nGroups int
}

// NewUnionFind creates a UnionFind of size n
func NewUnionFind(n int) UnionFind {
	unionFind := UnionFind{
		parent:  make([]int, n),
		size:    make([]int, n),
		nGroups: n,
	}

	// Initialize all vertices to be their own parent
	for i := range unionFind.parent {
		unionFind.parent[i] = i
		unionFind.size[i] = 1
	}

	return unionFind
}

// Returns the number of groups.
func (this *UnionFind) GroupCount() int {
	return this.nGroups
}

// Find swims up the tree until the last parent is found. The index of the last parent is returned.
func (this *UnionFind) Find(v int) int {
	// Find the top of the tree.
	newParent := v
	for this.parent[newParent] != newParent {
		newParent = this.parent[newParent]
	}

	// Set all visited nodes to point to the same parent.
	// This compresses the path.
	curr := v
	for this.parent[curr] != curr {
		next := this.parent[curr]
		this.parent[curr] = newParent
		curr = next
	}

	return newParent
}

// Union joins the trees of the vertices a and b.
func (this *UnionFind) Union(a, b int) {
	rootA, rootB := this.Find(a), this.Find(b)

	// If they are part of the same tree exit early.
	if rootA == rootB {
		return
	}

	// Merge with the smaller group with the larger group.
	if this.size[rootA] < this.size[rootB] {
		this.parent[rootA] = rootB
		this.size[rootA] += this.size[rootB]
	} else {
		this.parent[rootB] = rootA
		this.size[rootA] += this.size[rootA]
	}

	this.nGroups--
}
