package LeetCode

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {

	tests := []struct {
		name     string
		request  []int
		expected []int
	}{
		{
			name:     "runningSum1",
			request:  []int{1, 2, 3, 4},
			expected: []int{1, 3, 6, 10},
		},
		{
			name:     "runningSum2",
			request:  []int{1, 1, 1, 1, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "runningSum3",
			request:  []int{3, 1, 2, 10, 1},
			expected: []int{3, 4, 6, 16, 17},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expected, runningSum(test.request))
		})
	}

}
