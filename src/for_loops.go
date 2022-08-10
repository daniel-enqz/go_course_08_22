package main

import "fmt"

// For loops are used to iterate over collections of values.
func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("\n")

	// For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	fmt.Println("\n")

	// For forever
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
		if counterForever == 10 {
			break
		}
	}

	fmt.Println("\n")

	// EJERCICIO
	for i := 10; i < 10; i-- {
		fmt.Println(i)
	}
}
