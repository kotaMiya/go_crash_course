package main

import "fmt"

func main() {
	ids := []int{64, 29, 48, 1, 98, 19}

	// loop through ids with index
	for i, id := range ids {
		fmt.Printf("%d - IDs: %d\n", i, id)
	}

	// loop without index
	for _, id := range ids {
		fmt.Printf("IDs: %d\n", id)
	}

	// add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with map
	emails := map[string]string{
		"Kota":  "kota@gmail.com",
		"Jesse": "jesse@gmail.com",
		"Lyte":  "lyte@gmail.com",
	}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	// key only
	for k := range emails {
		fmt.Println("Name: " + k)
	}
}
