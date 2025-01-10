package jan

func stringMatching(words []string) []string {
	l := len(words)
	result := []string{}
	for i := 0; i < l; i++ {
		a := words[i]
		if arrayContains(result, a) {
			continue
		}
		for j := i + 1; j < l; j++ {
			b := words[j]
			if arrayContains(result, a) {
				break
			}
			if stringContains(b, a) {
				result = append(result, a)
			}
			if arrayContains(result, b) {
				continue
			}
			if stringContains(a, b) {
				result = append(result, b)
			}
		}
	}
	return result
}

func arrayContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func stringContains(s, substr string) bool {
	l := len(s)
	n := len(substr)
	if n > l {
		return false
	}
	t := l - n + 1
	for i := 0; i < t; i++ {
		if s[i:i+n] == substr {
			return true
		}
	}
	return false
}

func main() {

}
