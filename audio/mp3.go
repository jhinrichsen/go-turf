package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path"
	"strings"
	"time"
)

const (
	// Relative to user's home directory
	destinationBaseDir = "Music/audio"

	mp3info = "/usr/local/bin/mp3info"
	tagAlbum = "%l"
	tagYear = "%y"
	tagArtist = "%a"
)

// Apply mp3 gain
func adjustLoudness(dir string, mp3s []string) {
	cmd := exec.Command("mp3gain", "-p", "-s", "i", "-a")
	// Add list of mp3 file basenames
	for _, f := range mp3s {
		cmd.Args = append(cmd.Args, path.Base(f))
	}
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error executing mp3gain: %v\n", err)
	}
	fmt.Println(string(out))
}

// Determine artist for a given record directory
func artist(files []string) string {
	// Functional map a list of files to a list of artists
	artists := make([]string, len(files))
	for i, file := range files {
		output, err := exec.Command("/usr/local/bin/mp3info", "-p", "%a", file).Output()
		if err != nil {
			log.Fatalf("Error executing mp3info on file %s, error %s\n", file, err)
		}
		artists[i] = strings.TrimSpace(string(output))
	}

	// Distinct
	m := make(map[string] bool)
	for _, artist := range artists {
		m[artist] = true
	}
	if len(m) > 1 {
		log.Fatalf("More than one artist: %s\n", m)
	}

	// Check length of artist name
	a := artists[0]
	if len(a) < 2 || len(a) > 80 {
		log.Fatalf("Artist name %s has suspicious length %d\n", a, len(a))
	}
	return a
}

// Determine a relative destination directory for a given set of audio files
func destination(mp3s []string) string {
	album := tag(mp3s, tagAlbum)
	artist := tag(mp3s, tagArtist)
	year := tag(mp3s, tagYear)
	return path.Join(artist[0:1], artist, year + "--" + album)
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Determine a distinct id3 tag from an album folder
// Aborts if tag is not the same for all audio files in folder
func tag(files []string, id3tag string) string {
	tags := make([]string, len(files))
	for i, file := range files {
		output, err := exec.Command("/usr/local/bin/mp3info", "-p", id3tag, file).Output()
		if err != nil {
			log.Fatalf("Error executing mp3info on file %s: %s\n", file, err)
		}
		tags[i] = strings.TrimSpace(string(output))
	}

	// Distinct
	m := make(map[string] bool)
	for _, tag := range tags {
		m[tag] = true
	}
	if len(m) > 1 {
		log.Fatalf("More than one tag: %s\n", m)
	}

	// Check length of tag name
	a := tags[0]
	if len(a) < 2 || len(a) > 80 {
		log.Fatalf("Tag name %s has suspicious length %d\n", a, len(a))
	}
	return a
}

func main() {
	flag.Parse()

	mp3s := flag.Args()
	if len(mp3s) == 0 {
		log.Fatalf("No *.mp3 files specified\n")
	}
	// Make sure .mp3 are mp3 format
	for _, mp3 := range mp3s {
		verifyFormatIsMp3(mp3)
	}

	// User home dir
	user, err := user.Current()
	die(err)

	destDir := path.Join(user.HomeDir, destinationBaseDir, destination(mp3s))
	log.Printf("Using destination directory %s\n", destDir)
	mkdir(destDir, path.Dir(mp3s[0]))

	// Move files, this will probably not work above partitions/ filesystems
	for _, from := range mp3s {
		to := path.Join(destDir, path.Base(from))
		log.Printf("Moving from %s to %s\n", from, to)
		// Rename() will probably only work in the same filesystem
		if err := os.Rename(from, to); err != nil {
			log.Fatalf("Cannot move from %s to %s\n", from, to)
		}
	}

	adjustLoudness(destDir, mp3s)
}

// Create destination directory
// Will fatal() on about anything
func mkdir(destDir string, permDir string) {
	// Does destination directory exist already?
	if fi, _ := os.Stat(destDir); fi != nil {
		log.Fatalf("Destination directory %s already exists, aborting.", destDir)
	}

	// Use mode/ permission of base directory
	fi, err := os.Stat(permDir)
	if err != nil {
		log.Fatalf("Cannot stat %s: %s", permDir, err)
	}
	log.Printf("Creating %s using %v\n", destDir, fi.Mode())
	err = os.MkdirAll(destDir, fi.Mode())
	if err != nil {
		log.Fatalf("Cannot create directory %s, %v\n", destDir, err)
	}
}

func run() {
	start := time.Now()
	out, err := exec.Command("date").Output()
	duration := time.Since(start)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("It took %v to fetch the date %s\n", duration, out)
}

// Verify that a given filename congtains mp3 file data
func verifyFormatIsMp3(filename string) {
	output, err := exec.Command("file", "--brief", filename).CombinedOutput()
	if err != nil {
		log.Fatalf("Error executing 'file' on %s, error %s\n", path.Base(filename), err)
	}
	out := strings.TrimSpace(string(output))
	// Typical output: MPEG ADTS, layer III, v1, 320 kbps, 48 kHz, JntStereo
	if !strings.Contains(out, "MPEG") || !strings.Contains(out, "layer III") {
		log.Fatalf("Expected file %s to be of type MPEG and layer III but it is %s\n", filename, out)
	}
	log.Printf("Verified that file %s is an mp3 file (%s)\n", path.Base(filename), out)
}

