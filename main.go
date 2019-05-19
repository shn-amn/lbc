package main

import (
	"./fizzbuzz"

	"fmt"
)

func main() {
	f := fizzbuzz.New(3, 5, "fizz", "buzz")
	for i := 1; i < 50; i++ {
		fmt.Println(f(i))
	}
}
