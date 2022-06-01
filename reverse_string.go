package reverse_string

func ReverseString(input string) (output string) {
	runes1 := []rune(input)
	for i, j := 0, len(runes1)-1; i < j; i, j = i+1, j-1 {
		runes1[i], runes1[j] = runes1[j], runes1[i]
	}
	return string(runes1)
}
