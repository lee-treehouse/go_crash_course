package strutil

func Reverse(s string) string {
	// making an array of all the chars in the string
	// but what is rune, it's some int type i think
	runes := []rune(s)

	// i is being set to 0
	// j is being set to len(runes)-1
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
