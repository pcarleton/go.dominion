package dominion

import (
  "math/rand"
  "log"
  "sort"
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

// playAction looks up the action function associated with a cards
// name and executes it in the context of the provided Turn.
func (c Card) playAction(turn *Turn) {
  if (c.Type != ACTION) {
    log.Fatal("Card is not an action card")
  }
  ActionFunctions[c.Name](turn)
}


type Turn struct {
  P *Player
  G *Game
  Actions int
  Buys int
}

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

func (p Pile) Less(i, j int) bool {
  return p[i].Type < p[j].Type || p[i].CoinPrice < p[j].CoinPrice || p[i].Name < p[j].Name
}

func (p *Pile) Sort() {
  sort.Sort(p)
}

//Adds the specified card to the end of the pile.
func (p *Pile) Add(card Card) {
  s := *p
  *p = append(s, card)
}

func (p *Pile) AddAll(other Pile) {
  s := *p
  *p = append(s, other...)
}

func StartingDeck() Pile {
  var deck Pile
  for i := 0; i < 7; i++ {
    deck.Add(Copper)
  }
  for i := 0; i < 3; i++ {
    deck.Add(Estate) 
  }
  return deck
}

func smithyFunc(turn *Turn) {
  turn.P.Draw(3)  
}

var ActionFunctions = map[string](func(*Turn)){
  "Smithy": smithyFunc,
}
