package main

import "fmt"
// Note: using a slice will result in underlying array having the max needed
// memory although the MaxHeap may be small

type MaxHeap []int

func (h *MaxHeap) length() int {
    return len(*h)
}

func NewHeap(arr []int) MaxHeap {
    heap := MaxHeap(arr)

    n := len(arr)
    for i := n/2 - 1; i >= 0; i-- {
        heap.sink(i)
    }

    return heap
}

func (heap *MaxHeap) Push(val int) {
    *heap = append(*heap, val)
    heap.swim()
}

func (heap *MaxHeap) Pop() int{
    top := (*heap)[0]

    lastIdx := heap.length()-1
    heap.swap(0, lastIdx)
    *heap = (*heap)[:lastIdx]

    heap.sink(0)

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

func (heap *MaxHeap) sink(idx int) {
    curr := idx
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
    fmt.Println("Starting Data: ", data)

    heap := NewHeap(data)

    fmt.Println("New Heap: ", heap)

    fmt.Println("Pop:", heap.Pop())

    fmt.Println("Heap: ", heap)

    fmt.Println("Push: 7")
    heap.Push(7)
    fmt.Println("Heap: ", heap)

    fmt.Print("Pop All: " )
    for heap.length() > 0 {
        fmt.Print(heap.Pop())
    }
    fmt.Println()
}
