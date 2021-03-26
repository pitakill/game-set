package main

import "fmt"

// Deck of SET game
type Deck struct {
	// 81 cards
	cards [81][4]int8
}

func (d *Deck) debugCards(c ...int) {
	d.showCards(c...)

	if len(c) == 3 {
		fmt.Println(d.isSet(d.cards[c[0]], d.cards[c[1]], d.cards[c[2]]))
	}
}

func (d *Deck) showCards(cards ...int) {
	for _, i := range cards {
		fmt.Println(d.cards[i])
	}
}

func (d *Deck) isSet(a, b, c [4]int8) bool {
	var (
		control      = [4]int8{}
		i, t    int8 = 0, 0
	)

	for i < 4 {
		control[i] = (a[i] + b[i] + c[i]) % 3
		i++
	}

	for _, property := range control {
		t += property
	}

	return t == 0
}

func NewDeck() *Deck {
	deck := new(Deck)

	// A card is an array [4]int8
	var idCard int

	// Fill my deck with all posible cards
	for i := 1; i <= 3; i++ { // shape
		for j := 1; j <= 3; j++ { // color
			for k := 1; k <= 3; k++ { // number
				for l := 1; l <= 3; l++ { // shading
					deck.cards[idCard] = [4]int8{int8(i), int8(j), int8(k), int8(l)}
					idCard++
				}
			}
		}
	}

	return deck
}
