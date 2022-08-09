package main

import "fmt"

func printFunction() {
	// Variable Declaration
	helloMessage := "Hello,"
	var worldMessage string = "world"
	// Println
	fmt.Println(helloMessage, worldMessage)

	// Printf - Specify data type
	name := "Platzi"
	courses := 500
	// %v assigns the data type automatically but a good practice is to do it manually
	fmt.Printf("%v tiene mas de %d cursos\n", name, courses)

	// Sprintf - Store in a variable
	message := fmt.Sprintf("%s tiene más de %d cursos", name, courses)
	fmt.Println(message)

	// Print the type
	const nombre string = "UltiRequiem"
	fmt.Printf("La variable 'nombre' es de tipo : %T\n", nombre)
}

func welcomeMessage(animal string, position int) {
	message := fmt.Sprintf("%s is the top %d animal in the world", animal, position)
	fmt.Printf("%s, Data type: %T\n", message, message)
}

func main() {
	printFunction()
	welcomeMessage("cat", 5)
}
