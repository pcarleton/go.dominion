package main

import (
  dominion "github.com/pcarleton/go.dominion"
  "fmt"
)

var trials = 1

func main() {
  fmt.Println("Welcome to the Go Dominion Simulator")
  //piles := map[dominion.Card]int{ dominion.Copper: 30, dominion.Estate: 5}
  p1 := dominion.NewRobotPlayer("Robot 1")
  p2 := dominion.NewRobotPlayer("Robot 2")
  p3 := dominion.NewBMSPlayer("BMS 3")

  var g dominion.Game
  for i := 0; i < trials; i++ {
    g = dominion.NewGame(&p1, &p2, &p3)
    g.Play()
  fmt.Println("#####FINALS######")
    fmt.Printf("%+v\n", g.DetermineWinner())

  }
}

func RunGame(p1, p2, p3 dominion.Player, ch chan(string)) {
  

}
