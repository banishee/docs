package main

func findDuplicate(nums []int) int {
	l := 1
	r := len(nums) - 1

	for l < r {
		mid := (r + l) / 2
		count := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] <= mid {
				count++
			}
		}

		if count <= mid {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return r
}
