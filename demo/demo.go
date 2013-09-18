package main

import (
  dominion "github.com/pcarleton/go.dominion"
  "fmt"
)


func main() {
  fmt.Println("Welcome to the Go Dominion Simulator")
  //piles := map[dominion.Card]int{ dominion.Copper: 30, dominion.Estate: 5}
  p1 := dominion.NewRobotPlayer("Robot 1")
  p2 := dominion.NewRobotPlayer("Robot 2")
  p3 := dominion.NewBMSPlayer("BMS 3")

  winCounts := map[string]int{}
  var g dominion.Game
  for i := 0; i < 1000; i++ {
    g = dominion.NewGame(&p1, &p2, &p3)
    g.Play()
    winCounts[g.DetermineWinner().Name]++

    g = dominion.NewGame(&p3, &p1, &p2)
    g.Play()
    winCounts[g.DetermineWinner().Name]++

    g = dominion.NewGame(&p2, &p3, &p1)
    g.Play()
    winCounts[g.DetermineWinner().Name]++
  }
  fmt.Println("#####FINALS######")

  for p, count := range winCounts {
    fmt.Printf("%s: %d\n", p, count)
  }

}
