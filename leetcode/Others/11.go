package main

func maxArea(height []int) int {
	res := 0
	for front, rear := 0, len(height)-1; front < rear; {
		area := min(height[front], height[rear]) * (rear - front)
		if res < area {
			res = area
		}

		if height[front] <= height[rear] {
			front++
		} else {
			rear--
		}
	}
	return res
}
