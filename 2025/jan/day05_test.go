package jan

import "testing"

func shiftingLetters(s string, shifts [][]int) string {
	n := len(s)
	diff := make([]int, n+1)
	for _, shift := range shifts {
		start, end, direction := shift[0], shift[1], shift[2]
		if direction == 0 {
			diff[start]--
			diff[end+1]++
		} else {
			diff[start]++
			diff[end+1]--
		}
	}
	arr := make([]byte, n)
	currDiff := 0
	for i := range s {
		currDiff += diff[i]
		e := (int(s[i]-'a') + currDiff) % 26
		if e < 0 {
			e += 26
		}
		arr[i] = byte(e) + 'a'
	}

	return string(arr)
}

func TestShiftingLetters(t *testing.T) {
	s := "xuwdbdqik"
	shifts := [][]int{{4, 8, 0}, {4, 4, 0}, {2, 4, 0}, {2, 4, 0}, {6, 7, 1}, {2, 2, 1}, {0, 2, 1}, {8, 8, 0}, {1, 3, 1}}
	letters := shiftingLetters(s, shifts)
	t.Log(letters)
}
