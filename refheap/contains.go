package main

import (
	"fmt"
	"strings"
)

const (
	sub = "git-repos"
	s = "/mnt/sol1/usr/sys/inst.images/git-repos"
)

func main() {
	fmt.Printf("%s contained in %s: %t\n", sub, s, strings.Contains(s, sub))
}

