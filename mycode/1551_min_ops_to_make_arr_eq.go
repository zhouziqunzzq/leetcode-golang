package mycode

/*
1 3 5
t = 3
2

1 3 5 7 9
t = 5
4+2

1 3 5 7 9 11 13
t = 7
6+4+2

n%2==1:
2...(n-1)
#terms: (n-1)/2
(n-1)(n+1) / 4
(n*n-1)/4

==========

1 3
t = 2
1

1 3 5 7
t = 4
3+1

1 3 5 7 9 11
t = 6
5+3+1

n%2==0:
1...(n-1)
#terms: n/2
n*n/4
*/
func minOperations(n int) int {
	if n%2 == 0 {
		return n * n / 4
	} else {
		return (n*n - 1) / 4
	}
}
