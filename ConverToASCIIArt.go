package asciiArt

func ConvertToASCIIArt(input string) []int {
	var word []int
	for i := 0; i < len(input); i++ {
		char := int(rune(input[i]))
		if input[i] == '\\' && i < len(input) {
			if input[i+1] == 'n' {
				char = 127
				word = append(word, char)
				i = i + 1
			} else {
				word = append(word, char)
			}
		} else {
			word = append(word, char)
		}
	}
	return word
}
