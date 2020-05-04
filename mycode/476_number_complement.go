package mycode

func findComplement(num int) int {
	if num == 0 {
		return 1
	}

	var num32 int32 = int32(num)
	var mask int32 = ^0

	for num32 & mask != 0 {
		mask <<= 1
	}

	return int(^num32 & ^mask)
}
