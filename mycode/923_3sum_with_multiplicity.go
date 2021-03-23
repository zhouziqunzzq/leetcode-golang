package mycode

// https://leetcode.com/problems/3sum-with-multiplicity/solution/
// Solution 1: 3 ptrs
// Solution 2: counting and combination
// Solution 3: reduce to 3-sum
// We implement Solution 3 here.
func threeSumMulti(arr []int, target int) int {
	const MaxVal = 100
	const Mod = 1_000_000_007

	cnt := make([]int, MaxVal+1)
	uniqueCnt := 0
	for _, v := range arr {
		cnt[v]++
		if cnt[v] == 1 {
			uniqueCnt++
		}
	}

	keys := make([]int, uniqueCnt)
	pk := 0
	for i, c := range cnt {
		if c > 0 {
			keys[pk] = i
			pk++
		}
	}

	var ans int64 = 0
	// reduce to 3-sum on "keys" for i <= j <= k
	// note that the "<=" being used here because of multiplicity of original array
	for i, x := range keys {
		t := target - x
		j, k := i, len(keys)-1
		// reduce to 2-sum with target=t
		for j <= k {
			y, z := keys[j], keys[k]
			if y+z < t {
				j++
			} else if y+z > t {
				k--
			} else { // y + z = t, so x + y + z = target
				if i < j && j < k {
					ans += int64(cnt[x] * cnt[y] * cnt[z])
				} else if i == j && j < k {
					ans += int64(cnt[x] * (cnt[x] - 1) / 2 * cnt[z])
				} else if i < j && j == k {
					ans += int64(cnt[x] * cnt[y] * (cnt[y] - 1) / 2)
				} else { // i == j == k
					ans += int64(cnt[x] * (cnt[x] - 1) * (cnt[x] - 2) / 6)
				}
				ans %= Mod
				j++
				k--
			}
		}
	}

	return int(ans)
}
