package jan

func canConstruct(s string, k int) bool {
	if len(s) < k {
		return false
	}

	if len(s) == k {
		return true
	}

	bitmask := 0
	for _, char := range s {
		bitmask ^= 1 << (char - 'a')
	}

	oddCount := 0
	for bitmask > 0 {
		oddCount += bitmask & 1
		bitmask >>= 1
	}

	return oddCount <= k
}
