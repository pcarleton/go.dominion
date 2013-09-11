package dominion

import (
  "testing"
)

func TestStartingDeck(t *testing.T) {
  copperCount := 0
  estateCount := 0
  for _, card := range(startingDeck()) {
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
