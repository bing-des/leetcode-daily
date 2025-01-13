package top105

func merge(nums1 []int, m int, nums2 []int, n int) {
	k := m + n - 1
	j := n - 1
	i := m - 1
	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}

func removeElement(nums []int, val int) int {
	k := 0
	for index, value := range nums {
		if value == val {
			k++
		} else {
			if k > 0 {
				nums[index-k] = value
			}
		}

	}
	return k
}

func removeDuplicates(nums []int) int {
	last := 0
	k := 0
	twice := false
	for i, v := range nums {
		if i == 0 || last != v || (!twice && last == v) {
			twice = i != 0 && last == v
			nums[k] = v
			k++
			last = v
		}
	}
	return k
}
