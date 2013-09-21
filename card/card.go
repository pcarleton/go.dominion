package card

import (
  //dominion "github.com/pcarleton/go.dominion"
)

type CardType int

const (
  TREASURE CardType = iota
  VICTORY
  ACTION
)

type Stack interface {
  Deck() CardSlice
  Hand() CardSlice
  Discard() CardSlice
  DoDiscard(int)
  Draw(int) bool
  DiscardOne(string) bool
  Gain(*Card)
}


type Card struct {
  Name string
  VictoryValue int
  CoinValue int
  CoinPrice int
  Type CardType
}

type CardSlice []*Card

func (cs CardSlice) String() string {
  var pretty string
  for _, c := range cs {
    pretty += c.Name + ","
  }
  return pretty
}

type Pile struct {
  cards [100]*Card
  deck int
  hand int
  discard int
  size int
}

var Copper = Card{Name:"Copper"}
var Estate = Card{Name:"Estate"}

func NewPile() Pile {
  p := Pile{cards:[100]*Card{&Copper, &Copper, &Copper, &Copper, &Copper,
                             &Copper, &Copper, &Estate, &Estate, &Estate},
            deck:10,
            discard:10,
            size:10}
  return p  
}

func (p *Pile) Deck() CardSlice {
  return p.cards[0:p.deck]
}

func (p *Pile) Hand() CardSlice {
  return p.cards[p.deck:p.discard]
}

func (p *Pile) Discard() CardSlice {
  return p.cards[p.discard:p.size]
}

func (p *Pile) Gain(c *Card) {
  p.cards[p.size] = c
  p.size++
}

func (p *Pile) Draw(count int) bool {
  if count > p.size - len(p.Hand()) {
    return false
  }
  // We need to shuffle.
  for count > 0 && p.deck > 0 {
    count--
    p.deck--
  }
  // We've reached the end of the deck.
  // If we have still need cards, recycle discard
  if (count != 0) {
    p.deck = p.size - p.discard
    copy(p.cards[p.size:], p.Hand())
    copy(p.cards[0:p.deck], p.Discard())
    copy(p.cards[p.deck:p.size], p.cards[p.size:p.size + p.discard])
    p.discard = p.size
  }
  for count > 0 && p.deck > 0 {
    count--
    p.deck--
  }
  return true
}

func (p *Pile) DoDiscard(count int) {
  p.discard -= count
}

func (p *Pile) swap(i, j int) {
  p.cards[i], p.cards[j] = p.cards[j], p.cards[i]
}


//DiscardOne attempts to discard a card with the specified
//name.  Returns true on success.
func (p *Pile) DiscardOne(name string) bool {
  for i := p.discard - 1; i >= p.deck; i-- {
    if p.cards[i].Name == name {
      p.swap(i, p.discard)
      p.discard--
      return true
    }
  }
  return false
}


type LinkedStack {
  deck *LinkedNode
  hand *LinkedNode
  discard *LinkedNode
}

type LinkedNode struct {
  next *LinkedNode
  prev *LinkedNode
  c *Card
}




