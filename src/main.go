package main

import (
	"fmt"
	"math"
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

	// Artmethic Operators
	x := 10
	y := 50

	// Sum
	result := x + y
	fmt.Println("Suma:", result)

	// Sub
	result = x - y
	fmt.Println("Resta:", result)

	// Mult
	result = x * y
	fmt.Println("Mult:", result)

	// Incremental
	x++

	// Calculating Cicle area
	radio := 10.0
	areaCirculo := math.Pi * radio * radio
	fmt.Println("Area Circulo:", areaCirculo)
}
