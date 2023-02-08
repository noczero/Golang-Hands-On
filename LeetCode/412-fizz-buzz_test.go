package LeetCode

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFizzBuzz(t *testing.T) {

	tests := []struct {
		name     string
		request  int
		expected []string
	}{
		{
			name:     "fizzbuzz",
			request:  3,
			expected: []string{"1", "2", "Fizz"},
		}, {
			name:     "fizzbuzz",
			request:  5,
			expected: []string{"1", "2", "Fizz", "4", "Buzz"},
		}, {
			name:     "fizzbuzz",
			request:  15,
			expected: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := fizzBuzz412(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}
