package mycode

func isRobotBounded(instructions string) bool {
	curDir := 0
	var posAdj [4][2]int
	posAdj[0] = [2]int{0, 1}  // N
	posAdj[1] = [2]int{1, 0}  // E
	posAdj[2] = [2]int{0, -1} // S
	posAdj[3] = [2]int{-1, 0} // W
	x, y := 0, 0

	for _, i := range instructions {
		switch i {
		case 'G':
			x += posAdj[curDir][0]
			y += posAdj[curDir][1]
			break
		case 'L':
			curDir = (curDir + 4 - 1) % 4
			break
		case 'R':
			curDir = (curDir + 1) % 4
		}
	}

	return (x == 0 && y == 0) || curDir != 0
}
