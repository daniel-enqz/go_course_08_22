package main

import "fmt"

func message(text string, c chan string) {
	c <- text
}

func main() {
	c := make(chan string, 2)
	c <- "Messsage"
	c <- "Messsage2"

	fmt.Println(len(c), cap(c))

	// Range and Close
	close(c)

	// We cant add new lines because we closed the channel and the cap is 2
	// c <- "Message3"

	for message := range c {
		fmt.Println(message)
	}

	// Select
	email := make(chan string)
	email2 := make(chan string)
	go message("message1", email)
	go message("message2", email2)
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email:
			fmt.Println("Email recibido de email 1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email 2", m2)
		}
	}
}
