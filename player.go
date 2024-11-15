package main

type Player struct {
    Name  string
    Hand  []Card
}

func NewPlayer(name string) *Player {
    return &Player{Name: name, Hand: []Card{}}
}

func (p *Player) DrawCard() Card {
    card := p.Hand[0]
    p.Hand = p.Hand[1:]
    return card
}

func (p *Player) AddCards(cards []Card) {
    p.Hand = append(p.Hand, cards...)
}
