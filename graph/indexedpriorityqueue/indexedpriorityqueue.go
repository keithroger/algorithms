// This package implements a indexed priority queue using a min heap
package indexedpriorityqueue

import (
	"fmt"
	"strconv"
)

type IndexedPriorityQueue struct {
	heap     []*vertex
	keyToPos map[string]int
}

// NewIndexedPriorityQueueWithLen returns a new IndexedPriorityQueue with length n.
func NewIndexedPriorityQueueWithLen(n int) *IndexedPriorityQueue {
	return &IndexedPriorityQueue{
		heap:     make([]*vertex, 0, n),
		keyToPos: make(map[string]int, n),
	}
}

func (ipq *IndexedPriorityQueue) NewPriorityQueueWithArr() {
}

// Pop removes the element with the smallest value and the value is returned.
func (ipq *IndexedPriorityQueue) Pop() int {
	topElement := ipq.heap[0]

	ipq.swap(0, len(ipq.heap)-1)
	delete(ipq.keyToPos, topElement.key)
	ipq.heap = ipq.heap[:len(ipq.heap)-1]
	ipq.sink(0)

	return topElement.val
}

// Pop removes an element from the IndexedPriorityQueue using the key and returns the value.
func (ipq *IndexedPriorityQueue) PopByKey(key string) int {
	idx := ipq.keyToPos[key]

	result := ipq.heap[idx].val

	// Swap the element for removal with the last element.
	ipq.swap(idx, len(ipq.heap)-1)
	idx = ipq.keyToPos[key] // update index since swap moved the position.

	// Remove the last element in the heap and map.
	ipq.heap = ipq.heap[:len(ipq.heap)-1]
	delete(ipq.keyToPos, key)

	// If the key being removed is the last element, skip sink/swim.
	if idx == len(ipq.heap) {
		return result
	}

	ipq.swim(idx)
	ipq.sink(idx)

	return result
}

// Push adds a new element to the IndexedPriorityQueue.
func (ipq *IndexedPriorityQueue) Push(key string, val int) {
	v := &vertex{
		val: val,
		key: key,
	}

	ipq.keyToPos[key] = len(ipq.heap)
	ipq.heap = append(ipq.heap, v)

	ipq.swim(len(ipq.heap) - 1)
}

func (ipq *IndexedPriorityQueue) sink(idx int) {
	currIdx := idx

	for leftChild := currIdx*2 + 1; currIdx < len(ipq.heap); leftChild = currIdx * 2 {
		if leftChild < len(ipq.heap) && ipq.less(leftChild, currIdx) {
			ipq.swap(leftChild, currIdx)
			currIdx = leftChild
			continue
		}

		rightChild := leftChild + 1
		if rightChild < len(ipq.heap) && ipq.less(rightChild, currIdx) {
			ipq.swap(leftChild, currIdx)
			currIdx = rightChild
			continue
		}

		break

	}
}

func (ipq *IndexedPriorityQueue) swim(idx int) {
	currIdx := idx

	for parentIdx := (currIdx - 1) / 2; currIdx > 0 && ipq.less(currIdx, parentIdx); parentIdx = (currIdx - 1) / 2 {
		ipq.swap(parentIdx, currIdx)

		currIdx = parentIdx
	}
}

func (ipq *IndexedPriorityQueue) less(idxA, idxB int) bool {
	return ipq.heap[idxA].val < ipq.heap[idxB].val
}

func (ipq *IndexedPriorityQueue) swap(idxA, idxB int) {
	// Swap position of the vertices in the heap.
	ipq.heap[idxA], ipq.heap[idxB] = ipq.heap[idxB], ipq.heap[idxA]
	// Fix the nameToPos map to show the updated positions of the elements.
	ipq.keyToPos[ipq.heap[idxA].key], ipq.keyToPos[ipq.heap[idxB].key] = idxA, idxB
}

// String return a printable string to represent the priority queue.
func (ipq IndexedPriorityQueue) String() string {
	str := "Heap: \n"
	for idx, vertex := range ipq.heap {
		str += "\t" + strconv.Itoa(idx) + " value: " + strconv.Itoa(vertex.val) + " "
		str += "\tname: " + vertex.key + "\n"
	}
	str += fmt.Sprintf("NameToPos: %v\n", ipq.keyToPos)
	return str
}

type vertex struct {
	val int
	key string // key used in nameToPos map
}
