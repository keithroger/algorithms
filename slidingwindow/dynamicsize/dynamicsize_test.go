package dynamicsize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDynamicSize(t *testing.T) {
	// Setup
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	goal := 19
	want := 2

	// Action
	got := DynamicSize(arr, goal)

	// Assert
	assert.Equal(t, want, got)
}
