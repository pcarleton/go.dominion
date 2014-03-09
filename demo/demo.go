package main

import (
  dominion "github.com/pcarleton/go.dominion"
  "fmt"
  "time"
  "runtime/pprof"
  "flag"
  "os"
  "log"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var trials int

func init() {
  flag.IntVar(&trials, "trials", 100, "Number of games to run")

}

func main() {
  flag.Parse()
  if *cpuprofile != "" {
    f, err := os.Create(*cpuprofile)
      if err != nil {
        log.Fatal(err)
      }
    pprof.StartCPUProfile(f)
    defer pprof.StopCPUProfile()
  } 
  fmt.Println("Welcome to the Go Dominion Simulator")
    //piles := map[dominion.Card]int{ dominion.Copper: 30, dominion.Estate: 5}
    gg := dominion.NewGameGenerator(dominion.BigMoneyStrat, dominion.BigMoneyStrat, dominion.BMSStrat) 

  scoreChan := make(chan []dominion.Score)
  resultChan := make(chan map[string][]dominion.Score)

  t0 := time.Now()
  go CatchScores(scoreChan, resultChan)
  for i := 0; i < trials; i++ {
    go RunGame(gg.Generate(), scoreChan)
  }
  
  scoreMap := <-resultChan

  for name, scores := range scoreMap {
    fmt.Println(name)
    fmt.Printf("%+v\n", AggregateScores(scores))
  }
  t1 := time.Now()
  fmt.Printf("The simulation took %v to run.\n", t1.Sub(t0))
}

func RunGame(g dominion.Game, c chan []dominion.Score) {
  g.Play()
  c <- g.DetermineWinner()
}

func CatchScores(c chan []dominion.Score, r chan map[string][]dominion.Score) {
  scoreMap := make(map[string][]dominion.Score)
  caught := 0
  var scores []dominion.Score
  for caught < trials {
    scores = <-c
    for _, score := range scores {
      scoreMap[score.Name] = append(scoreMap[score.Name], score)
    }
    caught++
  }
  r <- scoreMap
}

type ScoreAgg struct {
  Wins int
  WinRate float64
  Games int
  AvgTurns float64
  AvgVP float64
}

func AggregateScores(scores []dominion.Score) ScoreAgg {
  totalVP := 0
  totalTurns := 0
  scagg := ScoreAgg{}
  for _, score := range scores {
    scagg.Games++
    if score.Place == 1 { scagg.Wins++ }
    totalTurns += score.Turns
    totalVP += score.VP
  }
  scagg.AvgTurns = float64(totalTurns)/float64(scagg.Games)
  scagg.AvgVP = float64(totalVP)/float64(scagg.Games)
  scagg.WinRate = float64(scagg.Wins)/float64(scagg.Games)
  return scagg
}
