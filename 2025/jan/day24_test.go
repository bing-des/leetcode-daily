package jan

import (
	"fmt"
	"testing"
)

func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	// stores the number of edges entering node x
	indegree := make([]int, n)
	// adjacency list
	adj := make([][]int, n)

	for i := 0; i < n; i++ {
		if len(graph[i]) > 0 {
			for _, node := range graph[i] {
				adj[node] = append(adj[node], i)
				indegree[i]++
			}
		}
	}

	var q []int
	for i := 0; i < n; i++ {
		if indegree[i] == 0 {
			q = append(q, i)
		}
	}

	var safed = make([]bool, n)
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		safed[node] = true
		for _, neigbor := range adj[node] {
			indegree[neigbor]--
			if indegree[neigbor] == 0 {
				q = append(q, neigbor)
			}
		}
	}

	var safeNodes []int
	for i := 0; i < n; i++ {
		if safed[i] {
			safeNodes = append(safeNodes, i)
		}
	}

	return safeNodes
}

func TestEventualSafeNodes(t *testing.T) {
	graph := [][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}
	safeNodes := eventualSafeNodes(graph)
	fmt.Printf("safe nodes: %v\n", safeNodes)
}
