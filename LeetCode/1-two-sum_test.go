package LeetCode

import (
	"github.com/stretchr/testify/require"
	"testing"
)

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]
*/

func TestTwoSum(t *testing.T) {

	tests := []struct {
		name     string
		request1 []int
		request2 int
		expected []int
	}{
		{
			name:     "TestTwoSum1",
			request1: []int{2, 7, 11, 15},
			request2: 9,
			expected: []int{0, 1},
		}, {
			name:     "TestTwoSum1",
			request1: []int{3, 2, 4},
			request2: 6,
			expected: []int{1, 2},
		}, {
			name:     "TestTwoSum1",
			request1: []int{3, 3},
			request2: 6,
			expected: []int{0, 1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expected, twoSum1(test.request1, test.request2))
		})
	}

}
