package jan

import "testing"

func TestMinOperations(t *testing.T) {
	operations := minOperations("110")
	t.Logf("%x \n", operations)
}
