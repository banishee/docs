package leetcode

func longestPalindrome(s string) string {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := getPalindromeLen(s, i, i)
		len2 := getPalindromeLen(s, i, i+1)
		if len1 < len2 {
			len1 = len2
		}

		if len1 > end-start+1 {
			start = i - (len1-1)/2
			end = i + len1/2
		}
	}

	return s[start : end+1]
}

func getPalindromeLen(s string, left, right int) int {
	for 0 <= left && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}
