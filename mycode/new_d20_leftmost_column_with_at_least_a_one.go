package mycode

/**
 * // This is the BinaryMatrix's API interface.
 * // You should not implement it, or speculate about its implementation
 * type BinaryMatrix struct {
 *     Get(int, int) int
 *     Dimensions() []int
 * }
 */

type BinaryMatrix interface {
	Get(int, int) int
	Dimensions() []int
}

// Solution 1: Scanning from the top-right corner.
func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {
	dims := binaryMatrix.Dimensions()
	n, m := dims[0], dims[1]
	i, j := 0, m-1
	ans := -1

	for i < n && j >= 0 {
		if binaryMatrix.Get(i, j) == 1 {
			ans = j
			j--
		} else {
			i++
		}
	}

	return ans
}

// Solution 2: Scanning from the bottom-right corner.
func leftMostColumnWithOneSolution2(binaryMatrix BinaryMatrix) int {
	dims := binaryMatrix.Dimensions()
	n, m := dims[0], dims[1]
	i, j := n-1, m-1
	ans := -1

	for i >= 0 && j >= 0 {
		if binaryMatrix.Get(i, j) == 1 {
			ans = j
			j--
		} else {
			i--
		}
	}

	return ans
}
