package jan

import "testing"

func TestXorAllNums(t *testing.T) {
	nums1 := []int{8, 6, 29, 2, 26, 16, 15, 29}
	nums2 := []int{24, 12, 12}
	t.Logf("%d", xorAllNums(nums1, nums2))
}
