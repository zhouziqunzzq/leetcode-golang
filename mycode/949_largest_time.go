package mycode

import "strconv"

func isValid(digits []int, hour int, minute int) bool {
	target := make([]int, 10)
	for i := 0; i < 10; i++ {
		target[i] = 0
	}
	if hour < 10 {
		target[0]++
		target[hour]++
	} else {
		target[hour/10]++
		target[hour%10]++
	}
	if minute < 10 {
		target[0]++
		target[minute]++
	} else {
		target[minute/10]++
		target[minute%10]++
	}
	for i := 0; i < 10; i++ {
		if target[i] != digits[i] {
			return false
		}
	}
	return true
}

func largestTimeFromDigits(A []int) string {
	digits := make([]int, 10)
	for i := 0; i < 10; i++ {
		digits[i] = 0
	}
	for _, v := range A {
		digits[v]++
	}
	for hour := 23; hour >= 0; hour-- {
		for minute := 59; minute >= 0; minute-- {
			if isValid(digits, hour, minute) {
				hs := strconv.Itoa(hour)
				if hour < 10 {
					hs = "0" + hs
				}
				ms := strconv.Itoa(minute)
				if minute < 10 {
					ms = "0" + ms
				}
				return hs + ":" + ms
			}
		}
	}
	return ""
}
