package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

type Card struct {
	Suit   string
	Value  string
	Points int
}

func (c Card) String() string {
	suitSymbol := c.Suit
	var colorPrint func(format string, a ...interface{}) string

	switch c.Suit {
	case "♥", "♦":
		colorPrint = color.New(color.FgRed).SprintfFunc()
	case "♠", "♣":
		colorPrint = color.New(color.FgWhite).SprintfFunc()
	}

	return fmt.Sprintf(`
┌─────────┐
│ %2s     │
│         │
│    %s   │
│         │
│      %2s│
└─────────┘`, 
		colorPrint(c.Value), 
		colorPrint(suitSymbol), 
		colorPrint(c.Value))
}

func displayCards(card1, card2 Card) {
	lines1 := strings.Split(card1.String(), "\n")
	lines2 := strings.Split(card2.String(), "\n")
	
	fmt.Println("\nPlayer 1     Player 2")
	for i := 0; i < len(lines1); i++ {
		fmt.Printf("%s   %s\n", lines1[i], lines2[i])
	}
}

func displayScores(score1, score2 int) {
	fmt.Println("\nScores:")
	fmt.Printf("Player 1: %d  Player 2: %d\n", score1, score2)
	fmt.Println(strings.Repeat("-", 30))
}
