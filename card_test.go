package dominion

import (
  "testing"
)

func TestStartingDeck(t *testing.T) {
  copperCount := 0
  estateCount := 0
  for _, card := range(StartingDeck()) {
    if card.Name == "Copper" { copperCount++ }
    if card.Name == "Estate" { estateCount++ }
  }
  if copperCount != 7 {
    t.Errorf("Needed 7 coppers, saw %d", copperCount)
  }
  if estateCount != 3 {
    t.Errorf("Needed 3 estates, saw %d", estateCount)
  }
}

func TestPop(t *testing.T) {
  var p Pile
  p.Add(BaseCards["Copper"])
  if len(p) != 1 {
    t.Error("Didn't add card")
  }
  c := p.Pop()
  if c.Name != "Copper" {
    t.Error("Pop didn't work")
  }
}

func TestRemove(t *testing.T) {
  deck := StartingDeck()
  card := deck.Remove(BaseCards["Copper"])
  if card.Name != "Copper" {
    t.Errorf("Removed %v instead of copper", card)
  }
  if deck.Len() != 9 {
    t.Errorf("Did not remove card from deck.")
  }
}
