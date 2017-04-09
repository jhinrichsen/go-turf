package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i<=100; i++ {
		mod3 := i % 3 == 0
		mod5 := i % 5 == 0
		s := ""
		if mod3 && mod5 {
			s = "FizzBuzz"
		} else if mod3 {
			s = "Fizz"
		} else if mod5 {
			s = "Buzz"
		} else {
			s = strconv.Itoa(i)
		}
		fmt.Println(s)
	}
}
