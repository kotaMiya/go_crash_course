package main

import "fmt"

func main() {
	// using var
	var name = "kota"
	var age = 24

	test := "shorthand assignment"

	fmt.Println(name, age, test)
	fmt.Printf("%T\n", age)
}
