package game

import (
	"cruiser/assets"
	"math"
)

type Player struct {
	Entity
	Vy                   float64
	accelerationConstant float64
}

func NewPlayer(accelerationConstant float64, x float64, y float64) *Player {
	p := &Player{}
	p.accelerationConstant = accelerationConstant
	p.Vy = 0
	p.X = x
	p.Y = y
	p.sprite = assets.PlayerSprite
	return p
}

func (p *Player) Update() {
	p.Y += GetFrameValue(p.Vy)
}

func (p *Player) UpdateVerticalVelocity(meteor *Meteor) {
	p.Vy += ((p.accelerationConstant) / (meteor.GetDistance(&p.Entity))) * ((meteor.Y - p.Y) / math.Sqrt(meteor.GetDistance(&p.Entity)))
}
