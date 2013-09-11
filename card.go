package dominion

type CardType int
const (
  TREASURE CardType = iota
  VICTORY
  ACTION
)

type Card struct {
  Name string
  VictoryValue int
  CoinValue int
  CoinPrice int
  Type CardType
}

var Copper = Card{"Copper", 0, 1, 0, TREASURE}
var Estate = Card{"Estate", 1, 0, 2, VICTORY}

func startingDeck() []Card {
  var deck []Card
  for i := 0; i < 7; i++ {
    deck = append(deck, Copper)
  }
  for i := 0; i < 3; i++ {
    deck = append(deck, Estate)
  }
  return deck
}
