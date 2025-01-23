package jan

import "testing"

func countServers(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	rows := make([]int, m)
	cols := make([]int, n)
	// 如果数据量在10000以内，for rang 性能更好一些，如果数据量超过10000，使用for i性能更好一些，需要结合具体使用场景来看。
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				rows[i]++
				cols[j]++
			}
		}
	}

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && (rows[i] > 1 || cols[j] > 1) {
				count++
			}
		}
	}

	return count
}

func TestCountServer(t *testing.T) {
	grid := [][]int{{1, 1, 1, 0, 0, 1}, {1, 1, 0, 0, 1, 1}, {0, 0, 0, 0, 0, 0}, {1, 1, 0, 0, 0, 1}}
	count := countServers(grid)
	t.Logf("count: %d\n", count)
}
