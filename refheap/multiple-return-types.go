package main

// Check if the signature of a function includes the return type to make it unique

type atype int

func (a atype) fn() int {
	return 0
}

func (a atype) fn() string {
	return "0"
}

func main() {
}

