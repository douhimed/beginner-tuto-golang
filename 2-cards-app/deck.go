package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardNumbers := []string{"Ace", "Two", "Threee", "Four"}

	for _, suit := range cardSuits {
		for _, number := range cardNumbers {
			cards = append(cards, number + " of " + suit)
		}
	}
 
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:];
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666);
}

func readDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error Reading from file ==> ", err)
		os.Exit(1)
	}
	
	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	myRand := rand.New(source)
	for i := range d {
		rp := myRand.Intn(len(d) - 1)
		d[i], d[rp] = d[rp], d[i]
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}