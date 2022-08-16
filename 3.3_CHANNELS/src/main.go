package main

func say(text string, c chan string) {
	c <- text
}

func main() {
	c := make(chan int, 1)

	fmt.Println("Hello")

	go say("Bye", c)
}
