package main

import "fmt"

func main() {
	// Define map (long version)
	// emails := make(map[string]string)

	// Assign key and values
	// emails["Kota"] = "kota@gmail.com"
	// emails["Jesse"] = "jesse@gmail.com"
	// emails["Lyte"] = "lyte@gmail.com"

	// Declare map and add key values (same as above)
	emails := map[string]string{
		"Kota":  "kota@gmail.com",
		"Jesse": "jesse@gmail.com",
		"Lyte":  "lyte@gmail.com",
	}

	fmt.Println(emails)
	fmt.Println(emails["Kota"])

	// add new key value
	emails["Jake"] = "jake@gmail.com"

	// delete from map
	delete(emails, "Jesse")
	fmt.Println(emails)
}
