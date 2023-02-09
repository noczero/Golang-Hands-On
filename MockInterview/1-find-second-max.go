package MockInterview

func findSecondMax(slice []int) int {
	firstMax := 0
	secondMax := 0

	for _, value := range slice {
		// [3, 2, 4, 1, 5, 11]
		// value : 3
		// max1 : 0->3
		// max2 : 0
		// --
		// value : 2
		// max1 : 3
		// max2 : 0->2
		// --
		// value : 4
		// max1 : 3->4
		// max2 : 2->3
		// --
		// value : 1
		// max1 : 4
		// max2 : 3
		// --
		// value : 5
		// max1 : 4->5
		// max2 : 3->4
		// --
		// value : 11
		// max1 : 5->11
		// max2 : 4->5

		if value > firstMax {
			secondMax = firstMax
			firstMax = value
		}

		if value > secondMax && value < firstMax {
			secondMax = value
		}
	}

	return secondMax
}

//func main() {
//	mySlice := []int{3, 2, 4, 1, 5, 11}
//
//	fmt.Println(findMax(mySlice))
//}
