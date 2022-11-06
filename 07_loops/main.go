package main

import "fmt"

func main() {

	// Long method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// short method
	for i := 1; i <= 10; i++ {
		fmt.Printf("Number: %d\n", i)
	}

	// Fizz Buzz
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			println("fizz buzz")
		} else if i%3 == 0 {
			println("fizz")
		} else if i%5 == 0 {
			println("buzz")
		} else {
			println(i)
		}
	}
}
