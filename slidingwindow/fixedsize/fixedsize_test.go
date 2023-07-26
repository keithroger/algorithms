package fixedsize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFixedSize(t *testing.T) {
	// Setup
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	size := 2
	want := 19

	// Action
	got := fixedSize(arr, size)

	// Assert
	assert.Equal(t, want, got)
}
