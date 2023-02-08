package LeetCode

/*
Given an integer num, return the number of steps to reduce it to zero.

In one step, if the current number is even, you have to divide it by 2, otherwise, you have to subtract 1 from it.



Example 1:

Input: num = 14
Output: 6
Explanation:
Step 1) 14 is even; divide by 2 and obtain 7.
Step 2) 7 is odd; subtract 1 and obtain 6.
Step 3) 6 is even; divide by 2 and obtain 3.
Step 4) 3 is odd; subtract 1 and obtain 2.
Step 5) 2 is even; divide by 2 and obtain 1.
Step 6) 1 is odd; subtract 1 and obtain 0.

*/

func numberOfSteps1342(num int) int {
	// get how many step to reduce num to zero
	// if current number is even -> /2, else -1

	if num == 0 {
		return 0
	}

	counter := 1
	for {
		if num%2 == 0 {
			// even number
			num = num / 2
		} else if num%2 == 1 {
			// odd number
			num = num - 1
		}

		counter++

		if num == 1 || num == 0 {
			break
		}
	}

	return counter
}
