package mycode

import "math"

// https://leetcode.com/problems/mirror-reflection/solution/
// Solution 1: simulation
func mirrorReflection(p int, q int) int {
	fp, fq := float64(p), float64(q)
	eps := 1e-6
	x, y := 0.0, 0.0 // current ray point
	rx, ry := fp, fq // ray direction vector

	for !((mirrorReflectionHelper(x, fp) &&
		(mirrorReflectionHelper(y, 0) || mirrorReflectionHelper(y, fp))) ||
		(mirrorReflectionHelper(x, 0) && mirrorReflectionHelper(y, fp))) {
		// find smallest positive t satisfying one of the following eqs:
		// x + rx * t = 0 (hit W)
		// y + ry * t = 0 (hit S)
		// x + rx * t = p (hit E)
		// y + ty * t = p (hit N)
		t := 1e9
		if (-x / rx) > eps {
			t = math.Min(t, -x/rx)
		}
		if (-y / ry) > eps {
			t = math.Min(t, -y/ry)
		}
		if ((fp - x) / rx) > eps {
			t = math.Min(t, (fp-x)/rx)
		}
		if ((fp - y) / ry) > eps {
			t = math.Min(t, (fp-y)/ry)
		}

		// update ray point
		x += rx * t
		y += ry * t

		// update direction vector
		if mirrorReflectionHelper(x, 0) || mirrorReflectionHelper(x, fp) {
			// hit W or E
			rx *= -1
		}
		if mirrorReflectionHelper(y, 0) || mirrorReflectionHelper(y, fp) {
			// hit S or N
			ry *= -1
		}
	}

	if mirrorReflectionHelper(x, fp) && mirrorReflectionHelper(y, fp) {
		return 1
	} else if mirrorReflectionHelper(x, fp) { // y -> 0
		return 0
	} else { // x -> 0, y -> p
		return 2
	}
}

func mirrorReflectionHelper(x, y float64) bool {
	return math.Abs(x-y) < 1e-6
}

// Solution 2: math
func mirrorReflection2(p int, q int) int {
	g := gcd(p, q)
	np := p / g // number of horizontal reflections
	nq := q / g // number of vertical reflections

	if np%2 == 1 && nq%2 == 0 {
		return 0
	} else if np%2 == 1 { // np % 2 == 1
		return 1
	} else {
		return 2
	}
}
