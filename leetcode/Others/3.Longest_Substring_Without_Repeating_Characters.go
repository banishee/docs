package main

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]struct{})
	length := 0

	for left, right := 0, -1; right < len(s)-1; {
		// next char in right must not exsit in m
		if _, ok := m[s[right+1]]; !ok {
			m[s[right+1]] = struct{}{}
			right++
		} else {
			delete(m, s[left])
			left++
		}

		length = max(length, right-left+1)
	}

	return length
}

func main() {
	lengthOfLongestSubstring("abcabcbb")
}
