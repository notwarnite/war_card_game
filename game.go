package main

import (
	"fmt"
)

type Game struct {
    Players []*Player
    Deck    *Deck
}

func NewGame(numPlayers int) *Game {
    game := &Game{}
    game.Deck = NewDeck()
    game.Deck.Shuffle()

    for i := 1; i <= numPlayers; i++ {
        player := NewPlayer(fmt.Sprintf("Player %d", i))
        game.Players = append(game.Players, player)
    }

    game.DistributeCards()
    return game
}

func (g *Game) DistributeCards() {
    for i := 0; i < len(g.Deck.Cards); i++ {
        g.Players[i%len(g.Players)].AddCards([]Card{g.Deck.Cards[i]})
    }
}

func (g *Game) Start() {
    round := 1
    for {
        fmt.Printf("Round %d\n", round)
        g.PlayRound()
        round++
        if g.IsGameOver() {
            break
        }
		fmt.Println("-----------------------")
    }
    g.DisplayWinner()
}

func (g *Game) PlayRound() {
    cardsInPlay := []Card{}
    
    for _, player := range g.Players {
        if len(player.Hand) == 0 {
            continue
        }
        card := player.DrawCard()
        fmt.Printf("%s plays %s of %s\n", player.Name, card.Rank, card.Suit)
        cardsInPlay = append(cardsInPlay, card)
    }
    
    g.DetermineRoundWinner(cardsInPlay)
}

func (g *Game) DetermineRoundWinner(cardsInPlay []Card) {

    winnerIndex := 0
    for i, _ := range cardsInPlay {
        if i > 0 {
            break
        }
    }

    winner := g.Players[winnerIndex]
    fmt.Printf("%s wins this round!\n", winner.Name)
    winner.AddCards(cardsInPlay)
}

func (g *Game) IsGameOver() bool {
    for _, player := range g.Players {
        if len(player.Hand) == 52 {
            return true
        }
    }
    return false
}

func (g *Game) DisplayWinner() {
    for _, player := range g.Players {
        if len(player.Hand) == 52 {
            fmt.Printf("%s wins the game!\n", player.Name)
            return
        }
    }
    fmt.Println("The game is over! No winner.")
}
