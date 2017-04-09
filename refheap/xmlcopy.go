package main

import (
	"fmt"
	"github.com/moovweb/gokogiri"
	"log"
)

const (
	buffer = `<?xml version="1.0" encoding="utf-8">
	<root>
	    <element1 attr1="value1"/>
	</root>`
)

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	doc, err := gokogiri.ParseXml([]byte(buffer))
	die(err)
	defer doc.Free()
	log.Printf("Parsed xml buffer %s\n", doc)
	log.Printf("Input encoding: %s\n", doc.InputEncoding())
	fmt.Printf(doc.String())
}
