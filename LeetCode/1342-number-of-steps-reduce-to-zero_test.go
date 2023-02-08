package LeetCode

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNumberOfStepsReducetoZero(t *testing.T) {

	tests := []struct {
		name     string
		request  int
		expected int
	}{
		{
			name:     "NumberOfStepsReducetoZero",
			request:  14,
			expected: 6,
		}, {
			name:     "NumberOfStepsReducetoZero",
			request:  8,
			expected: 4,
		}, {
			name:     "NumberOfStepsReducetoZero",
			request:  123,
			expected: 12,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := numberOfSteps1342(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}
