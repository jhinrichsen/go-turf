// BBP Formula
// Weisstein, Eric W. "BBP Formula." From MathWorld--A Wolfram Web Resource. http://mathworld.wolfram.com/BBPFormula.html

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Starting to calculate ")
	π := 0.0
	for n := 0; n < 50; n++ {
		step := ((4.0 / (8.0*float64(n) + 1.0)) - (2.0 / (8.0*float64(n) + 4.0)) - (1.0 / (8.0*float64(n) + 5.0)) - (1.0 / (8.0*float64(n) + 6.0))) * math.Pow(1.0/16.0, float64(n))
		π += step
		fmt.Printf("Step %n, step value: %.20f, sum: %.20f, delta: %.20f, digits: %.20f\n", n, step, π, math.Abs(math.Pi-π), math.Log10(1.0/math.Abs(math.Pi-π)))
	}
}
