package main

import (
  "github.com/pcarleton/go.dominion/card"
  "fmt"
)

type Player struct {
  Name string
  Deck []card.Card
  Discard []card.Card
  Hand []card.Card
}

func (p *Player) Gain(card card.Card) {
  p.Discard = append(p.Discard, card)
}

func (p *Player) Points() int {
  score := 0
  for _, card := range p.Deck {
    score += card.VictoryValue
  }
  for _, card := range p.Discard {
    score += card.VictoryValue
  }
  return score
}

func (p *Player) Cards() []card.Card {
  return append(p.Deck, p.Discard...)
}

func (p *Player) DeckCounts() (countMap map[string]int) {
  countMap = map[string]int{}
  for _, card := range p.Cards() {
    countMap[card.Name]++
  }
  return countMap
}

func (p *Player) String() string {
  return fmt.Sprintf("Player %s-> VP: %d Counts:%v\n Deck: %v Discard: %v", p.Name, p.Points(), p.DeckCounts(), p.Deck, p.Discard)
}


func main() {
  fmt.Println("Welcome to the Go Dominion Simulator")
  piles := map[card.Card]int{ card.Copper: 30, card.Estate: 5}
  deck := make([]card.Card, 10, 10)
  for i := 0; i < 7; i++ {
    deck = append(deck, card.Copper)
    //deck[i] = card.Copper
  }
  for i := 7; i < 10; i++ {
    deck = append(deck, card.Estate)
    //deck[i] = card.Estate
  }
  fmt.Println(deck)
  deck2 := make([]card.Card, len(deck))
  copy(deck2, deck)
  p1 := Player{"1", deck, nil, nil}
  p2 := Player{"2", deck2, nil, nil}
  for piles[card.Estate] > 0 {
    piles[card.Estate]--
    p1.Gain(card.Estate)
  }
  fmt.Println(p1.String())
  fmt.Println(p2.String())

}
