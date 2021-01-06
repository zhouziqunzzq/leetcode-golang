package mycode

func findKthPositive(arr []int, k int) int {
	prev := 0

	for _, n := range arr {
		if prev+1 != n {
			// some of nums are missing
			diff := n - prev - 1
			if k-diff <= 0 {
				return prev + k
			} else {
				k -= diff
			}
		}
		prev = n
	}
	return prev + k
}
