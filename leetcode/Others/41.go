package main

func firstMissingPositive(nums []int) int {
	tmp := 0
	for i := 0; i < len(nums); i++ {
		if 0 <= nums[i] && nums[i] <= len(nums) && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	return tmp
}
