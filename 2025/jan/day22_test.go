package jan

import (
	"testing"
)

func highestPeak(isWater [][]int) [][]int {
	m, n := len(isWater), len(isWater[0])
	// first in first out
	var q []int
	visited := make([]bool, m*n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if isWater[i][j] == 1 {
				// is a water cell
				p := i*n + j
				q = append(q, p)
				visited[p] = true
			}
		}
	}

	directions := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	var h int
	for len(q) > 0 {
		start := len(q)
		for _, v := range q {
			row, col := v/n, v%n
			isWater[row][col] = h
			for k := 0; k < 4; k++ {
				newRow, newCol := row+directions[k][0], col+directions[k][1]
				newV := newRow*n + newCol
				if newRow < 0 || newRow >= m || newCol < 0 || newCol >= n || visited[newV] {
					continue
				}
				visited[newV] = true
				q = append(q, newV)
			}
		}
		h++
		q = q[start:]
	}
	return isWater
}

func TestHighestPeak(t *testing.T) {
	isWater := [][]int{{0, 1}, {0, 0}}
	highestPeak := highestPeak(isWater)
	t.Logf("%v \n", highestPeak)
}
