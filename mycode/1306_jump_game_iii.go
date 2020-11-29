package mycode

// DFS
func canReach(arr []int, start int) bool {
	if start >= 0 && start < len(arr) && arr[start] >= 0 {
		if arr[start] == 0 {
			return true
		} else {
			// mark visited
			arr[start] = -arr[start]
			return canReach(arr, start-arr[start]) || canReach(arr, start+arr[start])
		}
	} else {
		return false
	}
}
