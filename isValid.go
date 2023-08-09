package asciiArtFs

// Check for valid of characters by runes from 32 to 126
func IsValid(s string) bool {
	for _, ch := range s {
		if ch < ' ' && ch != 10 || ch > '~' {
			return false
		}
	}
	return true
}