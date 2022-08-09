package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	// With package reflect we can print the type of variable
	fmt.Println(reflect.TypeOf(pi))
	fmt.Println(reflect.TypeOf(pi2))

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// Declaración de variables
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero Value
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Calcular cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println(areaCuadrado)

	age := 23
	fmt.Println(reflect.TypeOf(age))
}
