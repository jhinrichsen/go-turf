package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	for _, f := range os.Args[1:] {
		fmt.Printf("%-20s: %v\n", f, path.IsAbs(f))
	}
}
