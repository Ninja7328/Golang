package main

import "fmt"

func main() {
	a := 15
	if a%3 == 0 {
		fmt.Println("Fuzz")
	} else if a%5 == 0 {
		fmt.Println("Buzz")
	} else if a%15 == 0 {
		fmt.Println("fuzzBuzz")

	}
}
