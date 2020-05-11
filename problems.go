package main

import (
	"fmt"
	"math"
)

func problem1(n int) int {
	var sum int = 0
	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	return sum
}

func problem2(limit int) int{
	var sum int = 2
	var dp []int = make([]int, 10000000)
	dp[0] = 1
	dp[1] = 2
	for i := 2; i <= 100 ; i++ {
		dp[i] = dp[i-1] + dp[i-2]
		if dp[i] > limit {
			break
		}
		if dp[i]%2 == 0 {
			sum += dp[i]
		}
	}

	return sum
}

func problem3(number int) int {
	sqrt := int(math.Sqrt(float64(number)))
	max := 1
	for i := 2; i <= sqrt; i++ {
		if number % i == 0  && isPrime(i) {
			max = int(math.Max(float64(max), float64(i)))
		}
	}

	if isPrime(number) {
		return number
	}

	return max
}

func isPrime(number int) bool {
	sqrt := int(math.Sqrt(float64(number)))

	for i := 2; i <= sqrt; i++ {
		if number % i == 0 {
			return false
		}
	}

	return true
}

func problem4() int {
	max := 0
	for i := 100; i <= 999; i++ {
		for j := 100; j <= 999; j++ {
			if isPalindrome(i * j) {
				max = int(math.Max(float64(max), float64(i * j)))
				fmt.Println(i, j)
			}
		}
	}

	return max
}

func isPalindrome(number int) bool {
	revNumber := 0
	temp := number

	for ; temp != 0;  {
		revNumber = revNumber *10 + temp % 10
		temp /= 10
	}

	return revNumber == number
}

func main() {
	fmt.Println(problem4())
}
