package mycode

func floodFillDFS(image [][]int, r, c, oldColor, newColor int) {
	image[r][c] = newColor

	// up
	if r-1 >= 0 && image[r-1][c] == oldColor {
		floodFillDFS(image, r-1, c, oldColor, newColor)
	}
	// down
	if r+1 < len(image) && image[r+1][c] == oldColor {
		floodFillDFS(image, r+1, c, oldColor, newColor)
	}
	// left
	if c-1 >= 0 && image[r][c-1] == oldColor {
		floodFillDFS(image, r, c-1, oldColor, newColor)
	}
	// right
	if c+1 < len(image[r]) && image[r][c+1] == oldColor {
		floodFillDFS(image, r, c+1, oldColor, newColor)
	}
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] != newColor {
		floodFillDFS(image, sr, sc, image[sr][sc], newColor)
	}
	return image
}
