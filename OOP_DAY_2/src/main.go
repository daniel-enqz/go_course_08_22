package main

import (
	"fmt"
	pk "go_course_08_22/OOP_DAY_2/src/mypackage"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "BMW"
	myCar.Year = 2018
	fmt.Println(myCar)
}
