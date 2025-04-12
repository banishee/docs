package main

func partitionLabels(s string) []int {
	last := [26]int{}
	for i := 0; i < len(s); i++ {
		last[s[i]-'a'] = i
	}

	group := []int{}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		if end < last[s[i]-'a'] {
			end = last[s[i]-'a']
		}

		if i == end {
			group = append(group, end-start+1)
			start = end + 1
		}
	}

	return group
}
