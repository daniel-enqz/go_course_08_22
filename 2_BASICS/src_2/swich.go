package main

import "fmt"

func main() {
	switch modulo := 5 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	value := 0
	switch {
	case value > 100 && value < 200:
		fmt.Println("Greater than 100")
	case value < 100 && value > 0:
		fmt.Println("Less than 100")
	case value == 100:
		fmt.Println("Is 100")
	default:
		fmt.Printf("Out of range, you put %d\n", value)
	}
}
