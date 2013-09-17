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
  t := Turn{P:p, G:game, Actions:2, Buys:2}  
  p.DoActionPhase(&t)
  p.DoBuyPhase(&t)
  p.DoCleanUp()
}

func (p *Player) PrintHand() {
  p.Hand.Sort()
  for _, card := range(p.Hand) {
    fmt.Printf("%s ", card.Name)
  }
  fmt.Printf("\n")
}

func (p *Player) DoBuyPhase(turn *Turn) {
  fmt.Println("****Buy Phase****")
  p.PrintHand()
  monies := 0
  for _, card := range(p.Hand) {
    monies += card.CoinValue
  }
  // For now, only do 1 buy.
  fmt.Printf("Coins: %d, Buys: %d", monies, turn.Buys)
  fmt.Println("#  Name   Available    Price")
  i := 0
  options := []string{}
  for card, count := range(turn.G.Stacks) {
    if count > 0 {
      fmt.Printf("%d. %s: %d %d\n", i, card, count, BaseCards[card].CoinPrice)
      options = append(options, card)
      i++
    }
  }
  done := false
  for !done {
    selection := GetSelection(i)
    selected := BaseCards[options[selection]]
    if (selected.CoinPrice <= monies) {
      turn.G.Stacks[selected.Name]--
      p.Gain(selected)
      done = true
    } else {
      fmt.Println("You don't have enough money for that card.")
    }
  }
}

func (p *Player) DoActionPhase(turn *Turn) {
  fmt.Println("****Action Phase****")
  played := Pile{}
  for ; turn.Actions > 0; {
    // Get all actions from hand
    fmt.Printf("Actions Left: %d", turn.Actions)
    p.PrintHand()
    actionCards := Pile{}
    for _, card := range p.Hand {
      if card.Type == ACTION {
        actionCards.Add(card)
      }
    }
    if actionCards.Len() == 0  {
      fmt.Println("No actions to play.")
      return
    }
    // Choose an action.
    fmt.Println("Select Action to play")
    actionCard := SelectCard(actionCards)
    actionCard.playAction(turn)
    turn.Actions--
    played.Add(p.Hand.Remove(actionCard))
  }
  p.Hand.AddAll(played)
}

func (p *Player) DoCleanUp() {
  p.Discard.AddAll(p.Hand)
  p.Hand = nil
  p.Draw(5)
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

func (p *Player) String() string {
  return fmt.Sprintf("Player %s-> VP: %d Deck: %v Discard: %v", p.Name, p.Points(), p.Deck, p.Discard)
}

