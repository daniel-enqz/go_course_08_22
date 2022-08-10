package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar.year)

	// Other way
	var myCar2 car
	myCar2.brand = "Mercedes Benz"
	fmt.Println(myCar2, myCar2.year, myCar2.brand)

	// Zero Value
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)
}
