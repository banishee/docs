package main

func diStringMatch(S string) []int {
	maxNum, minNum := len(S), 0
	res := make([]int, len(S)+1)
	for i := 0; i < len(S); i++ {
		if S[i] == 'I' {
			res[i] = minNum
			minNum++
		} else {
			res[i] = maxNum
			maxNum--
		}
	}
	res[len(S)] = minNum

	return res
}
