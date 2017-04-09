package main

import (
	"log"
	"os"
	"time"
)

func now() {
	t := time.Now()
	log.Printf("%s\n", t)
}

func main() {

	now()

	/* file based
	file, err := os.Open("/tmp/now")
	if (err != nil) {
	  log.Fatal(err)
	}
	s, err := file.Stat()
	if (err != nil) {
	  log.Fatal(err)
	}
	log.Printf("%v\n", s)
	f := s.FileInfo()
	log.Printf("%v\n", f)
	*/

	// FileInfo shortcut
	stat, err := os.Stat("/tmp/now")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v\n", stat.ModTime())
}
