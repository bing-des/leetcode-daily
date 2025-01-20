package jan

import "testing"

func TestFirstCompleteIndex(t *testing.T) {
	mat := [][]int{{4, 3, 5}, {1, 2, 6}}
	arr := []int{1, 4, 5, 2, 6, 3}

	completeIndex := firstCompleteIndex(arr, mat)
	t.Logf("first complete index %d\n", completeIndex)
}
