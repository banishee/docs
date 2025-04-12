package main

func canJump(nums []int) bool {
	maxDistance := nums[0]
	for i := 1; i < len(nums); i++ {
		if maxDistance < i {
			return false
		}
		if maxDistance < i+nums[i] {
			maxDistance = i + nums[i]
		}
		if maxDistance >= len(nums)-1 {
			return true
		}
	}
	return true
}
