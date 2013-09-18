package dominion

import (
  //"fmt"

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
  // Deal out decks and draw first hand
  for _, player := range p {
    player.Deck = StartingDeck()
    player.Hand = nil
    player.Discard = nil
    player.Deck.Shuffle()
    player.Draw(5)
  }
  return Game{Players:p, Stacks:stacks}
}

func (g *Game) Play() {
  for !g.endCondition() {
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

func (g *Game) DetermineWinner() Player {
  maxPoints := 0
  var winner Player
  //fmt.Println("Final Scores")
  for _, p := range g.Players {
    vp := p.GetVictoryPoints()
    //fmt.Printf("%s: %d\n", p.Name, vp)
    if vp > maxPoints {
      maxPoints = vp
      winner = *p
    }
  }
  //fmt.Printf("Winner is %s\n", winner.Name)
  return winner
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
