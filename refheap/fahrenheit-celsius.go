// Based on K&R 'Programming in C', page 8

package main

import "fmt"

// print Fahrenheit-Celsius table for f = 0, 20, ..., 300
func main() {
	for fahr := 0.0; fahr <= 300; fahr += 20 {
		celsius := (5.0 / 9.0) * (fahr - 32.0)
		fmt.Printf("%4.0f %6.1f\n", fahr, celsius)
	}
}
