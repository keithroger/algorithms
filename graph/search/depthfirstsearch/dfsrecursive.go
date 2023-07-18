package depthfirstsearch

type Vertex struct {
	Val   string
	Left  *Vertex
	Right *Vertex
}

// PreOrder returns a string representation of the order of traversal.
func PreOrder(v *Vertex) []string {
	result := make([]string, 0)
	preOrder(v, &result)

	return result
}

func preOrder(v *Vertex, result *[]string) {
	if v == nil {
		return
	}

	*result = append(*result, v.Val)
	preOrder(v.Left, result)
	preOrder(v.Right, result)
}

// InOrder returns a string representation of the order of traversal. The order of the vertices should represent
// how the binary tree values appear from left to right.
func InOrder(v *Vertex) []string {
	result := make([]string, 0)
	inOrder(v, &result)

	return result
}

func inOrder(v *Vertex, result *[]string) {
	if v == nil {
		return
	}

	inOrder(v.Left, result)
	*result = append(*result, v.Val)
	inOrder(v.Right, result)
}

// PreOrder returns a string representation of the order of traversal.
func PostOrder(v *Vertex) []string {
	result := make([]string, 0)
	postOrder(v, &result)

	return result
}

func postOrder(v *Vertex, result *[]string) {
	if v == nil {
		return
	}

	postOrder(v.Left, result)
	postOrder(v.Right, result)
	*result = append(*result, v.Val)
}
