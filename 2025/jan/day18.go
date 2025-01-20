package jan

import "container/heap"

type Point struct {
	row, col, cost int
}

// MinHeap 有序队列
type MinHeap []Point

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Point))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minCost(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	dist := make([][]int, rows)
	for i := range dist {
		dist[i] = make([]int, cols)
		for j := range dist[i] {
			dist[i][j] = 1e9
		}
	}

	dist[0][0] = 0

	directions := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	h := &MinHeap{}
	heap.Push(h, Point{0, 0, 0})

	for h.Len() > 0 {
		p := h.Pop().(Point)
		row, col, cost := p.row, p.col, p.cost
		if cost > dist[row][col] {
			continue
		}
		for i := 0; i < 4; i++ {
			newRow, newCol := row+directions[i][0], col+directions[i][1]
			if newRow < 0 || newCol < 0 || newRow >= rows || newCol >= cols {
				continue
			}
			newCost := cost
			if grid[row][col] != i+1 {
				newCost++
			}
			if newCost < dist[newRow][newCol] {
				dist[newRow][newCol] = newCost
				heap.Push(h, Point{newRow, newCol, newCost})
			}
		}
	}
	return dist[rows-1][cols-1]
}
