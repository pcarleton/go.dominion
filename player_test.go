package dominion

import(
  "testing"
)

var emptyPlayer = Player{Name:"empty"}

func TestGain(t *testing.T) {
  expected := []Card{BaseCards["Copper"]}
  emptyPlayer.Gain(BaseCards["Copper"])
  if emptyPlayer.Discard[0] != BaseCards["Copper"] || len(emptyPlayer.Discard) > 1 {
    t.Errorf("Expected: %v, but got %v", expected, emptyPlayer.Discard)
  }
}

func TestDraw(t *testing.T) {
  player := Player{Name:"test", Deck:Pile{BaseCards["Estate"], BaseCards["Estate"]}, Discard:Pile{BaseCards["Copper"], BaseCards["Copper"]}}

  //First draw.
  player.Draw(1)
  if player.Deck.Len() != 1 || player.Hand.Len() != 1 {
    t.Errorf("Didn't draw first card correctly.")
  }

  //Draw again
  player.Draw(1)
  if player.Deck.Len() != 0 || player.Hand.Len() != 2 {
    t.Errorf("Didn't draw 2nd card correctly.")
  }

  //Draw again with shuffle from discard.
  player.Draw(1)
  if player.Deck.Len() != 1 || player.Hand.Len() != 3 || player.Discard.Len() != 0 {
    t.Errorf("Didn't draw with shuffle correctly: %v, %v, %v", player.Deck, player.Hand, player.Discard)
  }
}

func TestSelection(t *testing.T) {
  c := GetSelection(3)
  if c == -1 {
    t.Errorf("Error\n")
  }
}
