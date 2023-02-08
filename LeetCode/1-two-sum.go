package LeetCode

import "fmt"

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

func twoSum1(nums []int, target int) []int {
	// bruteforce O(n^2)
	//for i, val1 := range nums {
	//	for j, val2 := range nums {
	//		if val1+val2 == target && i != j {
	//			return []int{i, j}
	//		}
	//	}
	//}

	// use map key value O(n)
	resultMap := make(map[int]int)

	for index, value := range nums {
		_, exist := resultMap[value] // store value as key in map, _, exist is bool return false if not exist, return true is exist
		fmt.Println("index", index)
		fmt.Println("value", value)
		fmt.Println("exist", exist)
		fmt.Println("resultMap", resultMap)
		if exist {
			return []int{resultMap[value], index}
		}

		// compliment to the target
		resultMap[target-value] = index // save value to index
	}

	return nil
}
