package MockInterview

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMockInterview(t *testing.T) {
	tests := []struct {
		name     string
		request  []int
		expected int
	}{
		{
			name:     "find-second-max",
			request:  []int{1, 3, 6, 2, 4},
			expected: 4,
		},
		{
			name:     "find-second-max",
			request:  []int{3, 2, 4, 1, 5, 11},
			expected: 5,
		},
		{
			name:     "find-second-max",
			request:  []int{77, 32, 33, 55, 67},
			expected: 67,
		},
		{
			name:     "find-second-max",
			request:  []int{-12, -34, 43, 10, 9},
			expected: 10,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expected, findSecondMax(test.request))
		})
	}
}
