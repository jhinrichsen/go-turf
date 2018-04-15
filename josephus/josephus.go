package main

import "math"

// Golang only offers math.Pow(float64) float64
func pow(n, m int) int {
	e := 1
	for i := m; i > 0; i-- {
		e *= n
	}
	return e
}

// O(log(n))
func josephus1(n int) int {
	if n == 1 {
		return 1
	}

	if (n % 2) == 0 {
		return 2*josephus1(n/2) - 1
	}

	// if (n % 2) == 1 {
	return 2*josephus1((n-1)/2) + 1
}

// http://oeis.org/A006257
// O(1)
func josephus2(n int) int {
	return 1 + (2 * (n - pow(2, int(math.Floor(math.Log2(float64(n)))))))
}

func Josephus(n int) int {
	return josephus2(n)
}
