package dominion

import (
  "math/rand"
)

type CardType int
const (
  TREASURE CardType = iota
  VICTORY
  ACTION
)

type Card struct {
  Name string
  VictoryValue int
  CoinValue int
  CoinPrice int
  Type CardType
}

type Pile []Card

func (p *Pile) Pop() Card {
  s := *p
  if len(s) == 0 {
    return Card{}
  }
  card := s[len(s)-1]
  *p = s[:len(s)-1]
  return card
}

func (p *Pile) Len() int {
  return len(*p)
}

func (p *Pile) Shuffle() {
  for i := p.Len() - 1; i > 0; i-- {
    if j := rand.Intn(i + 1); i != j {
      p.Swap(i, j)
    }
  }
}

func (p Pile) Swap(i, j int) {
  p[i], p[j] = p[j], p[i]
}

func (p *Pile) Add(card Card) {
  s := *p
  *p = append(s, card)
}

var Copper = Card{"Copper", 0, 1, 0, TREASURE}
var Estate = Card{"Estate", 1, 0, 2, VICTORY}

func startingDeck() []Card {
  var deck []Card
  for i := 0; i < 7; i++ {
    deck = append(deck, Copper)
  }
  for i := 0; i < 3; i++ {
    deck = append(deck, Estate)
  }
  return deck
}
