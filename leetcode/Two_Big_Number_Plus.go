package leetcode

func add(n1 string, n2 string) string {
	a,b, plus := len(n1)-1, len(n2)-1, uint8(0)
	var s []uint8
	for a >= 0 || b >= 0 || plus != 0 {
		res := lambda(n1, a) + lambda(n2, b) + plus
		plus = res/10
		s = append(s, res)

		a--
		b--
	}
	
	return string(s)
}

func lambda(s string, i int) uint8 {
	if i >= 0 {
		return s[i] - '0'
	} else {
		return 0
	}
}

//function add(a, b) {
//let pos = 0, res = ''
//while(a.length > pos || b.length > pos) {
//let carry = res.length - pos++
//res = (~~a[a.length - pos] + ~~b[b.length - pos] + carry) + res.substring(carry)
//}
//return res
//}