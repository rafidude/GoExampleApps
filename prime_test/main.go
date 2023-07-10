package main

import "fmt"

// isPrime checks if a number is prime, returns a bool and a string
func isPrime(n int) (bool, string) {
	if n == 1 {
		return false, "1 is not prime"
	}
	for i := 2; i < n/2+1; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not prime, %d is a divisor", n, i)
		}
	}
	return true, fmt.Sprintf("%d is prime", n)
}

func main() {
	fmt.Println(isPrime(15))
}
