package main

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	// [[1,3],[2,6],[8,10],[15,18]]
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := make([][]int, 1)
	merged[0] = intervals[0]
	for i := 1; i < len(intervals); i++ {
		pre := len(merged) - 1
		// [1,     5]
		//     [3,          10]
		//        [5,   7]
		if merged[pre][0] <= intervals[i][0] && intervals[i][0] <= merged[pre][1] {
			merged[pre][1] = max(merged[pre][1], intervals[i][1])
		} else { // merged[pre][0] < intervals[i][0]
			merged = append(merged, intervals[i])
		}
	}
	return merged
}

// func main() {
// 	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
// }
