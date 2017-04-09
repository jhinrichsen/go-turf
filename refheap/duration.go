package main

import (
	"fmt"
	"time"
)

// Calculate the duration between two points in time
func main() {
	d1 := time.Date(2012, time.September, 22, 0, 0, 0, 0, time.UTC)
	d2 := time.Date(2012, time.November, 7, 0, 0, 0, 0, time.UTC)
	fmt.Println(d2.Sub(d1))
	fmt.Println(d2.Sub(d1).Hours() / 24, "days")
}
