package jan

func minimizeXor(num1 int, num2 int) int {
	b1 := countSetBits(num1)
	b2 := countSetBits(num2)

	if b1 == b2 {
		return num1
	}
	if b1 > b2 {
		// 从末尾删除
		return remove(num1, b1-b2)
	}
	return add(num1, b2-b1)
}

func remove(num int, n int) int {
	for n > 0 {
		num &= num - 1
		n--
	}
	return num
}

func add(num int, n int) int {
	pos := 0
	for n > 0 {
		for (num>>pos)&1 == 1 {
			pos++
		}
		num |= 1 << pos
		n--
	}
	return num
}

func countSetBits(num int) int {
	c := 0
	for num != 0 {
		c += num & 1
		num >>= 1
	}
	return c
}
