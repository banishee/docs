package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i:=0; i < len(nums); i++ {
		complement := target - nums[i]
		if k, ok := m[complement]; ok {
			return []int{k, i}
		}
		m[nums[i]] = i
	}
	return nil
}
