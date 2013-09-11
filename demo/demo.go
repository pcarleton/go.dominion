package main

import (
  "github.com/pcarleton/go.dominion/dominion"
  "fmt"
)


func main() {
  fmt.Println("Welcome to the Go Dominion Simulator")
  piles := map[dominion.Card]int{ dominion.Copper: 30, dominion.Estate: 5}
  deck := make([]dominion.Card, 10, 10)
  for i := 0; i < 7; i++ {
    deck = append(deck, dominion.Copper)
    //deck[i] = dominion.Copper
  }
  for i := 7; i < 10; i++ {
    deck = append(deck, dominion.Estate)
    //deck[i] = dominion.Estate
  }
  fmt.Println(deck)
  deck2 := make([]dominion.Card, len(deck))
  copy(deck2, deck)
  p1 := Player{"1", deck, nil, nil}
  p2 := Player{"2", deck2, nil, nil}
  for piles[dominion.Estate] > 0 {
    piles[dominion.Estate]--
    p1.Gain(dominion.Estate)
  }
  fmt.Println(p1.String())
  fmt.Println(p2.String())

}
