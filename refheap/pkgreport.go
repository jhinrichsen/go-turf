// Report package use on build server

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	dir = "/mnt/sol1/usr/sys/inst.images"
	ignore = "git-repos"
)

func main() {
	filepath.Walk(dir, visit)
}

func visit(path string, info os.FileInfo, err error) error {
	if err == nil {
		// log.Printf("%s contained in %s: %t\n", ignore, path, strings.Contains(path, ignore))
		if strings.Contains(path, ignore) {
			log.Printf("Skipping dir %s\n", path)
			return filepath.SkipDir
		}
		if strings.Contains(path, "AIX") && strings.Contains(path, ".usr.") {
			// Dissect package name
			pkgname := filepath.Base(path)
			parts := strings.Split(pkgname, ".")
			service := parts[0]
			function := parts[1]
			fmt.Printf("Found AIX package for service %s, function %s\n", service, function)
		} else if strings.Contains(path, "Linux") && strings.Contains(path, ".rpm") && !strings.Contains(path, "-") {
			// Dissect package name
			pkgname := filepath.Base(path)
			parts := strings.Split(pkgname, ".")
			name := parts[0]
			fmt.Printf("Found Linux package for %s\n", name)

		} else if strings.Contains(path, "SunOS") && strings.Contains(path, ".pkg") {
			// Dissect package name
			pkgname := filepath.Base(path)
			parts := strings.Split(pkgname, ".")
			name := parts[0]
			fmt.Printf("Found SunOS package for %s\n", name)
		} else {
			fmt.Printf("Ignoring %s\n", path)
		}
	} else {
		log.Println("Error accessing directory %s, error: %s\n", path, err)
	}
	return nil
}
