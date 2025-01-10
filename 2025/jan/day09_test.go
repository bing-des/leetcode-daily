package jan

import "testing"

func TestPrefixCount(t *testing.T) {
	words := []string{"pay", "attention", "practice", "attend"}
	count := prefixCount(words, "at")
	t.Log(count)
}
