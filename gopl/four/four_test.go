package four

import (
	"reflect"
	"testing"
)

func TestDedupAdjacent(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "base",
			input:    []string{"a", "a", "b"},
			expected: []string{"a", "b"},
		}, {
			name:     "multiple",
			input:    []string{"a", "a", "a", "b"},
			expected: []string{"a", "b"},
		}, {
			name:     "end",
			input:    []string{"a", "b", "b"},
			expected: []string{"a", "b"},
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			res := deduplicateAdjacent(tt.input)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("input slice %v not changed to look like expected %v", res, tt.expected)
			}
		})
	}
}

func TestUnic(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{
			name:     "base",
			input:    []byte("ba   se"),
			expected: []byte("ba se"),
		}, {
			name:     "multiple",
			input:    []byte("ba   se     bb"),
			expected: []byte("ba se bb"),
		}, {
			name:     "start",
			input:    []byte("   ba   se"),
			expected: []byte(" ba se"),
		}, {
			name:     "end",
			input:    []byte("ba   se  "),
			expected: []byte("ba se "),
		}, {
			name:     "no change",
			input:    []byte("ba se"),
			expected: []byte("ba se"),
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			res := squashSpace(tt.input)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("input slice %v not changed to look like expected %v", res, tt.expected)
			}
		})
	}
}
