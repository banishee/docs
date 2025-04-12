package main

func generate(numRows int) [][]int {
	ans := [][]int{{1}}
	for i := 1; i < numRows; i++ {
		newAns := make([]int, len(ans[i-1])+1)
		newAns[0] = 1
		newAns[len(newAns)-1] = 1
		for j := 1; j < len(newAns)-1; i++ {
			newAns[j] = ans[i-1][j-1] + ans[i-1][j]
		}
		ans = append(ans, newAns)
	}
	return ans
}

func main() {
	generate(5)
}
