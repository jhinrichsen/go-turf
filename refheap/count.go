// Based on K&R 'Programming in C', page 16

package main

import (
	"bufio"
	"fmt"
	"os"
)

// count characters in input
func main() {
	reader := bufio.NewReader(os.Stdin)
	nc := 0
	for {
		_, err := reader.ReadByte()
		if err != nil {
			break
		}
		nc++
	}
	fmt.Printf("%d\n", nc)
}
