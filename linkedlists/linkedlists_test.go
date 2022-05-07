package linkedlists_test

import (
	"reflect"
	"testing"

	"github.com/keithroger/algorithms/linkedlists"
)

type testcase struct {
	items []int
	want  []int
}

func TestStack(t *testing.T) {
	tc := testcase{[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}}

	stack := linkedlists.Stack{}
	for _, item := range tc.items {
		stack.Push(item)
	}

	ans := make([]int, len(tc.items))
	for i := range tc.items {
		ans[i] = stack.Pop()
	}

	if !reflect.DeepEqual(ans, tc.want) {
		t.Errorf("Got: %v\nWant: %v", ans, tc.want)
	}
}

func TestQueue(t *testing.T) {
	tc := testcase{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}}

	queue := linkedlists.Queue{}
	for _, item := range tc.items {
		queue.Enqueue(item)
	}

	ans := make([]int, len(tc.items))
	for i := range tc.items {
		ans[i] = queue.Dequeue()
	}

	if !reflect.DeepEqual(ans, tc.want) {
		t.Errorf("Got: %v\nWant: %v", ans, tc.want)
	}
}
