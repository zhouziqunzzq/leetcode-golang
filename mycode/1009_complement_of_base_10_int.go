package mycode

func bitwiseComplement(N int) int {
	if N == 0 {
		return 1
	}

	mask := 1 << 31
	for mask&N == 0 {
		mask >>= 1
	}
	mask -= 1

	return (^N) & mask
}
