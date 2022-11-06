package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver)
func (person Person) greet() string {
	return "Hello, my name is " + person.firstName + " " + person.lastName + " and I am " + strconv.Itoa(person.age)
}

// hasBirthday method (pointer receiver)
func (person *Person) hasBirthday() {
	person.age++
}

// getMarried (pointer receiver)
func (person *Person) getMarried(spouseLastName string) {
	if person.gender != "f" {
		return
	}
	person.lastName = spouseLastName
}

func main() {
	// in it person usin gstruct
	person1 := Person{firstName: "Lee", lastName: "Treehouse", gender: "f", age: 25}
	person2 := Person{"Holly", "Treehouse", "Melbourne", "f", 9}

	person1.hasBirthday()
	person1.getMarried("Smith")
	fmt.Println(person1)
	fmt.Println(person2)

	fmt.Println(person1.greet())
}
