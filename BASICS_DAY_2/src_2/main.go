package main

import "fmt"

func main() {
	// Defer - Se puede utlizar por ejemplo para cerrar archivos o base de datos.
	// Buena practica es utilizar un defer por función
	defer fmt.Println("defer")
	fmt.Println("Hello")

	// Continue and Break
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue // a pesar de que puede sucede un error continuar el siguiente ciclo
		}
		if i == 7 {
			break // Pasa algo determinado y detener el codigo si o sí
		}
		fmt.Println(i)
	}
}
