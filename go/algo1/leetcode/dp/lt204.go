package dp

import "fmt"

func countPrimes(n int) int {
	return eraPrimes(n)
}

func eraPrimes(n int) int {
	if n < 2 {
		return 0
	}
	isPrime := make([]bool, n)
	for i := 0; i < n; i++ {
		isPrime[i] = true
	}
	ans := 0
	for i := 2; i < n; i++ {
		if isPrime[i] {
			ans++
			for j := 2 * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	return ans
}

func forceCountPrimes(n int) int {
	if n < 2 {
		return 0
	}

	ans := 0

	for i := 2; i < n; i++ {
		isPrime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%d is prime\n", i)
			ans++
		}
	}
	return ans

}
