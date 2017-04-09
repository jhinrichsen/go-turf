package main

import (
	"path/filepath"
	"flag"
	"fmt"
	"image/jpeg"
	"log"
	"os"
)

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Compresses an image by reducing quality until it is smaller than a given size.
// Useful for ebay and alike that have an upper limit on image size, that is easily
// blown by most modern cameras.
func compress(filename,  outfilename string) {
	// Open a file
	file, err := os.Open(filename)
	die(err)

	// Read image metadata
	meta, err := jpeg.DecodeConfig(file)
	die(err)
	fmt.Printf("Image size is %dx%d\n", meta.Width, meta.Height)

	// Read the image
	file, err = os.Open(filename)
	img, err := jpeg.Decode(file)
	die(err)

	// Create an output file
	out, err := os.Create(outfilename)
	die(err)

	// And save
	err = jpeg.Encode(out, img, &jpeg.Options{Quality:95})
	die(err)
}

func init() {
	flag.Parse()
}

func main() {
	for i := 0; i<flag.NArg(); i++ {
		filename := flag.Arg(i)
		fmt.Println("Processing ", filename)
		compress(filename, Outfilename(filename))
	}
}

// Convert an input filename into an output filename
func Outfilename(filename string) string {
	abs, err := filepath.Abs(filename)
	die(err)

	base := filepath.Base(abs)
	die(err)
	ext := filepath.Ext(base)

	outfilename := base[0:len(base)-len(ext)] + "-compressed" + ext
	log.Printf("%s -> %s\n", filename, outfilename)
	return outfilename
}

