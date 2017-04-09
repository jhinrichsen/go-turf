package main

// Import section
import (
	"fmt"
	"math"
)

// Prints pi
func main() {
	fmt.Println(Pi(5000))
}

// pi launches n goroutines to compute an
// approximation of pi.
func Pi(n int) float64 {
	ch := make(chan float64)
	for k := 0; k <= n; k++ {
		go term(ch, float64(k))
	}
	f := 0.0
	for k := 0; k <= n; k++ {
		f += <-ch
	}
	return f
}

// Term computes an iteration for pi
func term(ch chan float64, k float64) {
	ch <- 4 * math.Pow(-1, k) / (2*k + 1)
}
