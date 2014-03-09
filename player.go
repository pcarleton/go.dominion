package dominion
// Player struct and methods

import (
  "fmt"
)

type BuyFunc func(i int,t *Turn) Card
type ActionFunc func(p Pile) Card

type Strat struct {
  Name string
  Buy BuyFunc
  Act ActionFunc
}

type Player struct {
  Name string
  Deck Pile
  Discard Pile
  Hand Pile
  Plan Strat
}


type GameGenerator struct {
  sts []Strat
}

func NewGameGenerator(ss ...Strat) GameGenerator {
  return GameGenerator{ sts:ss}
}

func (s Strat) GetPlayer() *Player {
  return &Player{Name:s.Name, Plan:s}
}

func (g *GameGenerator) Generate() Game {
  //Rotate starting order
  first := g.sts[0]
  for i := 1; i < len(g.sts); i++ {
    g.sts[i-1] = g.sts[i]
  }
  g.sts[len(g.sts)-1] = first

  //Create players
  ps := make([]*Player, len(g.sts))
  for i, strat := range g.sts {
    ps[i] = strat.GetPlayer()
  }
  return NewGame(ps...)
}

var HumanStrat = Strat{Name:"Human", Buy:HumanSelect, Act:HumanSelectAction }
var BigMoneyStrat = Strat{Name:"BigMoney", Buy:RobotSelect }
var BMSStrat = Strat{Name:"BigMoneySmith", Buy:BMSSelect, Act:DumbSelectAction}

func NewHumanPlayer(name string) Player {
  return Player{Name: name, Plan: HumanStrat}
}

func NewBMSPlayer(name string) Player {
  return Player{Name: name, Plan: BMSStrat}
}

func NewRobotPlayer(name string) Player {
  return Player{Name: name, Plan: BigMoneyStrat}
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

func (p *Player) DoDiscard(cards Pile) {
}

func (p *Player) PlayTurn(game *Game) {
  //fmt.Printf("#### %s ####\n", p.Name)
  t := Turn{P:p, G:game, Actions:1, Buys:1}  
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
  //fmt.Println("****Buy Phase****")
  //p.PrintHand()
  monies := 0
  for _, card := range(p.Hand) {
    monies += card.CoinValue
  }
  // For now, only do 1 buy.
  done := false
  for !done {
    //selection := GetSelection(i)
    //selected := BaseCards[options[selection]]
    selected := p.Plan.Buy(monies, turn)
    if selected.Name == "" {
      done = true
      break
    }
    if (selected.CoinPrice <= monies) {
      turn.G.Stacks[selected.Name]--
      p.Gain(selected)
      done = true
      //fmt.Printf("Bought: %s\n", selected.Name)
    } else {
      //fmt.Println("You don't have enough money for that card.")
    }
  }
}

func BMSSelect(monies int, turn *Turn) Card {
  var priority []string
  if (turn.P.getAllCards().Count("Smithy") < 2) {
    priority = []string{"Smithy", "Province", "Gold", "Silver", "Copper"}
  } else {
    priority = []string{ "Province", "Gold", "Silver", "Copper"}
  }
  for _, name := range priority {
    if monies >= BaseCards[name].CoinPrice && turn.G.Stacks[name] > 0 {
      return BaseCards[name]
    }
  }
  return BaseCards["Copper"]
}


func RobotSelect(monies int, turn *Turn) Card {
  priority := []string{ "Province", "Gold", "Silver", "Copper"}
  for _, name := range priority {
    if monies >= BaseCards[name].CoinPrice && turn.G.Stacks[name] > 0 {
      return BaseCards[name]
    }
  }
  return BaseCards["Copper"]
}

func HumanSelect(monies int, turn *Turn) Card {
  fmt.Printf("Coins: %d, Buys: %d\n", monies, turn.Buys)
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
  selection := GetSelection(i)
  return BaseCards[options[selection]]
}


func HumanSelectAction(actionCards Pile) Card {
  fmt.Println("Select Action to play")
  return SelectCard(actionCards)
}

func DumbSelectAction(actionCards Pile) Card {
  return actionCards[0]
}

func (p *Player) DoActionPhase(turn *Turn) {
  //fmt.Println("****Action Phase****")
  played := Pile{}
  for ; turn.Actions > 0; {
    // Get all actions from hand
    //fmt.Printf("Actions Left: %d\n", turn.Actions)
    //p.PrintHand()
    actionCards := Pile{}
    for _, card := range p.Hand {
      if card.Type == ACTION {
        actionCards.Add(card)
      }
    }
    if actionCards.Len() == 0  {
      //fmt.Println("No actions to play.")
      return
    }
    // Choose an action.
    actionCard := p.Plan.Act(actionCards)
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

func (p *Player) GetVictoryPoints() int {
  vp := 0
  for _, card := range p.getAllCards() {
    vp += card.VictoryValue
  }
  return vp
}

func (p *Player) getAllCards() Pile {
  allCards := Pile{}
  allCards.AddAll(p.Hand)
  allCards.AddAll(p.Discard)
  allCards.AddAll(p.Deck)
  return allCards
}

func (p *Player) String() string {
  return fmt.Sprintf("Player %s-> VP: %d Deck: %v Discard: %v", p.Name, p.Points(), p.Deck, p.Discard)
}

