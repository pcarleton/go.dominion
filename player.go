package dominion
// Player struct and methods

import (
  "fmt"
)

type Player struct {
  Name string
  Deck Pile
  Discard Pile
  Hand Pile
  Plan Strat
}

type Strat interface{}

func (p *Player) Gain(card Card) {
  p.Discard.Add(card)
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

func (p *Player) draw() {
  p.Hand.Add(p.Deck.Pop())
}

func (p *Player) Draw(num int) {
  for i := 0; i < num; i++ {
    if len(p.Deck) == 0 {
      p.Deck = p.Discard
      p.Discard = nil
      p.Deck.Shuffle()
    }
    p.draw()
  }
}

func (p *Player) DoDiscard(cards Pile) {
}

func (p *Player) PlayTurn(game *Game) {
  t := Turn{P:p, G:g, Actions:1, Buys:1}  
  p.DoActionPhase(t)

}

func (p *Player) DoActionPhase(turn *Turn) {
  // Get all actions from hand
  actionCards := Pile{}
  for _, card := range p.Hand {
    if card.Type == ACTION {
      actionCards.Add(card)
    }
  }
  // Choose an action.
  actionCard := SelectCard(actionCards)
  actionCard.playAction(
}

func SelectCard(pile Pile) Card{
  for i, card := range(pile) {
    fmt.Printf("%d. %s\n", i, card.Name)
  }
  choice := GetSelection(len(pile))
  return pile[choice]
}

// Get selection gets an input from the command line
func GetSelection(opts int) int {
  fmt.Printf("Enter selection (%d-%d):\n", 0, opts)
  var choice int
  n, err := fmt.Scan(&choice)
  if n == 0 || err != nil {
    fmt.Println(err)
  }
  return choice
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

