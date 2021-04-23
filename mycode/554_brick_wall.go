package mycode

// https://leetcode.com/problems/brick-wall/discuss/888577/IntuitionC%2B%2BWith-PicturesHashMapDetailed-ExplanationCommentsSolutionCode
// using hashmap to store number of bricks that share the same 'edge'
func leastBricks(wall [][]int) int {
	edgeCnt := make(map[int]int) // edge -> num of bricks that share this edge
	maxCnt := 0

	for _, row := range wall {
		curEdge := 0
		for i := 0; i < len(row)-1; i++ {
			curEdge += row[i]
			edgeCnt[curEdge] += 1
			maxCnt = maxInt(maxCnt, edgeCnt[curEdge])
		}
	}

	return len(wall) - maxCnt
}
