package jan

func prefixCount(words []string, pref string) int {
	total := 0
	for _, word := range words {
		if stringPrefixes(word, pref) {
			total++
		}
	}
	return total
}

func stringPrefixes(s, prefix string) bool {
	n := len(prefix)
	if n > len(s) {
		return false
	}
	return s[:n] == prefix
}
