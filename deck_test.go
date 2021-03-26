package main

import (
	"reflect"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := NewDeck()

	tests := []struct {
		want  [4]int8
		input int8
	}{
		// First card
		{[4]int8{1, 1, 1, 1}, 0},
		// Eleventh card
		{[4]int8{1, 2, 1, 3}, 11},
		// Twelve card
		{[4]int8{1, 2, 2, 1}, 12},
		// Fifteenth card
		{[4]int8{1, 2, 3, 1}, 15},
		// Last card
		{[4]int8{3, 3, 3, 3}, 80},
	}

	for _, test := range tests {
		if !reflect.DeepEqual(test.want, deck.cards[test.input]) {
			t.Logf("Deck in %d is %v, but wants %v\n", test.input, deck.cards[test.input], test.want)
		}
	}
}
