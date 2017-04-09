package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func die(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// Create a set of subfolder named for each month with same permissions as base folder
func main() {
	folder, err := os.Getwd()
	die(err)

	fi, err := os.Stat(folder)
	die(err)

	t := time.Date(2012, time.January, 1, 1, 1, 1, 1, time.UTC)
	for i := 1; i < 13; i++ {
		s := fmt.Sprintf("%02d-%s", i, t.Month().String())
		fmt.Println("Making directory ", s)
		os.Mkdir(s, fi.Mode())
		t = t.AddDate(0, 1, 0)
	}
}

