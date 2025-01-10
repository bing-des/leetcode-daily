package jan

func countPrefixSuffixPairs(words []string) int {
	wordsLen := len(words)
	pairs := 0
	for i := 0; i < wordsLen; i++ {
		for j := i + 1; j < wordsLen; j++ {
			if isPrefixAndSuffix(words[i], words[j]) {
				pairs++
			}
		}
	}
	return pairs
}

func isPrefixAndSuffix(str1, str2 string) bool {
	str1Len := len(str1)
	str2Len := len(str2)
	if str1Len > str2Len {
		return false
	}
	prefix := str2[0:str1Len]
	suffix := str2[str2Len-str1Len:]
	return str1 == prefix && str1 == suffix
}
