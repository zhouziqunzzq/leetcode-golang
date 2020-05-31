package mycode

// Solution 1: quick select with O(N) avg runtime and O(N^2) worst runtime. <=
// Solution 2: max-heap with O(N logK) runtime.

import (
	"math/rand"
	"time"
)

func kClosest(points [][]int, K int) [][]int {
	rand.Seed(time.Now().UnixNano())
	kClosestQuickSelect(points, K, 0, len(points)-1)
	return points[:K]
}

func kClosestQuickSelect(points [][]int, K, l, r int) {
	if l > r {
		return
	}

	// partition by pivot
	pivot := l + rand.Intn(r - l + 1)	// [l, r - l + l + 1)
	pivotDist := kClosestDist(points[pivot])
	points[l][0], points[l][1], points[pivot][0], points[pivot][1] =
		points[pivot][0], points[pivot][1], points[l][0], points[l][1]
	i := l
	for j := l+1; j <= r; j++ {
		if kClosestDist(points[j]) <= pivotDist {
			i++
			points[i][0], points[i][1], points[j][0], points[j][1] =
				points[j][0], points[j][1], points[i][0], points[i][1]
		}
	}
	points[l][0], points[l][1], points[i][0], points[i][1] =
		points[i][0], points[i][1], points[l][0], points[l][1]

	leftLen := i - l + 1
	if K < leftLen {
		kClosestQuickSelect(points, K, l, i-1)
	} else if K > leftLen {
		kClosestQuickSelect(points, K, i+1, r)
	}
	// else K = leftLen and we can just return. The points[:K] would be the answer.
}

func kClosestDist(p []int) int{
	return p[0] * p[0] + p[1] * p[1]
}
