package main

import (
	"bytes"
	"crypto/md5"
//	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
	"strings"
)

func die(err error ) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
        flag.Usage = func() {
                fmt.Fprintf(os.Stderr, "Usage: %s <string to md5>\n", os.Args[0])
                flag.PrintDefaults()
        }
        flag.Parse()
        if flag.NArg() != 1 {
                flag.Usage()
                os.Exit(1)
        }
	m := md5.New()
	io.WriteString(m, flag.Arg(1))
	checksum := m.Sum(nil)

	// big.Int uses big endian, MD5 uses little endian
	dec := new(big.Int)
	dec.SetBytes(checksum)

	// hex := new([]byte)
	upperCase := new(bytes.Buffer)
	fmt.Fprintf(upperCase, "%x", checksum)
	upperCase = strings.ToUpper(upperCase)
	fmt.Printf("%x %s %v %d\n", checksum, upperCase, dec, checksum)
}

