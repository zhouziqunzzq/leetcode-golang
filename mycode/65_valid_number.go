package mycode

import "regexp"

// Solution 1: DFA
// https://leetcode.com/problems/valid-number/discuss/23728/A-simple-solution-in-Python-based-on-DFA
// Solution 2: regex
// Solution 3: if-else
// https://leetcode.com/problems/valid-number/discuss/23738/Clear-Java-solution-with-ifs
func isNumber(s string) bool {
	re := regexp.MustCompile(`^\s*[-+]?((\d+(\.\d*)?)|(\.\d+))([eE][-+]?\d+)?\s*$`)
	return re.MatchString(s)
}
