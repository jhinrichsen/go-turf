// Validate that a file is UTF-8 (only)

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"unicode/utf8"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please provide some filenames as commandline parameter")
		return
	}
	for _, filename := range os.Args[1:] {
		log.Printf("Validating file %s\n", filename)
		buffer, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatalf("Cannot read file %s, %s\n", filename, err)
		}
		valid := utf8.Valid(buffer)
		if !valid {
			fmt.Println("%s: Not valid UTF-8", filename)
		}
	}
}
