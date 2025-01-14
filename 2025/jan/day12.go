package jan

func canBeValid(s string, locked string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}

	open := 0
	for i := 0; i < n; i++ {
		if s[i] == '(' || locked[i] == '0' {
			open++
		} else {
			open--
		}
		if open < 0 {
			return false
		}
	}

	end := 0
	for i := n - 1; i > 0; i-- {
		if s[i] == ')' || locked[i] == '0' {
			end++
		} else {
			end--
		}
		if end < 0 {
			return false
		}
	}
	return true
}
