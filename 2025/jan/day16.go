package jan

func xorAllNums(nums1 []int, nums2 []int) int {
	// xor 交换律
	l1, l2 := len(nums1), len(nums2)
	if l1%2 == 0 {
		if l2%2 == 0 {
			return 0
		}
		return xorArrays(nums1)
	} else {
		if l2%2 == 0 {
			return xorArrays(nums2)
		}
		return xorArrays(nums1) ^ xorArrays(nums2)
	}
}

func xorArrays(nums []int) int {
	res := 0
	for _, v := range nums {
		res ^= v
	}
	return res
}
