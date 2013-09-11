package strategy


type Interface interface {
}

type Strategy struct {
  p *Player
}


func NewStrategy(p *Player) Strategy {
  return Strategy{p: p}
}


