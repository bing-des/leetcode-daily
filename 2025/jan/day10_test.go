package jan

import "testing"

func TestWordSubsets(t *testing.T) {
	words1 := []string{"amazon", "apple", "facebook", "google", "leetcode"}
	words2 := []string{"e", "o"}

	subsets := wordSubsets(words1, words2)
	t.Logf("subsets: %v", subsets)
}
