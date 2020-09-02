package mycode

import "math"

// https://leetcode.com/problems/contains-duplicate-iii/discuss/61645/AC-O(N)-solution-in-Java-using-buckets-with-explanation
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if k <= 0 || t < 0 {
		return false
	}
	w := t + 1
	buckets := make(map[int]int)
	for i, n := range nums {
		n = n - math.MaxInt32
		bucketId := n / w
		if _, ok := buckets[bucketId]; ok {
			return true
		}
		if candidate, ok := buckets[bucketId-1]; ok && n-candidate <= t {
			return true
		}
		if candidate, ok := buckets[bucketId+1]; ok && candidate-n <= t {
			return true
		}
		if len(buckets) >= k {
			lastBucketId := (nums[i-k] - math.MaxInt32) / w
			delete(buckets, lastBucketId)
		}
		buckets[bucketId] = n
	}
	return false
}
