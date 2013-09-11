package dominion

import (
  "fmt"
)

type Player struct {
  Name string
  Deck []Card
  Discard []Card
  Hand []Card
}


func (p *Player) Gain(card Card) {
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

func (p *Player) Cards() []Card {
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

