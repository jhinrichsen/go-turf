package main

import (
	"fmt"
	"math"
)

func main() {
	a := 24.0
	c := 18.0
	α := math.Atan(a / c)
	fmt.Printf("%f\n", a/c)
	fmt.Printf("α = %f\n", α)
	fmt.Printf("α = %f\n", math.Tan(24.0/18.0))
	fmt.Printf("α = %f\n", math.Atan(1.3333333))
}
