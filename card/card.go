package card

import (
  //dominion "github.com/pcarleton/go.dominion"
  "math/rand"
  "sort"
  "time"
  "fmt"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))
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
  Use(string) bool
  Gain(*Card)
}

func DefaultStack() Stack {
  s := NewArrayStack()
  return &s
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

func (cs *CardSlice) Swap(i, j int) {
  (*cs)[i], (*cs)[j] = (*cs)[j], (*cs)[i]
}

func (cs CardSlice) Less(i, j int) bool {
  return cs[i].Type < cs[j].Type || cs[i].CoinPrice < cs[j].CoinPrice || cs[i].Name < cs[j].Name
}

func (cs *CardSlice) Len() int {
  return len(*cs)
}

func (cs *CardSlice) Sort() {
  sort.Sort(cs)
}

func (cs *CardSlice) GetActions() CardSlice {
  cs.Sort()
  actionIndex := 0
  for i, c := range *cs {
    if c.Type == ACTION {
      actionIndex = i
      break
    }
  }
  return (*cs)[actionIndex:]
}


type ArrayStack struct {
  cards [100]*Card
  deck int
  used int
  discard int
  size int
}

var Copper = Card{Name:"Copper", Type:TREASURE}
var Estate = Card{Name:"Estate", Type:VICTORY}
var Village = Card{Name:"Village", Type:ACTION}

func NewArrayStack() ArrayStack {
  p := ArrayStack{cards:[100]*Card{&Copper, &Copper, &Copper, &Copper, &Copper,
                             &Copper, &Copper, &Estate, &Estate, &Estate},
            deck:10,
            discard:10,
            size:10}
  return p  
}

func (p *ArrayStack) Deck() CardSlice {
  return p.cards[0:p.deck]
}

// 0|--deck--|--hand--|-used-|--discard--|size

func (p *ArrayStack) Hand() CardSlice {
  return p.cards[p.deck:p.discard]
}

func (p *ArrayStack) Discard() CardSlice {
  return p.cards[p.discard:p.size]
}

func (p *ArrayStack) Gain(c *Card) {
  p.cards[p.size] = c
  p.size++
}

func (p *ArrayStack) recycleDiscard() {
  // Error catching for out of bounds.
  defer func() {
    if r := recover(); r != nil {
      fmt.Printf("Deck: %d, Discard: %d, Size: %d\n", p.deck, p.discard, p.size)
    }
  }()

  k := p.size - 1
  for i := 0; i < p.discard && k > i; i++ {
    p.cards[i], p.cards[k] = p.cards[k], p.cards[i]
    k--
  }
  p.deck = p.size - p.discard
  p.discard = p.size
  deck := p.Deck()
  deck.Shuffle()
}

func (p *ArrayStack) Use(cname string) bool {
  return true
}

func (p *ArrayStack) Draw(count int) bool {
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
    p.recycleDiscard()
  }
  for count > 0 && p.deck > 0 {
    count--
    p.deck--
  }
  return true
}

func (p *ArrayStack) DoDiscard(count int) {
  p.discard -= count
}

func (p *ArrayStack) swap(i, j int) {
  p.cards[i], p.cards[j] = p.cards[j], p.cards[i]
}

// Shuffles the pile of cards randomly.
func (cs *CardSlice) Shuffle() {
  for i := len(*cs) - 1; i > 0; i-- {
    if j := r.Intn(i + 1); i != j {
      cs.Swap(i, j)
    }
  }
}


//DiscardOne attempts to discard a card with the specified
//name.  Returns true on success.
func (p *ArrayStack) DiscardOne(name string) bool {
  for i := p.discard - 1; i >= p.deck; i-- {
    if p.cards[i].Name == name {
      p.swap(i, p.discard)
      p.discard--
      return true
    }
  }
  return false
}

type CopyStack struct {
  deck Queue
  hand Queue
  discard Queue
}

func NewCopyStack() CopyStack {
  return CopyStack{deck:StartingDeck()}
}

func (c *CopyStack) Deck() CardSlice {
  return CardSlice(c.deck)
}

func (c *CopyStack) Hand() CardSlice {
  return CardSlice(c.hand)
}

func (c *CopyStack) Discard() CardSlice {
  return CardSlice(c.discard)
}

func (c *CopyStack) DoDiscard(count int) {
  for i := 0; i < count; i++ {
    c.discard.Add(c.hand.Pop())
  }
}

func (c *CopyStack) Draw(count int) bool {
  if count > (len(c.deck) + len(c.discard)) {
    return false
  }
  for i := 0; i < count; i++ {
    if (len(c.deck) == 0) {
      for _, _ = range c.discard {
        c.deck.Add(c.discard.Pop())
      }
      c.deck.Shuffle()
    }
    c.hand.Add(c.deck.Pop())
  }
  return true
}
func (c *CopyStack) DiscardOne(name string) bool {
  return true
}


func (cs *CopyStack) Gain(c *Card) {
  cs.discard.Add(c)
}

func StartingDeck() Queue {
  var deck Queue
  for i := 0; i < 7; i++ {
    deck.Add(&Copper)
  }
  for i := 0; i < 3; i++ {
    deck.Add(&Estate) 
  }
  return deck
}

type Queue []*Card

//Removes and returns the last card in the pile.
func (p *Queue) Pop() *Card {
  s := *p
  if len(s) == 0 {
    return &Card{}
  }
  card := s[len(s)-1]
  *p = s[:len(s)-1]
  return card
}

func (p *Queue) Len() int {
  return len(*p)
}

// Shuffles the pile of cards randomly.
func (p *Queue) Shuffle() {
  for i := p.Len() - 1; i > 0; i-- {
    if j := r.Intn(i + 1); i != j {
      p.Swap(i, j)
    }
  }
}

//Finds the specified card, removes it from the pile and returns it.
// Returns nil if the card is not found in the pile.
func (p *Queue) Remove(target *Card) *Card {
  for i, card := range *p {
    if card.Name == target.Name {
      p.Swap(i, p.Len() - 1)
      return p.Pop()
    }
  }
  return &Card{}
}

func (p *Queue) Swap(i, j int) {
  s := *p
  //fmt.Printf("i: %d, j: %d\n", i, j)
  //fmt.Printf("s[i]: %v, s[j] %v\n", s[i], s[j])
  s[i], s[j] = s[j], s[i]
  //fmt.Printf("s[i]: %v, s[j] %v\n", s[i], s[j])
  *p = s
  //fmt.Println(p)
}

func (p Queue) Less(i, j int) bool {
  return p[i].Type < p[j].Type || p[i].CoinPrice < p[j].CoinPrice || p[i].Name < p[j].Name
}

func (p *Queue) Sort() {
  sort.Sort(p)
}

func (p Queue) Count(cardName string) int {
  count := 0
  for _, card := range p {
    if card.Name == cardName {
      count++
    }
  }
  return count
}

//Adds the specified card to the end of the pile.
func (p *Queue) Add(card *Card) {
  s := *p
  *p = append(s, card)
}

func (p *Queue) AddAll(other Queue) {
  s := *p
  *p = append(s, other...)
}

