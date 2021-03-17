package mycode

import "math"
import "math/rand"

// https://leetcode.com/problems/generate-random-point-in-a-circle/discuss/154037/Polar-Coordinates-10-lines
// https://meyavuz.wordpress.com/2018/11/15/generate-uniform-random-points-within-a-circle/
// Use polar coordinate! PDF and CDF proof is provided in the 2nd link.

type Solution478 struct {
	R float64
	X float64
	Y float64
}

func Constructor478(radius float64, x_center float64, y_center float64) Solution478 {
	return Solution478{radius, x_center, y_center}
}

func (this *Solution478) RandPoint() []float64 {
	r := math.Sqrt(rand.Float64()) * this.R
	theta := rand.Float64() * 2.0 * math.Pi
	x := r*math.Cos(theta) + this.X
	y := r*math.Sin(theta) + this.Y
	return []float64{x, y}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */
