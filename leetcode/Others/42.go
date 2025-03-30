package main

func trap(height []int) (ans int) {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if height[left] < height[right] { // height[right] is the max in current state, so move the left
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right] // height[left] is the max in current state,
			right--
		}
	}
	return
}

// dp
// func trap(height []int) (ans int) {
//     n := len(height)
//     if n == 0 {
//         return
//     }

//     leftMax := make([]int, n)
//     leftMax[0] = height[0]
//     for i := 1; i < n; i++ {
//         leftMax[i] = max(leftMax[i-1], height[i])
//     }

//     rightMax := make([]int, n)
//     rightMax[n-1] = height[n-1]
//     for i := n - 2; i >= 0; i-- {
//         rightMax[i] = max(rightMax[i+1], height[i])
//     }

//     for i, h := range height {
//         ans += min(leftMax[i], rightMax[i]) - h
//     }
//     return
// }
