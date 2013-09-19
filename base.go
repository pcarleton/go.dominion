package dominion

import (
  "sort"
  "fmt"

)

var BaseCards = map[string]Card{
  "Copper": Card{"Copper", 0, 1, 0, TREASURE},
  "Silver": Card{"Silver", 0, 2, 3, TREASURE},
  "Gold": Card{"Gold", 0, 3, 6, TREASURE},
  "Estate":Card{"Estate", 1, 0, 2, VICTORY},
  "Duchy":Card{"Duchy", 3, 0, 5, VICTORY},
  "Province":Card{"Province", 6, 0, 8, VICTORY},
  "Smithy":Card{"Smithy", 0, 0, 4, ACTION},
  "Village":Card{"Village", 0, 0, 3, ACTION},
}


type Game struct {
  Players []*Player
  Stacks map[string] int
  turns int
}


func NewGame(p ...*Player) Game {
  stacks := map[string]int{
    "Copper":60,
    "Silver":40,
    "Gold":30,
    "Estate":24,
    "Province":12,
    "Duchy":12,
    "Smithy":10,
    "Village":10,
  }
  return Game{Players:p, Stacks:stacks}
}

func (g *Game) Play() {
  // Deal out decks and draw first hand
  for _, player := range g.Players {
    player.Deck = StartingDeck()
    player.Hand = nil
    player.Discard = nil
    player.Deck.Shuffle()
    player.Draw(5)
  }
  for !g.endCondition() {
    g.turns++
    for _, p := range g.Players {
      p.PlayTurn(g)
      if g.endCondition() {
        break
      }
    }
  }
  g.DetermineWinner()
}

func (g *Game) endCondition() bool {
  zeroCount := 0
  for _, count := range g.Stacks {
    if count == 0 {
      zeroCount++
    }
  }
  return g.Stacks["Province"] == 0 || zeroCount >= 3
}

type Score struct {
  Name string
  Place int
  VP int
  Turns int
  Order int
}

type scoreSorter struct {
  ss []Score
}

func (c scoreSorter) Swap(i, j int) {
  c.ss[i], c.ss[j] = c.ss[j], c.ss[i]
}

func (c scoreSorter) Len() int {
  return len(c.ss)
}

func (c scoreSorter) Less(i, j int) bool {
  return c.ss[i].VP < c.ss[j].VP
}

func (g *Game) DetermineWinner() []Score {
  scores := make([]Score, len(g.Players))
  for i, p := range g.Players {
    scores[i] = Score{VP:p.GetVictoryPoints(), Turns:g.turns, Order:i}
  }
  fmt.Println(scores)
  sort.Sort(&scoreSorter{ss:scores})
  fmt.Println(scores)
  for i, score := range scores {
    score.Place = i
  }
  return scores
}

func StartingDeck() Pile {
  var deck Pile
  for i := 0; i < 7; i++ {
    deck.Add(BaseCards["Copper"])
  }
  for i := 0; i < 3; i++ {
    deck.Add(BaseCards["Estate"]) 
  }
  return deck
}

func smithyFunc(turn *Turn) {
  turn.P.Draw(3)  
}

func villageFunc(turn *Turn) {
  turn.P.Draw(1)
  turn.Actions += 2
}

var ActionFunctions = map[string](func(*Turn)){
  "Smithy": smithyFunc,
  "Village": villageFunc,
}
