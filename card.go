package dominion

import (
  "math/rand"
  "log"
  "sort"
  "time"
  //"fmt"
)

type CardType int

const (
  TREASURE CardType = iota
  VICTORY
  ACTION
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

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
    if j := r.Intn(i + 1); i != j {
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

func (p *Pile) Swap(i, j int) {
  s := *p
  //fmt.Printf("i: %d, j: %d\n", i, j)
  //fmt.Printf("s[i]: %v, s[j] %v\n", s[i], s[j])
  s[i], s[j] = s[j], s[i]
  //fmt.Printf("s[i]: %v, s[j] %v\n", s[i], s[j])
  *p = s
  //fmt.Println(p)
}

func (p Pile) Less(i, j int) bool {
  return p[i].Type < p[j].Type || p[i].CoinPrice < p[j].CoinPrice || p[i].Name < p[j].Name
}

func (p *Pile) Sort() {
  sort.Sort(p)
}

func (p Pile) Count(cardName string) int {
  count := 0
  for _, card := range p {
    if card.Name == cardName {
      count++
    }
  }
  return count
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
