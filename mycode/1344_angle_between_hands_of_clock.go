package mycode

func angleClock(hour int, minutes int) float64 {
	hAngle := float64(hour%12)*30.0 + float64(minutes)*0.5
	mAngle := float64(minutes) * 6.0
	ans := mAngle - hAngle

	if ans < 0 {
		ans = -ans
	}
	if ans > 180 {
		ans = 360 - ans
	}

	return ans
}
