package main

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	gcdLength := gcd(len(str1), len(str2))
	return str1[:gcdLength]
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// brute force
// func gcdOfStrings(str1 string, str2 string) string {
// 	targetStr := str1
// 	if len(str1) < len(str2) {
// 		targetStr = str2
// 	}

// 	for i := len(targetStr); i > 0; i-- {
// 		divisor := targetStr[:i]
// 		if isDivided(str1, divisor) && isDivided(str2, divisor) {
// 			return divisor
// 		}
// 	}

// 	return ""
// }

// func isDivided(s string, t string) bool {
// 	lenS, lenT := len(s),len(t)
// 	if lenS % lenT != 0 {
// 		return false
// 	}

// 	for i := 0; i < lenS; i += lenT {
// 		if s[i:i+lenT] != t {
// 			return false
// 		}
// 	}
// 	return true
// }
