package main

func jump(nums []int) int {
	count := 0
	maxIndex := 0
	endIndex := 0
	for i := 0; i < len(nums)-1; i++ {
		maxIndex = max(maxIndex, i+nums[i])
		if i == endIndex {
			endIndex = maxIndex // get the next end index
			count++
		}
	}

	return count
}
