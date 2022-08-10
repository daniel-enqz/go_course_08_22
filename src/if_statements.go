package main

import "fmt"

func main() {
	valor1 := 1
	valor2 := 2
	const nombre string = "UltiRequiem"
	is_type := fmt.Sprintf("%T", nombre)

	if valor1 == 1 && valor2 == 2 {
		fmt.Println("valor1 is 1 and valor2 is 2")
	} else {
		fmt.Println("valor1 is not 1 and valor2 is not 2")
	}

	if is_type == "string" {
		fmt.Println("is_type is a", is_type)
	} else {
		fmt.Println("is_type is not a string")
	}
}
