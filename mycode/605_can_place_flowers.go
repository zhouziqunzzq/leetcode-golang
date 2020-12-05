package mycode

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	fn := len(flowerbed)
	for i := range flowerbed {
		if flowerbed[i] == 0 &&
			(i == 0 || flowerbed[i-1] != 1) &&
			(i == fn-1 || flowerbed[i+1] != 1) {
			flowerbed[i] = 1
			n--
			if n == 0 {
				return true
			}
		}
	}

	return false
}
