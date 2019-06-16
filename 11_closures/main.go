package main

import "fmt"

// return func, so that sum in adder() is not initialise when it is called
func adder() func(int) int {
	sum := 0

	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// instance function
	sum := adder()

	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}
}
