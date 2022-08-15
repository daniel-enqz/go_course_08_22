package main

import "fmt"

type pc struct {
	ram int
	cpu int
}

func (myPc pc) ping() {
	fmt.Println(myPc.ram, "Pong")
}

func (myPc *pc) duplicateRAM() {
	myPc.ram = myPc.ram * 2
}

func main() {
	a := 50
	b := &a

	fmt.Println(a, b) // 50 0xc420014080
	fmt.Println(*b)   // 50

	*b = 100       // Will change that space in memory value, then a will be 100
	fmt.Println(a) // 100

	myPc := pc{ram: 8, cpu: 200} // Instance
	fmt.Println(myPc)            // {8, 200}

	myPc.ping() // {8, Pong}

	myPc.ram = 16     // Changing this by accessing the instance
	fmt.Println(myPc) // {16, 200}

	myPc.duplicateRAM() // Changing this by using pointer method
	fmt.Println(myPc)   // {32, 200}

	myPc.duplicateRAM()
	fmt.Println(myPc) // {64, 200}
}
