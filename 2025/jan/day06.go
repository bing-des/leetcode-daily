package jan

import "math"

func minOperations(boxes string) []int {
	nums := make([]int, len(boxes))
	for i, _ := range boxes {
		for j, v1 := range boxes {
			if i == j {
				continue
			}
			if v1 == '1' {
				diff := j - i
				nums[i] += int(math.Abs(float64(diff)))
			}
		}
	}
	return nums
}
