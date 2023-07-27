package kadanes

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKadanes(t *testing.T) {
	tt := []struct {
		arr  []int
		want int
	}{
		{[]int{-5, -4, -3, -2, -1}, -1},
		{[]int{1, 2, 3, -1, 1, 2, 3, -10, 1, 2, 3, 4}, 11},
		{[]int{0, 0}, 0},
	}

	for i, tc := range tt {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := Kadanes(tc.arr)
			assert.Equal(t, tc.want, got)
		})
	}
}
