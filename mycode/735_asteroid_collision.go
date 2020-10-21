package mycode

func asteroidCollision(asteroids []int) []int {
	ans := make([]int, 0)
	if len(asteroids) == 0 {
		return ans
	}

	for _, a := range asteroids {
		if len(ans) == 0 {
			ans = append(ans, a)
		} else if ans[len(ans)-1] > 0 && a < 0 {
			// perform collision
			cur := -a
			isCurAlive := true
			for len(ans) > 0 && ans[len(ans)-1] > 0 {
				top := ans[len(ans)-1]
				if cur == top {
					ans = ans[:len(ans)-1]
					isCurAlive = false
					break
				} else if cur > top {
					ans = ans[:len(ans)-1]
				} else {
					isCurAlive = false
					break
				}
			}
			if isCurAlive {
				ans = append(ans, a)
			}
		} else {
			ans = append(ans, a)
		}
	}

	return ans
}
