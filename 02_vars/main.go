package main

import "fmt"

func main() {
	var name string = "Lee"
	email, age := "leentaylor@gmail.com", 40

	fmt.Println(name, age, email)
	fmt.Printf("%T\n", name)
}
