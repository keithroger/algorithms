package main

import "fmt"
// Note: using a slice will result in underlying array having the max needed
// memory although the MaxHeap may be small

type MaxHeap []int

func (h *MaxHeap) length() int {
    return len(*h)
}

// create constructor for new heap from array

func (heap *MaxHeap) Push(val int) {
    *heap = append(*heap, val)
    heap.swim()
}

func (heap *MaxHeap) Pop() int{
    top := (*heap)[0]

    lastIdx := heap.length()-1
    heap.swap(0, lastIdx)
    *heap = (*heap)[:lastIdx]

    heap.sink()
    // TODO fix sink

    return top
}

func (heap *MaxHeap) swim() {
    curr := heap.length() - 1
    for curr > 0  {
        parent := (curr - 1) / 2
        if (*heap)[parent] < (*heap)[curr] {
            heap.swap(parent, curr)
        }
        curr = parent
    }
}

func (heap *MaxHeap) sink() {
    curr := 0
    largest := curr
    for curr < heap.length() {
        left, right := curr*2 + 1, curr*2 + 2

        if right < heap.length() && (*heap)[right] > (*heap)[largest] {
            largest = right
        }
        if left < heap.length() && (*heap)[left] > (*heap)[largest] {
            largest = left
        }
        if largest == curr {
            break
        }

        heap.swap(largest, curr)
        curr = largest
    }
}

func (heap *MaxHeap) swap(a, b int) {
    (*heap)[a], (*heap)[b] = (*heap)[b], (*heap)[a]
}

func main() {
    data := []int{1, 3, 2, 4, 5, 6}

    heap := MaxHeap{}
    for _, datum := range data {
        heap.Push(datum)
    }

    fmt.Println("Heap: ", heap)

    fmt.Println("Pop:", heap.Pop())

    fmt.Println("Heap: ", heap)

    fmt.Println("Insert: 7")
    heap.Push(7)
    fmt.Println("Heap: ", heap)
}
