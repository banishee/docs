package main

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]bool, 10)
	length := 0
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	r := -1
	for i := 0; i < len(s); i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for r+1 < len(s) && !m[s[r+1]] {
			// 不断地移动右指针
			m[s[r+1]] = true
			r++
		}

		curLen := r - i + 1
		if length < curLen {
			length = curLen
		}
	}

	return length
}
