package main

import (
	"math/rand"
	"time"
)

type Card struct {
    Rank string
    Suit string
}

type Deck struct {
    Cards []Card
}

func NewDeck() *Deck {
    suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
    ranks := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}
    
    deck := &Deck{}
    
    for _, suit := range suits {
        for _, rank := range ranks {
            deck.Cards = append(deck.Cards, Card{Rank: rank, Suit: suit})
        }
    }
    
    return deck
}

func (d *Deck) Shuffle() {
    rand.Seed(time.Now().UnixNano())
    for i := range d.Cards {
        j := rand.Intn(i + 1)
        d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
    }
}
