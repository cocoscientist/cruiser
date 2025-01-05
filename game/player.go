package game

import (
	"cruiser/assets"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
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
	if (&Meteor{}) != meteor {
		p.Vy += ((p.accelerationConstant * float64(meteor.weight)) / (meteor.GetDistance(&p.Entity))) * ((meteor.Y - p.Y) / math.Sqrt(meteor.GetDistance(&p.Entity)))
	} else {
		p.Vy += 0
	}
}

func (p *Player) Draw(screen *ebiten.Image, engineOn bool) {
	op := &ebiten.DrawImageOptions{}
	halfW := float64(p.getWidth()) / 2
	halfH := float64(p.getHeight()) / 2
	if !engineOn {
		op.GeoM.Translate((2*p.X)-(halfW*2)-40, (2*p.Y)-(halfH*2)+35)
		op.GeoM.Scale(0.5, 0.5)
		screen.DrawImage(assets.ExhaustSprite, op)
		op.GeoM.Scale(2.0, 2.0)
		op.GeoM.Translate((2*halfW)-(2*p.X)+40, (2*halfH)-(2*p.Y)-35)
	}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Translate(p.X, p.Y)
	screen.DrawImage(p.sprite, op)
}

func (p *Player) ResetVerticalVelocity() {
	p.Vy = 0
}

func (p *Player) OutOfBounds() bool {
	return (p.Y < (0)) || (p.Y > (600))
}
