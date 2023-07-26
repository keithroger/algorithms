package priorityqueue

// The following demonstrates how a priority queue using a min-heap.

type PriorityQueue []int

// NewPriorityQueueWithArr returns a new PriorityQueue initialized with the
// from arr.
func NewPriorityQueueWithArr(arr []int) PriorityQueue {
	// Convert arr into a priority queue
	pq := PriorityQueue(arr)

	n := len(pq)
	for idx := n/2 - 1; idx >= 0; idx-- {
		pq.sink(idx)
	}

	return pq
}

// NewPriorityQueueWithLen return a new uninitialized PriorityQueue with
// length n.
func NewPriorityQueueWithLen(n int) PriorityQueue {
	return make(PriorityQueue, n)
}

// Pop returns the smallest value in the array and removes it from the array.
func (pq *PriorityQueue) Pop() int {
	// Get min value from the heap
	minValue := (*pq)[0]

	// Swap the first and last value in the heap
	pq.swap(0, len(*pq)-1)

	// Truncate the array
	(*pq) = (*pq)[:len(*pq)-1]

	// Sink the current top value to restore the heap
	pq.sink(0)

	return minValue
}

// Push adds val to the heap
func (pq *PriorityQueue) Push(val int) {
	// Add val as the next node in the heap
	*pq = append(*pq, val)

	// Move new value into position to restore the heap
	pq.swim()
}

// swim puts the last entry in the priority queue into the correct position.
func (pq *PriorityQueue) swim() {
	currIdx := len(*pq) - 1
	for currIdx > 0 {
		parentIdx := (currIdx - 1) / 2

		// If parent value is greater, then break.
		if (*pq)[currIdx] > (*pq)[parentIdx] {
			break
		}

		// If parent has a large value, then swap.
		if (*pq)[parentIdx] > (*pq)[currIdx] {
			pq.swap(currIdx, parentIdx)
		}

		currIdx = parentIdx
	}
}

// sink moves the value lower in the tree until it reaches the correct position.
func (pq *PriorityQueue) sink(idx int) {
	currIdx := idx
	largest := idx
	for currIdx < len(*pq) {
		// If right child is greater than current node, then swap values.
		rightChild := idx*2 + 2
		if rightChild < len(*pq) && (*pq)[rightChild] < (*pq)[currIdx] {
			largest = rightChild
		}

		// If left child is greater than current node, then swap values.
		leftChild := rightChild - 1
		if leftChild < len(*pq) && (*pq)[leftChild] < (*pq)[currIdx] {
			largest = leftChild
		}

		// Break early if finished sinking.
		if largest == currIdx {
			break
		}

		pq.swap(largest, currIdx)
		currIdx = largest
	}
}

// swap takes the values at idx1 and idx2 and swaps their position.
func (pq *PriorityQueue) swap(idx1, idx2 int) {
	(*pq)[idx1], (*pq)[idx2] = (*pq)[idx2], (*pq)[idx1]
}
