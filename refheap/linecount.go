// Based on K&R 'Programming in C', page 17

package main

import (
	"fmt"
)

// count characters in input
func main() {
	var s string
	nl := 0
	for {
		_, err := fmt.Scanln(&s)
		if err != nil {
			break
		}
		nl++
	}
	fmt.Printf("%d\n", nl)
}
