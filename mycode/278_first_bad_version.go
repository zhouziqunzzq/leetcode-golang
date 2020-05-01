package mycode

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func isBadVersion(version int) bool {
	if version > 3 {
		return false
	}
	return true
}

func firstBadVersion(n int) int {
	lo, hi := 0, n

	for lo < hi {
		mid := (lo + hi) / 2
		if !isBadVersion(mid) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return lo
}
