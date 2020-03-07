package questions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuestionOne(t *testing.T) {
	tests := []struct {
		name     string
		x        int64
		y        int64
		expected int64
	}{
		{
			name:     "simple",
			x:        1,
			y:        1,
			expected: 2,
		}, {
			name:     "larger",
			x:        759,
			y:        674,
			expected: 1433,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := QuestionOne(tt.x, tt.y)
			assert.Equal(t, tt.expected, result)
		})
	}
}
