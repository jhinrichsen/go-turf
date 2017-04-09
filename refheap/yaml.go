package main

import (
	"fmt"
	// I fiddled around with zombiezen's version but could not find the handle, therefore:
	"github.com/kylelemons/go-gypsy/yaml"
	"log"
)
const (
	filename = "yaml.yml"
)

func main() {
	config, err := yaml.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error parsing config file %s: %s\n", filename, err)
	}
	fmt.Printf("key1 = %s\n", config.Get("key1"))
	fmt.Printf("anotherKey = %s\n", config.Get("anotherKey"))
}
