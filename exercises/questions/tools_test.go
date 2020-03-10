package questions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name     string
		array    []int
		expected []int
		lessFunc LessFunc
	}{
		{
			name:     "base",
			array:    []int{5, 8, 1, 3},
			expected: []int{1, 3, 5, 8},
			lessFunc: intComp,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			Quicksort(tt.array, tt.lessFunc)
			assert.Equal(t, tt.expected, tt.array)
		})
	}
}

func intComp(a, b interface{}) bool {
	return a.(int) < b.(int)
}
