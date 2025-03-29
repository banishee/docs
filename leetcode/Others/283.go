package main

func moveZeroes(nums []int) {
	zero_counts := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zero_counts++
			continue
		}

		nums[i-zero_counts] = nums[i]
	}

	for i := 1; i <= zero_counts; i++ {
		index := len(nums) - i
		nums[index] = 0
	}
}
