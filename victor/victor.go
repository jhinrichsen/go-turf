// Victor, the cleaner
// A successor to logrotate

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
        "path/filepath"
	"time"
)

const (
	// Format in golang Duration
	maxAge = "240h"
)

// return true if file info is a regular file
func isFile(f os.FileInfo) bool {
	return 0 == f.Mode() & os.ModeType
}

func main() {
	logdir, err := filepath.Abs("/var/log/sync-portage")
        if err != nil {
		panic("Cannot determine absolute path for relative dir")
	}
	list, err := ioutil.ReadDir(logdir)
	if err != nil {
		log.Fatalf("Error reading directory: %v\n", err)
	}
	for i, f := range list {
		log.Printf("Testing file %d: %v\n", i, f.Name())
		log.Printf("\tModified date: %v, regular file: %t\n", f.ModTime(), isFile(f))
		log.Printf("\tAge: %v\n", time.Since(f.ModTime()))
		du, err := time.ParseDuration(maxAge)
		if err != nil {
			log.Fatalf("Error parsing duration %s: %s\n", maxAge, err)
		}
		expired := f.ModTime().Add(du).Before(time.Now())
		log.Printf("\tExpired: %t\n", expired)
                // FileInfo.Name() returns base name only, so we need to prepend path
		path := filepath.Join(logdir, f.Name())
		log.Printf("\tPath: %v\n", path)

		if isFile(f) && expired {
			fmt.Printf("Removing %s\n", path)
			err := os.Remove(path)
			if err != nil {
				log.Printf("Error removing %s: %s\n", path, err)
			}
			log.Printf("Removed %s\n", path)
		}
	}
}

// EOF

