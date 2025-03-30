package main

import "sort"

func threeSum(nums []int) [][]int {
	// increasing
	sort.Ints(nums)

	ans := [][]int{}
	// negative, 0, positive
	for i := 0; i < len(nums); i++ {
		if 0 < nums[i] { // if target > 0, the next two numbers must be larger than 0
			break
		}
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		front, rear := i+1, len(nums)-1
		for front < rear {
			sum := nums[i] + nums[front] + nums[rear]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[front], nums[rear]})
				// find next front
				for front = front + 1; front < rear; front++ {
					if nums[front-1] != nums[front] {
						break
					}
				}
				// find next rear
				for rear = rear - 1; rear > front; rear-- {
					if nums[rear] != nums[rear+1] {
						break
					}
				}
			} else if sum > 0 {
				// find next rear
				for rear = rear - 1; rear > front; rear-- {
					if nums[rear] != nums[rear+1] {
						break
					}
				}
			} else if sum < 0 {
				// find next front
				for front = front + 1; front < rear; front++ {
					if nums[front-1] != nums[front] {
						break
					}
				}
			}
		}
	}

	return ans
}

func main() {
	//threeSum([]int{2, -3, 0, -2, -5, -5, -4, 1, 2, -2, 2, 0, 2, -4, 5, 5, -10})
	threeSum([]int{-1, 0, 1, 2, -1, -4})
}
