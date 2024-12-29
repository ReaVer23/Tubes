/*package main

import (
	"fmt"
	"math"
)

// Fungsi untuk memeriksa apakah sebuah bilangan prima (iteratif)
func isPrimeIterative(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Fungsi untuk mencari bilangan prima dalam rentang (iteratif)
func primesInRangeIterative(a, b int) []int {
	primes := []int{}
	for n := a; n <= b; n++ {
		if isPrimeIterative(n) {
			primes = append(primes, n)
		}
	}
	return primes
}

func main() {
	a, b := 10, 30
	fmt.Println("Bilangan prima (iteratif):", primesInRangeIterative(a, b))
}*/

/*package main

import (
	"fmt"
	"math"
)

func isPrimeRecursive(n, divisor int) bool {
	if n < 2 {
		return false
	}
	if divisor < 2 {
		return true
	}
	if n%divisor == 0 {
		return false
	}
	return isPrimeRecursive(n, divisor-1)
}

func primesInRangeRecursive(a, b int) []int {
	if a > b {
		return []int{}
	}
	if isPrimeRecursive(a, int(math.Sqrt(float64(a)))) {
		return append([]int{a}, primesInRangeRecursive(a+1, b)...)
	}
	return primesInRangeRecursive(a+1, b)
}

func main() {
	a, b := 10, 30
	fmt.Println("Bilangan prima (rekursif):", primesInRangeRecursive(a, b))
}*/

package main

import (
	"fmt"
	"math"
	"time"
)

func isPrimeIterative(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primesInRangeIterative(a, b int) []int {
	primes := []int{}
	for n := a; n <= b; n++ {
		if isPrimeIterative(n) {
			primes = append(primes, n)
		}
	}
	return primes
}

func isPrimeRecursive(n, divisor int) bool {
	if n < 2 {
		return false
	}
	if divisor < 2 {
		return true
	}
	if n%divisor == 0 {
		return false
	}
	return isPrimeRecursive(n, divisor-1)
}

func primesInRangeRecursive(a, b int) []int {
	if a > b {
		return []int{}
	}
	if isPrimeRecursive(a, int(math.Sqrt(float64(a)))) {
		return append([]int{a}, primesInRangeRecursive(a+1, b)...)
	}
	return primesInRangeRecursive(a+1, b)
}

func main() {
	a, b := 1, 500

	startIterative := time.Now()
	iterativePrimes := primesInRangeIterative(a, b)
	durationIterative := time.Since(startIterative)
	fmt.Println("Bilangan prima (iteratif):", iterativePrimes)
	fmt.Printf("Waktu eksekusi (iteratif): %v\n", durationIterative)

	startRecursive := time.Now()
	recursivePrimes := primesInRangeRecursive(a, b)
	durationRecursive := time.Since(startRecursive)
	fmt.Println("Bilangan prima (rekursif):", recursivePrimes)
	fmt.Printf("Waktu eksekusi (rekursif): %v\n", durationRecursive)
}
