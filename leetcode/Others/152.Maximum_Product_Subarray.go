package leetcode

func maxProduct(nums []int) int {
	globalMax, tmpMax, tmpMin := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			tmpMax, tmpMin = tmpMin, tmpMax
		}
		tmpMax = max152(nums[i], tmpMax*nums[i])
		tmpMin = min152(nums[i], tmpMin*nums[i])
		globalMax = max152(globalMax, tmpMax)
	}

	return globalMax
}

func max152(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min152(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}