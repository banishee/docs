package main

type node struct {
	next          [26]*node
	shortestIndex int
}

func stringIndices(wordsContainer []string, wordsQuery []string) []int {
	root := &node{}

	// insert string s into trie
	for index, s := range wordsContainer {
		cur := root
		for i := len(s) - 1; i >= 0; i-- {
			if len(s) < len(wordsContainer[cur.shortestIndex]) {
				cur.shortestIndex = index
			}

			c := s[i] - 'a'
			if cur.next[c] == nil {
				cur.next[c] = &node{shortestIndex: index}
			}

			cur = cur.next[c]
			if len(s) < len(wordsContainer[cur.shortestIndex]) {
				cur.shortestIndex = index
			}
		}
	}

	ans := make([]int, len(wordsQuery))
	// check answer
	for index, s := range wordsQuery {
		cur := root
		for i := len(s) - 1; i >= 0; i-- {
			c := s[i] - 'a'
			if cur.next[c] == nil {
				ans[index] = cur.shortestIndex
				break
			}
			ans[index] = cur.next[c].shortestIndex
			cur = cur.next[c]
		}
	}

	return ans
}

// func main() {
// 	stringIndices(
// 		[]string{"a", "b", "b"},
// 		[]string{"a", "b", "b"},
// 	)
// }
