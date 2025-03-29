package main

func groupAnagrams(strs []string) [][]string {
	group := make(map[[26]int][]string)
	for _, str := range strs {
		k := [26]int{}
		for _, c := range str {
			k[c-'a'] += 1
		}
		group[k] = append(group[k], str)
	}

	res := make([][]string, len(group))
	i := 0
	for _, arr := range group {
		res[i] = arr
		i++
	}

	return res
}
