package main

import "fmt"

func main() {
	fmt.Println("Hello, welcome to my poker app!!")

	d := makeDeck()
	d.shuffleDeck()
	d.printDeck()
}
