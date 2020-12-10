package mycode

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	i := 1
	up, down := false, false
	// go up
	for i < len(arr) && arr[i] > arr[i-1] {
		i++
		up = true
	}
	// go down
	for i < len(arr) && arr[i] < arr[i-1] {
		i++
		down = true
	}

	return i == len(arr) && up && down
}
