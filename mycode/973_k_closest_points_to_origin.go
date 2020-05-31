package mycode

// Solution 1: quick select with O(N) avg runtime and O(N^2) worst runtime. <=
// Solution 2: max-heap with O(N logK) runtime.

func kClosest(points [][]int, K int) [][]int {
	l, r := 0, len(points)-1

	for l <= r {
		mid := kClosestQuickSelect(points, l, r)
		if mid == K {
			break
		} else if mid < K {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return points[:K]
}

func kClosestQuickSelect(points [][]int, l, r int) int {
	pivot := make([]int, 2)
	pivot[0], pivot[1] = points[l][0], points[l][1]

	// l ... <= pivot ... pivot ... >= pivot ... r
	for l < r {
		for l < r && kClosestDist(pivot) <= kClosestDist(points[r]) {
			r--
		}
		points[l][0], points[l][1] = points[r][0], points[r][1]
		for l < r && kClosestDist(points[l]) <= kClosestDist(pivot) {
			l++
		}
		points[r][0], points[r][1] = points[l][0], points[l][1]
	}
	points[l][0], points[l][1] = pivot[0], pivot[1]
	return l
}

func kClosestDist(p []int) int {
	return p[0]*p[0] + p[1]*p[1]
}
