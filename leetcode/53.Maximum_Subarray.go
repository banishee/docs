package leetcode

func maxSubArray(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return nums[0]
	}

	res := nums[0]
	dp := make([]int, n , n)
	dp[0]= nums[0]
	for i :=1; i < n; i++ {
		if dp[i-1]+nums[i] > dp[i-1] {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}

		if res < dp[i] {
			res = dp[i]
		}
	}

	return res
}

// 数组中至少有一个非负数
func maxSubArray1(nums []int) int {
	maxSum, res, p := nums[0], 0, 0
	for p < len(nums) {
		res += nums[p]
		if res > maxSum {
			maxSum = res
		}
		if res < 0 {
			res = 0
		}
		p++
	}
	return maxSum
}