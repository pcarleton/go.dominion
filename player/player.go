package player

import (
  "github.com/pcarleton/go.dominion/card"
)


type Player struct {
  Name string
  Cards card.Stack
}

type Strategy struct {
  // Choose an action card given the players current cards.
  ChooseAction(card.Stack) &Card
  ChooseBuy(card.Stack) &Card
}
