func longestCommonPrefix(arr1 []int, arr2 []int) int {
	prefix_map := make(map[int]struct{}, 0)	
	for _, i := range arr1 {
		for ; i > 0; i = i / 10 {
			prefix_map[i] = struct{}{}
		}
	}

	v := 0
	for _, i := range arr2 {
		for ; i > 0; i = i / 10 {
			if _, ok := prefix_map[i]; ok {
				v = max(v, i)
			}
		} 
	}
    
	if v == 0 {
		return 0
	}
	return len(strconv.Itoa(v))
}