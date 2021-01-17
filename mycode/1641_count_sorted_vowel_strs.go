package mycode

// https://leetcode.com/problems/count-sorted-vowel-strings/discuss/939122/3-solutions-(with-pictures)
// Solution 1: math
// multiset coefficients
// https://en.wikipedia.org/wiki/Multiset#Counting_multisets
// MultiC(5, n) = C(5+n-1, n) = C(n+4, 4)
func countVowelStrings(n int) int {
	return (n + 4) * (n + 3) * (n + 2) * (n + 1) / 24
}

// Solution 2: recursive method
//
func countVowelStringsRecursive(n int) int {
	return countVowelStringsRecursiveHelper(n, 1)
}

// f(n, i): num of strings of length n constructed using only vowels[i...k]
// f(n, i) = f(n-1, i) + f(n-1, i+1) + ... + f(n-1, k)
// where k is the number of vowels
func countVowelStringsRecursiveHelper(n, i int) int {
	// base case: empty string
	if n == 0 {
		return 1
	}

	rst := 0
	// try different vowels: vowel[i], vowel[i+1], ..., vowel[k]
	// here k = 5
	for v := i; v <= 5; v++ {
		rst += countVowelStringsRecursiveHelper(n-1, i)
	}

	return rst
}
