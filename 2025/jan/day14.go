package jan

func findThePrefixCommonArray(A []int, B []int) []int {
	n := len(A)
	res := make([]int, n)

	countMap := make(map[int]int)
	count := 0
	for i := 0; i < n; i++ {
		a := A[i]
		b := B[i]
		countMap[a]++
		if countMap[a] == 2 {
			count++
		}
		countMap[b]++
		if countMap[b] == 2 {
			count++
		}
		res[i] = count
	}
	return res
}
