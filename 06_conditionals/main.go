package main

import "fmt"

func main() {
	x := 5
	y := 10

	// if else
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	}

	// if else if
	color := "green"

	if color == "red" {
		fmt.Println("Color is red")
	} else if color == "blue" {
		fmt.Println("Color is blue")
	} else {
		fmt.Println("Color is NOT red or blue")
	}

	// switch
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is NOT red or blue")
	}
}
