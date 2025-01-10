package jan

import (
	"fmt"
	"testing"
)

func TestCountPrefixSuffixPairs(t *testing.T) {
	words := []string{"a", "aba", "ababa", "aa"}
	pairs := countPrefixSuffixPairs(words)
	fmt.Println(pairs)
}
