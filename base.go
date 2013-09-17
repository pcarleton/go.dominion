package dominion

var Copper = Card{"Copper", 0, 1, 0, TREASURE}
var Estate = Card{"Estate", 1, 0, 2, VICTORY}
var Smithy = Card{"Smithy", 0, 0, 4, ACTION}

var BaseCards = map[string]Card{
  "Copper":Copper,
  "Estate":Estate,
  "Smithy":Smithy,
}


type Game struct {
  Players []*Player
  Stacks map[string] int
}


func NewGame(p ...*Player) Game {
  stacks := map[string]int{
    "Copper":50,
    "Estate":12,
    "Smithy":10,
  }
  return Game{Players:p, Stacks:stacks}
}

