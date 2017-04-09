package main

// TODO
// - looks as if tmp wav files do not get deleted all times
// - add genre tag optionally

// Convert audio files from flac into mp3 format keeping all meta information
// Depends on externals commands 'flac' and 'lame' for decoding/ encoding

import (
	//	"bytes"
	"code.google.com/p/goflac-meta"
	// "goflac-meta"
	// "flac"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path"
	"runtime"
	"strings"
)

// An interface that knows how to derive output filenames from input filenames
//type Namer interface {
//	Derive(os.File) os.File
//}

type Namer func(os.File) os.File
type Locater func(Tags) string

// ID3 tag information
type Tags map[string]string

func Basename(f os.File) string {
	b := path.Base(f.Name())
	return b[0 : len(path.Ext(b))-1]
}

func Homedir() string {
    usr, err := user.Current()
    if err != nil {
        log.Fatal( err )
    }
    return usr.HomeDir
}

// Return a wav file for a given flac file
func wavFile(flacFile os.File) os.File {
	wavFile, err := ioutil.TempFile("", Basename(flacFile)+".wav-")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Using temporary wav file %q\n", wavFile)
	return *wavFile
}

// Convert a flac file to wav format
func flac2wav(in os.File, n Namer) (os.File, error) {
	out := n(in)
	cmd := exec.Command("flac",
		"-f",            // overwrite any existing file
		"--silent",	// output is useless because we are multiplexing goroutines
		"-d", in.Name(), // decode file (input)
		"-o", out.Name()) // output file
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	return out, err
}

// Return an mp3 filename including path for a given set of id tags
func mp3File(t Tags) string {
	fqp := path.Join(Homedir(), "Music", "audio", hierarchy(t)) + ".mp3"
	log.Printf("Using mp3 file %s\n", fqp)
	return fqp
}

// Convert a wav file to mp3 format
func wav2mp3(in os.File, l Locater, t Tags) (os.File, error) {
	out := l(t)
	// Make sure out exists
	dir := path.Dir(out)
	log.Printf("Creating album directory %s\n", dir)
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Using mp3 output location %s\n", out)
	cmd := exec.Command("lame",
		"--noreplaygain",	// TODO
		"-V", "4",		// high variable bitrate
		"--nohist",		// do not display histogram
	        "--silent",		// output is useless because we are multiplexing goroutines
		"--tt", t["title"],
		"--ta", t["artist"],
		"--tl", t["album"],
		"--ty", t["date"],
		"--tn", t["tracknumber"] + "/" + t["totaltracks"],
		"--add-id3v2",
		"--id3v2-only",
		in.Name(),		// input file
		out)			// output file
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err = cmd.Run()

	mp3File, err := os.Open(out)
	if err != nil {
		log.Fatal(err)
	}
	return *mp3File, err
}

func init() {
	flag.Parse()
}

// Remove a file
func cleanup(f os.File) {
	log.Printf("Removing temporary file %q\n", f)
	err := os.Remove(f.Name())
	if err != nil {
		log.Printf("Error: Cannot remove temporary file %q\n", f)
	}
}

// Convert a flac audio file to mp3
// id3 tag support is not primetime as of Dec 2012, so lame is used
func convert(flacFile string, c chan string) {
	flac, err := os.Open(flacFile)
	if err != nil {
		log.Fatal(err)
	}

	// Step 1. Create a wav from flac
	wav, err := flac2wav(*flac, wavFile)
	if err != nil {
		log.Fatal(err)
	}
	// Everything vanishes once its time has come
	defer cleanup(wav)
	log.Printf("wav: %q\n", wav)

	// Step 2. Extract meta tags
	tags := metadata(*flac)

	// Step 2. Create mp3 from wav
	mp3, err := wav2mp3(wav, mp3File, tags)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created file %s\n", mp3.Name())
	c <- mp3.Name()
}

// Determine where the generated mp3 should be created (some audio archive probably)
// <A..Z>/<artist>/<year>--<album>/<2 digit track>-<title>
func hierarchy(t Tags) string {
	// Left pad tracknumber to two digits
	track := fmt.Sprintf("%02s", t["tracknumber"])
	artist := t["artist"]
	s := path.Join(artist[0:1],
		artist,
		t["date"] + "--" + t["album"],
		track + "-" + t["title"])
	log.Printf("Group path for %q is %s\n", t, s)
	return s
}

// Use as many CPUs as possible to convert flac files in commandline parameters
// and wait until all goroutines have signaled end of processing
func main() {
	// Full speed ahead
	runtime.GOMAXPROCS(runtime.NumCPU())

	sync := make(chan string)
	for _, filename := range flag.Args() {
		go convert(filename, sync)
	}
	for i, _ := range flag.Args() {
		s := <- sync
		fmt.Printf("Processed #%d: %s\n", i, s)
	}
}

// Metadata for given flac file
func metadata(flacFile os.File) map[string]string {
	meta := flac.Metadata{}
	err := meta.Read(&flacFile)
	if err != nil {
		log.Fatal(err)
	}

	m := make(map[string]string)
	// fmt.Printf("Metadata for %s: %s\n", file, meta)
	for _, tag := range meta.VorbisComment.Data.Comments {
		log.Printf("Flac tag %s\n", tag)
		// Split into key and value
		s := strings.Split(tag, "=")
		if len(s) == 2 {
			m[strings.ToLower(s[0])] = s[1]
		} else {
			log.Println("Skipping unparsable flac tag %s\n", tag)
		}
	}
	return m
}
