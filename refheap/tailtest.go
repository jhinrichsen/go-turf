package main

import (
	"fmt"
	"github.com/ActiveState/tail"
	"log"
)

func main() {
	t, err := tail.TailFile("/tmp/test.log", tail.Config{Follow: true})
	if err != nil {
		log.Fatal(err)
	}
	for line := range t.Lines {
		fmt.Println(line.Text)
	}
}
