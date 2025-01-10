package jan

import (
	"fmt"
	"testing"
)

func TestStringMatch(t *testing.T) {
	words := []string{"jak", "yjakdn", "tj", "yyjakdn"}
	matching := stringMatching(words)
	fmt.Printf("%v \n", matching)
}
