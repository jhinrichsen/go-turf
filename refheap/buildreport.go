package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	dir = "/mnt/sol1/var/stage/hudson/jobs"
)

var baselevel = Level(dir)

var totalBuilds = 0

// Extract a hudson job name from a given directory
func Jobname(path string) string {
	l := strings.Split(path, "/")
	return l[len(l)-2]
}

// Determine level determined from root level for a given path/ directory
func Level(path string) int {
	return len(strings.Split(path, "/"))
}

func main() {
	filepath.Walk(dir, visit)
	fmt.Printf("Total of all build numbers: %d\n", totalBuilds)
}

func visit(path string, info os.FileInfo, err error) error {
        if err == nil {
		// Ugly: 2 is necessary because directories will end w/o slash
		if Level(path) - baselevel > 2 {
			log.Printf("Skipping %s because it is too deep\n", path)
			return filepath.SkipDir
		}
		if "nextBuildNumber" == filepath.Base(path) {
			log.Printf("Processing %s\n", path)
			buffer, err := ioutil.ReadFile(path)
			if err != nil {
				log.Fatalf("Error reading %s: %s\n", path, err)
			}
			// Strip trailing newline
			n, err := strconv.Atoi(string(buffer[0:len(buffer)-1]))
			if err != nil {
				log.Fatalf("Cannot parse build number in file %s: %s\n", path, err)
			}
			fmt.Printf("%40s %5d\n", Jobname(path), n)
			totalBuilds += n
		} else {
			log.Printf("Ignoring %s\n", path)
		}
	}
	return nil
}

