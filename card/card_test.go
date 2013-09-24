package card

import (
  "testing"
)

var GetStack func() Stack

func init() {
GetStack = func() Stack {
  p := NewArrayStack()
  return &p
}
}

func TestNew(t *testing.T) {

  p := GetStack()
  t.Logf("%+v\n", p)
  t.Logf("%+v\n", p.Deck())
  t.Logf("%+v\n", p.Hand())
  t.Logf("%+v\n", p.Discard())
  if !CheckPiles(p, 10, 0, 0) {
    t.Error("Improper new deck")
  }
}

func TestGain(t *testing.T) {
  p := GetStack()
  p.Gain(&Copper)
  if !CheckPiles(p, 10, 0, 1) {
    t.Error("Gain didn't gain a card")
  }
}

func TestDraw(t *testing.T) {
  p := GetStack()
  p.Draw(5)
  if !CheckPiles(p, 5, 5, 0) {
    t.Log(p.Deck(), p.Hand())
    t.Error("Hand wasn't drawn properly")
  }
  p.Draw(2)
  if !CheckPiles(p, 3, 7, 0) {
    t.Error("Couldn't draw through deck")

  }
  // Over drawing is not okay.
  if ok := p.Draw(6); ok {
    t.Log(p.Deck())
    t.Error("Didn't return false on overdraw")
  }
}

func CheckPiles(p Stack, d, h, dc int) bool {
  return len(p.Deck()) == d && len(p.Hand()) == h && len(p.Discard()) == dc
}

func TestDiscard(t *testing.T) {
  p := GetStack()
  p.Draw(5)
  p.DoDiscard(5)
  if !CheckPiles(p, 5, 0, 5) {
    t.Error("Improper discard")
  }
}

func BenchmarkStartArrayDeck(b *testing.B) {
  for i := 0; i < b.N; i++ {
    p := NewArrayStack()
    p.Draw(5)
  }
}

func BenchmarkStartCopyStack(b *testing.B) {
  for i := 0; i < b.N; i++ {
    p := NewCopyStack()
    p.Draw(5)
  }
}


func BenchmarkArrayStackGamePlay(b *testing.B) {
  for i := 0; i < b.N; i++ {
    p := NewArrayStack()
    for j := 0; j < 20; j++ {
      p.Draw(5)
      p.Gain(&Copper)
      p.DoDiscard(5)
    }
  }
}

func BenchmarkCopyStackGamePlay(b *testing.B) {
  for i := 0; i < b.N; i++ {
    p := NewCopyStack()
    for j := 0; j < 20; j++ {
      p.Draw(5)
      p.Gain(&Copper)
      p.DoDiscard(5)
    }
  }
}
