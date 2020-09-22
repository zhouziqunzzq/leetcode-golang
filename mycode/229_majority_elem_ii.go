package mycode

// https://leetcode.com/problems/majority-element-ii/solution/
func majorityElement(nums []int) []int {
	f1, f2 := false, false
	num1, num2 := 0, 0
	cnt1, cnt2 := 0, 0

	// 1st pass
	for _, n := range nums {
		if cnt1 == 0 && (!f2 || n != num2) {
			f1 = true
			num1 = n
		}
		if cnt2 == 0 && (!f1 || n != num1) {
			f2 = true
			num2 = n
		}

		if n == num1 {
			cnt1++
		} else if n == num2 {
			cnt2++
		} else {
			cnt1--
			cnt2--
		}
	}

	// 2nd pass
	cnt1, cnt2 = 0, 0
	ans := make([]int, 0)
	for _, n := range nums {
		if f1 && n == num1 {
			cnt1++
		}
		if f2 && n == num2 {
			cnt2++
		}
	}
	if f1 && cnt1 > len(nums)/3 {
		ans = append(ans, num1)
	}
	if f2 && cnt2 > len(nums)/3 {
		ans = append(ans, num2)
	}

	return ans
}
