package jan

func wordSubsets(words1 []string, words2 []string) []string {
	hotMap := wordsToHotMap(words2)
	hotLen := len(hotMap)

	result := make([]string, 0)
	for _, word := range words1 {
		if len(word) < hotLen {
			continue
		}
		if stringSubsets(word, hotMap) {
			result = append(result, word)
		}
	}

	return result
}

func wordsToHotMap(words []string) map[rune]int {
	hotMap := make(map[rune]int)
	for _, word := range words {
		wordHot := stringToHotMap(word)
		for key, value := range wordHot {
			if hotMap[key] < value {
				hotMap[key] = value
			}
		}
	}
	return hotMap
}

func stringToHotMap(s string) map[rune]int {
	hotMap := make(map[rune]int)
	for _, l := range s {
		hotMap[l]++
	}
	return hotMap
}

func stringSubsets(word string, hotMap map[rune]int) bool {
	wordHot := stringToHotMap(word)
	for key, value := range hotMap {
		if wordHot[key] < value {
			return false
		}
	}
	return true
}
