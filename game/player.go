package game

type Player struct {
	Entity
	Vx float64
	Vy float64
}

func (p *Player) Update(ax float64, ay float64) {
	p.X += GetFrameValue(p.Vx)
	p.Y += GetFrameValue(p.Vy)
	p.Vx += GetFrameValue(ax)
	p.Vy += GetFrameValue(ay)
}
