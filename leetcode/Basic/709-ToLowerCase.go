package basic

func toLowerCase(s string) string {
	result := []rune(s)
	for i, char := range result {
		if char >= 'A' && char <= 'Z' {
			result[i] = char + ('a' - 'A')
		}
	}
	return string(result)
}
