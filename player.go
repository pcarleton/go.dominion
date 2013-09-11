package dominion

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

func (p *Player) discard(cards Pile) {
  

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

