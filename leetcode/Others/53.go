package main

func maxSubArray(nums []int) int {
	// dp[i] represents the max subarray considering the nums[i] (might not be included)``
	// dp[i] = max(nums[i], dp[i-1] + nums[i])
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < dp[i-1]+nums[i] {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
	}

	max := 0
	for i := 0; i < len(dp); i++ {
		if max < dp[i] && dp[i] >= 0 {
			max = dp[i]
		}
	}
	return max
}

// func maxSubArray(nums []int) int {
// 	n := len(nums)
// 	if n <= 1 {
// 		return nums[0]
// 	}

// 	res := nums[0]
// 	dp := make([]int, n, n)
// 	dp[0] = nums[0]
// 	for i := 1; i < n; i++ {
// 		if dp[i-1]+nums[i] > nums[i] {
// 			dp[i] = dp[i-1] + nums[i]
// 		} else {
// 			dp[i] = nums[i]
// 		}

// 		if res < dp[i] {
// 			res = dp[i]
// 		}
// 	}

// 	return res
// }

// // 数组中至少有一个非负数
// func maxSubArray1(nums []int) int {
// 	maxSum, res, p := nums[0], 0, 0
// 	for p < len(nums) {
// 		res += nums[p]
// 		if res > maxSum {
// 			maxSum = res
// 		}
// 		if res < 0 {
// 			res = 0
// 		}
// 		p++
// 	}
// 	return maxSum
// }

// func main() {
// 	maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
// }
