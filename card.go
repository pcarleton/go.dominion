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

type Game struct {
  players []Player
}

type Action interface{
  Play(*Game) 
}

type ActionCard Card



type Pile []Card

//Removes and returns the last card in the pile.
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

// Shuffles the pile of cards randomly.
func (p *Pile) Shuffle() {
  for i := p.Len() - 1; i > 0; i-- {
    if j := rand.Intn(i + 1); i != j {
      p.Swap(i, j)
    }
  }
}

//Finds the specified card, removes it from the pile and returns it.
// Returns nil if the card is not found in the pile.
func (p *Pile) Remove(target Card) Card {
  for i, card := range *p {
    if card.Name == target.Name {
      p.Swap(i, p.Len() - 1)
      return p.Pop()
    }
  }
  return Card{}
}

func (p Pile) Swap(i, j int) {
  p[i], p[j] = p[j], p[i]
}

//Adds the specified card to the end of the pile.
func (p *Pile) Add(card Card) {
  s := *p
  *p = append(s, card)
}

var Copper = Card{"Copper", 0, 1, 0, TREASURE}
var Estate = Card{"Estate", 1, 0, 2, VICTORY}

func startingDeck() Pile {
  var deck Pile
  for i := 0; i < 7; i++ {
    deck.Add(Copper)
  }
  for i := 0; i < 3; i++ {
    deck.Add(Estate) 
  }
  return deck
}
