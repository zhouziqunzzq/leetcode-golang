package mycode

func reverseBits(num uint32) uint32 {
	var ans uint32 = 0
	for i := 0; i < 32; i++ {
		ans |= num & 1
		num >>= 1
		if i < 31 {
			ans <<= 1
		}
	}
	return ans
}
