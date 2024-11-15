package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Deck []Card

func NewDeck() Deck {
	suits := []string{"♠", "♥", "♦", "♣"}
	values := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	points := map[string]int{
		"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10,
		"J": 11, "Q": 12, "K": 13, "A": 14,
	}

	var deck Deck
	for _, suit := range suits {
		for _, value := range values {
			deck = append(deck, Card{
				Suit:   suit,
				Value:  value,
				Points: points[value],
			})
		}
	}
	return deck
}

func (d Deck) Shuffle() Deck {
	rand.Seed(time.Now().UnixNano())
	for i := range d {
		j := rand.Intn(len(d))
		d[i], d[j] = d[j], d[i]
	}
	return d
}

func playWar() {
	fmt.Println("\n=== Welcome to War Card Game ===")
	fmt.Println("Press Enter to start each round...")

	// Initialize and shuffle deck
	deck := NewDeck().Shuffle()
	
	// Deal cards
	player1 := deck[:len(deck)/2]
	player2 := deck[len(deck)/2:]
	
	score1, score2 := 0, 0
	round := 1

	for len(player1) > 0 && len(player2) > 0 {
		fmt.Printf("\nRound %d", round)
		fmt.Scanln() // Wait for Enter key

		// Draw cards
		card1 := player1[0]
		card2 := player2[0]
		player1 = player1[1:]
		player2 = player2[1:]

		// Display cards
		displayCards(card1, card2)

		// Determine winner
		switch {
		case card1.Points > card2.Points:
			score1++
			fmt.Println("\nPlayer 1 wins the round!")
		case card2.Points > card1.Points:
			score2++
			fmt.Println("\nPlayer 2 wins the round!")
		default:
			fmt.Println("\nIt's a tie!")
		}

		displayScores(score1, score2)
		round++
	}

	// Game over
	fmt.Println("\n=== Game Over ===")
	if score1 > score2 {
		fmt.Println("Player 1 wins the game!")
	} else if score2 > score1 {
		fmt.Println("Player 2 wins the game!")
	} else {
		fmt.Println("It's a tie game!")
	}
}