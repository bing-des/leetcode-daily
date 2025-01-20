package jan

import (
	"testing"

	"container/heap"
)

func TestTrapRainWater(t *testing.T) {
	heightMap := [][]int{{1, 4, 3, 1, 3, 2}, {3, 2, 1, 3, 2, 4}, {2, 3, 3, 2, 3, 1}}
	waterTrapped := trapRainWater(heightMap)
	t.Logf("rain water %d\n", waterTrapped)
}

type Cell struct {
	row, col, height int
}

type PriorityQueue []Cell

// Len 计算长度
func (q PriorityQueue) Len() int           { return len(q) }
func (q PriorityQueue) Less(i, j int) bool { return q[i].height < q[j].height }
func (q PriorityQueue) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q *PriorityQueue) Push(x any) {
	*q = append(*q, x.(Cell))
}
func (q *PriorityQueue) Pop() any {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[0 : n-1]
	return x
}

func trapRainWater(heightMap [][]int) int {
	rows, cols := len(heightMap), len(heightMap[0])

	if rows < 3 && cols < 3 {
		return 0
	}

	q := &PriorityQueue{}
	heap.Init(q)
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}
	for row := 0; row < rows; row++ {
		heap.Push(q, Cell{row, 0, heightMap[row][0]})
		heap.Push(q, Cell{row, cols - 1, heightMap[row][cols-1]})
		visited[row][0] = true
		visited[row][cols-1] = true
	}

	for col := 0; col < cols; col++ {
		heap.Push(q, Cell{0, col, heightMap[0][col]})
		heap.Push(q, Cell{rows - 1, col, heightMap[rows-1][col]})
		visited[0][col] = true
		visited[rows-1][col] = true
	}

	directions := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	waterTrapped := 0
	height := 0
	for q.Len() > 0 {
		cell := heap.Pop(q).(Cell)
		height = intMax(height, cell.height)

		for _, direction := range directions {
			newRow, newCol := cell.row+direction[0], cell.col+direction[1]
			if newRow < 0 || newCol < 0 || newRow >= rows || newCol >= cols || visited[newRow][newCol] {
				continue
			}
			visited[newRow][newCol] = true
			neighbour := heightMap[newRow][newCol]
			if neighbour < height {
				waterTrapped += height - neighbour
			}
			heap.Push(q, Cell{newRow, newCol, heightMap[newRow][newCol]})
		}
	}
	return waterTrapped
}

func intMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
