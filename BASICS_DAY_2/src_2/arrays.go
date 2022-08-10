package main

import "fmt"

// Arrays are inmutable, fixed. Slices doesn't

func main() {
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	// Slices
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Métodos en el slice
	fmt.Println(slice[0])   // 1
	fmt.Println(slice[:3])  // [1 2 3]
	fmt.Println(slice[3:])  // [4 5 6]
	fmt.Println(slice[1:4]) // [2 3 4]

	// Append
	slice = append(slice, 7, 8, 9)
	fmt.Println(slice)

	// Append list
	slice = append(slice, []int{10, 11, 12}...)
	fmt.Println(slice)

	// Specify capacity and length
	x := make([]float64, 5, 10)
	fmt.Println(x)
}
