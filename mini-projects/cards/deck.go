package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func (d *deck) print() {
	for i, card := range *d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cardSuites := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	cards := deck{}

	for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func deal(d *deck, handSize int) (deck, deck) {
	return (*d)[:handSize], (*d)[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(bytes), ","))
}

func (d deck) shuffle() {
	// Generate a new random source based on the current time to ensure different shuffles each time
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// Swap each card in the deck with a randomly chosen card
	for idx := range d {
		newPos := r.Intn(len(d))
		d[idx], d[newPos] = d[newPos], d[idx]
	}
}
