package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20

	// Recorrer map

	for i, v := range m {
		fmt.Println(i, v)
	}

	// Encontrar un valor
	// Ok is used to see if a value is in a map
	value, ok := m["Jose"]
	fmt.Println(value, ok)
}
