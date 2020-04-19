package mycode

func binSearchHelper(nums []int, target, lo, hi int) int {
	for lo <= hi {
		mid := (lo + hi) / 2
		if nums[mid] == target {
			return mid
		} else {
			if target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		}
	}

	return -1
}

func search(nums []int, target int) int {
	if len(nums) <= 2 {
		for i, v := range nums {
			if v == target {
				return i
			}
		}
		return -1
	}

	if nums[0] < nums[len(nums)-1] {
		// not rotated, just do bin-search
		return binSearchHelper(nums, target, 0, len(nums)-1)
	} else {
		// find pivot
		lo, hi := 0, len(nums)-1
		pivot := 0
		r := nums[len(nums)-1]

		for lo <= hi {
			mid := (lo + hi) / 2
			// just do some lucky guessing :p
			if nums[mid] == target {
				return mid
			}
			if mid-1 >= 0 && nums[mid-1] > nums[mid] {
				pivot = mid
				break
			} else {
				if nums[mid] < r {
					hi = mid - 1
				} else {
					lo = mid + 1
				}
			}
		}

		// do bin-search in [0, pivot - 1] or [pivot, len(nums)-1]
		if target > r {
			// do bin-search in [0, pivot - 1]
			return binSearchHelper(nums, target, 0, pivot-1)
		} else {
			// do bin-search in [pivot, len(nums)-1]
			return binSearchHelper(nums, target, pivot, len(nums)-1)
		}
	}
}
