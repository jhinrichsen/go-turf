package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please specify folder(s) to delete")
	}
	for _, folder := range os.Args[1:] {
		fmt.Printf("Cleaning %v\n", folder)
		fis, err := ioutil.ReadDir(folder)
		if err != nil {
			log.Fatalf("cannot list directory %v\n", fis)
		}
		for _, fi := range fis {
			if isLiveMovie(fis, fi) {
				delete(filepath.Join(folder, fi.Name()), fi)
			}
		}
	}
}

func delete(filename string, fi os.FileInfo) {
	fmt.Printf("deleting %v (size %.2f MB)\n", filename,
		float64(fi.Size())/1024/1024)
	os.Remove(filename)
}

func hasCorrespondingJpg(fis []os.FileInfo, fi os.FileInfo) bool {
	// Name w/ extension
	filename := fi.Name()
	basename := strings.TrimSuffix(filename, path.Ext(filename))
	for _, f := range fis {
		if f.Name() == basename+".JPG" {
			return true
		}
	}
	return false
}

func isLiveMovie(fis []os.FileInfo, fi os.FileInfo) bool {
	return isMovie(fi) && hasCorrespondingJpg(fis, fi)
}

func isMovie(fi os.FileInfo) bool {
	// Live movies are lowercase while real movies are uppercase ext
	return strings.HasSuffix(fi.Name(), ".MOV") ||
		strings.HasSuffix(fi.Name(), ".mov")
}
