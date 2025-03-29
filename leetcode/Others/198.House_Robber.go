package main

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	preMax, curMax := 0, nums[0]
	for i := 1; i < len(nums); i++ {
		tmp := curMax
		if preMax+nums[i] > curMax {
			curMax = preMax + nums[i]
		}
		preMax = tmp
	}

	return curMax
}
