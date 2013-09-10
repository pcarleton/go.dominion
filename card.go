package dominion

type Card struct {
  Name string
  VictoryValue int
  CoinValue int
  CoinPrice int
}

Copper := Card{"Copper", 0, 1, 0}
Estate := Card{"Estate", 1, 0, 2}
