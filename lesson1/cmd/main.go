package main

import (
	"fizzbuzz/fizzbuzz"
	"fmt"
)

func main() {
	for i := 1; i < 101; i++ {
		fmt.Println(fizzbuzz.FizzBuzz(i))
	}
}
