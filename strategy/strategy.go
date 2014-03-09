package strategy

import (
  "github.com/pcarleton/go.dominion/card"
)

type Strategy interface {
  // Choose an action card given the players current cards.
  ChooseAction(card.Stack) &Card
  ChooseBuy(card.Stack) &Card
}


func ChooseRandomAction(stack card.Stack) &Card {
  stack.Hand().GetActions()

}


/* Player should be responsible for state of the game.  He does all the mechanical work and provides options.  Strategy makes the decisions about what to do based on the state of the game.  However, how do I handle restricted actions? The Stack handles categories of cards, so it should also handle a used stack. */
