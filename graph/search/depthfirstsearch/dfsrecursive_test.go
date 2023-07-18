package depthfirstsearch

import (
	"reflect"
	"testing"
)

func binaryTree() *Vertex {
	return &Vertex{
		Val: "A",
		Left: &Vertex{
			Val: "B",
			Left: &Vertex{
				Val: "D",
			},
			Right: &Vertex{
				Val: "E",
			},
		},
		Right: &Vertex{
			Val: "C",
			Left: &Vertex{
				Val: "F",
			},
			Right: &Vertex{
				Val: "G",
			},
		},
	}
}

func TestPreOrder(t *testing.T) {
	want := []string{"A", "B", "D", "E", "C", "F", "G"}
	got := PreOrder(binaryTree())

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Got: %v\t Want: %v", got, want)
	}
}

func TestInOrder(t *testing.T) {
	want := []string{"D", "B", "E", "A", "F", "C", "G"}
	got := InOrder(binaryTree())

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Got: %v\t Want: %v", got, want)
	}
}

func TestPostOrder(t *testing.T) {
	want := []string{"D", "E", "B", "F", "G", "C", "A"}
	got := PostOrder(binaryTree())

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Got: %v\t Want: %v", got, want)
	}
}
