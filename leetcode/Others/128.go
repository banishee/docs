package main

func longestConsecutive(nums []int) int {
	m := make(map[int]struct{})
	for _, n := range nums {
		m[n] = struct{}{}
	}

	maxListLength := 0
	for n, _ := range m {
		// create list
		length := 1
		delete(m, n)
		// 0 <- n
		for left := n - 1; left >= -1e9; left-- {
			if _, ok := m[left]; !ok {
				break
			}
			length++
			delete(m, left)
		}
		// n -> max
		for right := n + 1; right <= 1e9; right++ {
			if _, ok := m[right]; !ok {
				break
			}
			length++
			delete(m, right)
		}

		// check
		if maxListLength < length {
			maxListLength = length
		}
	}

	return maxListLength
}

func main() {
	longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
}
