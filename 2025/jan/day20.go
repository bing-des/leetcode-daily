package jan

type point struct {
	x int
	y int
}

func firstCompleteIndex(arr []int, mat [][]int) int {
	m, n := len(mat), len(mat[0])

	nums := make([]int, m*n+1)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			nums[mat[i][j]] = i*m + j
		}
	}

	rowCount := make([]int, m)
	colCount := make([]int, n)

	l := len(arr)
	for i := 0; i < l; i++ {
		v := arr[i]
		p := nums[v]
		row, col := p/m, p%n
		rowCount[row]++
		colCount[col]++

		if rowCount[row] == n || colCount[col] == m {
			return i
		}
	}

	return -1
}
