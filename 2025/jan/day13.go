package jan

func minimumLength(s string) int {
	l := len(s)
	if l < 3 {
		return l
	}
	hots := [26]int{} // 数组的性能优于map

	for _, e := range s {
		hots[e-'a']++
	}

	size := 0
	for _, v := range hots {
		if v > 0 {
			if v&1 == 1 {
				// 奇数个最终保留一个
				size += 1
			} else {
				// 偶数个最终保留两个
				size += 2
			}
		}
	}
	return size
}
