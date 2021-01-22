package main

import (
	"fmt"
	"math/rand"
)

//Card is an object containing one Suit and one Rank.
type Card struct {
	Suit string
	Rank int
}

//Deck of cards containing a slice of cards
type Deck struct {
	Cards []Card
}

func makeDeck() Deck {
	suits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	ranks := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

	var d Deck
	for _, suit := range suits {
		for _, rank := range ranks {
			d.Cards = append(d.Cards, Card{Suit: suit, Rank: rank})
		}
	}
	return d
}

func (d Deck) shuffleDeck() {
	rand.Seed(1) //TODO pick different seed once actually wanting random shuffles
	for i := range d.Cards {
		r := rand.Intn(len(d.Cards))
		d.Cards[i], d.Cards[r] = d.Cards[r], d.Cards[i]
	}
}

func (c Card) printCard() {
	fmt.Printf("%v of %v", c.Rank, c.Suit)
}

func (d Deck) printDeck() {
	for _, card := range d.Cards {
		card.printCard()
		fmt.Print("\n")
	}
}
