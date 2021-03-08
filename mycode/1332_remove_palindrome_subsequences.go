package mycode

// this one pretend to be a hard problem and fxxked my brain...
// use the fact that it consists of only two kind of chars!
func removePalindromeSub(s string) int {
	if len(s) == 0 {
		return 0
	} else if isPalindrome(s) {
		return 1
	} else {
		return 2 // remove all 'a's first and then all 'b's...
	}
}
