package jan

import "testing"

func TestMinCost(t *testing.T) {
	grid := [][]int{{1, 1, 1, 1}, {2, 2, 2, 2}, {1, 1, 1, 1}, {2, 2, 2, 2}}
	cost := minCost(grid)
	t.Logf("min cost %d\n", cost)
}
