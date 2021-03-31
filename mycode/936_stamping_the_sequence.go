package mycode

type windowInfo936 struct {
	done map[int]bool // positions that match the target
	todo map[int]bool // positions that yet to be made to match the target
}

// https://leetcode.com/problems/stamping-the-sequence/solution/
// greedily working backwards
func movesToStamp(stamp string, target string) []int {
	m, n := len(stamp), len(target)
	// a queue of positions where is made matched and is needed to check all possible windows
	// that overlap with it
	q := make([]int, 0)
	flags := make([]bool, n) // flags[i]: target[i] has been made to a match
	ans := make([]int, 0)    // a stack storing stamp positions in reversed order since we're working backwards
	wInfos := make([]windowInfo936, n-m+1)

	// initialize window infos and queue
	for i := range wInfos {
		done := make(map[int]bool)
		todo := make(map[int]bool)
		for j, c := range stamp {
			if target[i+j] == uint8(c) {
				done[i+j] = true
			} else {
				todo[i+j] = true
			}
		}

		wInfos[i] = windowInfo936{done, todo}

		if len(todo) == 0 {
			// exact match, use this window to initialize the queue.
			// conceptually, we're replacing all chars in this window with '?'
			// as described in the solution.
			ans = append(ans, i)
			for j := i; j < i+m; j++ {
				q = append(q, j)
				flags[j] = true
			}
		}
	}

	for len(q) > 0 {
		i := q[0] // the position that could potentially affect other windows that overlap with it
		q = q[1:]

		for j := maxInt(0, i-m+1); j <= minInt(n-m, i); j++ {
			if _, ok := wInfos[j].todo[i]; ok { // this window (starting at j) is affected
				delete(wInfos[j].todo, i)
				if len(wInfos[j].todo) == 0 {
					ans = append(ans, j)
					for pos := range wInfos[j].done {
						if !flags[pos] {
							q = append(q, pos)
							flags[pos] = true
						}
					}
				}
			}
		}
	}

	// check flags
	for _, f := range flags {
		if !f {
			return make([]int, 0)
		}
	}

	reverseSlice(ans)
	return ans
}
