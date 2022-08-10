package main

import (
	"fmt"
	"log"
	"strconv"
)

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

	// Covertir texto a número
	value, err := strconv.Atoi("56")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value:", value)

	// EJERCICIOS
	num1 := 2

	if num1%2 == 0 {
		fmt.Println("Is even")
	}

	if checkUserRegistration("user1", "123456") {
		fmt.Println("Valid User")
	}
}

func checkUserRegistration(user, password string) bool {
	return user == "user1" && password == "123456"
}
