package jan

import (
	"testing"
)

func gridGame(grid [][]int) int64 {
	firstSum := 0
	secondSum := 0
	for _, col := range grid[0] {
		firstSum += col
	}

	minSum := 2<<31 - 1
	for i, col := range grid[0] {
		firstSum -= col
		minSum = min(minSum, max(firstSum, secondSum))
		secondSum += grid[1][i]
	}
	return int64(minSum)
}

func TestGridGame(t *testing.T) {
	grid := [][]int{{2, 5, 4}, {1, 5, 1}}
	result := gridGame(grid)
	t.Logf("game over. result: %d\n", result)
}
