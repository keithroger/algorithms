package algorithms

import (
	"fmt"

	"github.com/keithroger/algorithms/linked"
)

func main() {
	items := []int{1, 2, 3, 4, 5, 56, 3, 2, 3, 4, 89, 3, 2, 4, 3, 2, 99, 34, 3, 0, 24, 12, 41, 31}

	queue := linked.Queue{}

	for item := range items {
		queue.Enqueue(item)
	}

	for queue.N != 0 {
		fmt.Println(queue.Dequeue())
	}

}
