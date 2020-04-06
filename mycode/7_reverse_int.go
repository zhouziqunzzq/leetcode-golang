package mycode

func reverse(x int) int {
	const IntMax = int32(^uint32(0) >> 1)
	const IntMin = ^IntMax
	var ans int32 = 0
	for x != 0 {
		i := int32(x % 10)
		x /= 10
		if (ans > IntMax/10) || (ans == IntMax/10 && i > 7) {
			return 0
		}
		if (ans < IntMin/10) || (ans == IntMin/10 && i < -8) {
			return 0
		}
		ans = ans*10 + i
	}
	return int(ans)
}
