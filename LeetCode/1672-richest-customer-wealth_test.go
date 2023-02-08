package LeetCode

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRichestCustomerWealth(t *testing.T) {

	tests := []struct {
		name     string
		request  [][]int
		expected int
	}{
		{
			name:     "richetscustomerwealth",
			request:  [][]int{{1, 2, 3}, {3, 2, 1}},
			expected: 6,
		}, {
			name:     "richetscustomerwealth",
			request:  [][]int{{1, 5}, {7, 3}, {3, 5}},
			expected: 10,
		}, {
			name:     "richetscustomerwealth",
			request:  [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}},
			expected: 17,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expected, maximumWealth1672(test.request))
		})
	}
}
