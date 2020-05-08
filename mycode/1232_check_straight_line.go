package mycode

func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) == 2 {
		return true
	}

	var slopeBase float64 = float64(coordinates[1][1]-coordinates[0][1]) /
		float64(coordinates[1][0]-coordinates[0][0])
	for i := 2; i < len(coordinates); i++ {
		if (float64(coordinates[i][1]-coordinates[i-1][1]) /
			float64(coordinates[i][0]-coordinates[i-1][0])) != slopeBase {
			return false
		}
	}
	return true
}
