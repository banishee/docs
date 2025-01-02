package leetcode

func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	preMax1, curMax1 := 0, nums[0]
	for i := 1; i < len(nums)-1; i++ {
		tmp := curMax1
		if preMax1+ nums[i] > curMax1 {
			curMax1 = preMax1 + nums[i]
		}
		preMax1 = tmp
	}

	preMax2, curMax2 := 0, nums[1]
	for i := 2; i < len(nums); i++ {
		tmp := curMax2
		if preMax2+ nums[i] > curMax2 {
			curMax2 = preMax2 + nums[i]
		}
		preMax2 = tmp
	}

	if curMax1 > curMax2{
		return curMax1
	}
	return curMax2
}
