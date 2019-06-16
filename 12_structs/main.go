package main

import (
	"fmt"
	"strconv"
)

// Define Person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// greeting method (value receiver)
// this function is for Person struct, like class and this
func (p Person) greeting() string {
	return "Hi, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
// this function is for Person struct,
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried method (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "male" {
		return
	}
	if p.gender == "female" {
		p.lastName = spouseLastName
	}
}

func main() {
	// init person using struct
	person1 := Person{
		firstName: "Kota",
		lastName:  "Miyamoto",
		city:      "Tokyo",
		gender:    "male",
		age:       24,
	}

	person2 := Person{
		firstName: "Anna",
		lastName:  "Kurata",
		city:      "Tokyo",
		gender:    "female",
		age:       24,
	}

	// do not need to say key
	// person1 := Person{
	// 	"Kota",
	// 	"Miyamoto",
	// 	"Tokyo",
	// 	"male",
	// 	24,
	// }

	fmt.Println(person1)
	fmt.Println(person1.firstName)
	// you can call value receiver method following struct
	person1.hasBirthday()

	// Anna is going to be Miyamoto
	person2.getMarried("Miyamoto")

	fmt.Println(person1.greeting())
	fmt.Println(person2.greeting())
}
