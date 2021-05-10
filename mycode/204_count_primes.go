package mycode

// https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
func countPrimes(n int) int {
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}

	for i := 2; i*i < n; i++ {
		if !isPrime[i] {
			continue
		} else {
			// i is prime, mark all multiples of i as non-prime
			for j := i * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}

	ans := 0
	for i := 2; i < n; i++ {
		if isPrime[i] {
			ans++
		}
	}
	return ans
}
