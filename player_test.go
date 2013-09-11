package dominion

import(
  "testing"
)

var emptyPlayer = Player{"empty", nil, nil, nil}

func TestGain(t *testing.T) {
  expected := []Card{Copper}
  emptyPlayer.Gain(Copper)
  if emptyPlayer.Discard[0] != Copper || len(emptyPlayer.Discard) > 1 {
    t.Errorf("Expected: %v, but got %v", expected, emptyPlayer.Discard)
  }
}


