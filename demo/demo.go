package main

import (
  dominion "github.com/pcarleton/go.dominion"
  "fmt"
)


func main() {
  fmt.Println("Welcome to the Go Dominion Simulator")
  //piles := map[dominion.Card]int{ dominion.Copper: 30, dominion.Estate: 5}
  p1 := dominion.Player{"1", dominion.StartingDeck(), nil, nil, nil}
  //p2 := dominion.Player{"2", dominion.StartingDeck(), nil, nil, nil}

  p1.Gain(dominion.Smithy)
  p1.Draw(4)
  p1.Hand.Add(dominion.Smithy)
  fmt.Println(p1.Hand)
  p1.DoActionPhase()

  fmt.Println(p1.Hand)
}
