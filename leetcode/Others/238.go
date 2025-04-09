package main

/*
1   2        3           4
1 1*2    1*2*3         1*2*3*4

	    1*2 * 4        1*2*3

		      4*3          4
*/
func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))
	answer[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		answer[i] = answer[i-1] * nums[i]
	}
	// 1 2 3 4
	// 1 2 6 24
	//     2*4  6
	//

	answer[len(nums)-1] = answer[len(nums)-2]
	r := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 1; i-- {
		answer[i] = answer[i-1] * r
		r = r * nums[i]
	}
	answer[0] = r
	return answer
}

// func productExceptSelf(nums []int) []int {
// 	L, R := make([]int, len(nums)), make([]int, len(nums))
// 	L[0], R[len(nums)-1] = nums[0], nums[len(nums)-1]
// 	for i := 1; i < len(nums); i++ {
// 		L[i] = L[i-1] * nums[i]
// 		R[len(nums)-1-i] = R[len(nums)-i] * nums[len(nums)-1-i]
// 	}

// 	answer := make([]int, len(nums))
// 	answer[0], answer[len(nums)-1] = R[1], L[len(nums)-2]
// 	for i := 1; i < len(nums)-1; i++ {
// 		answer[i] = L[i-1] * R[i+1]
// 	}
// 	return answer
// }

func main() {
	productExceptSelf([]int{1, 2, 3, 4})
}
