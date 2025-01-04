package game

import (
	"cruiser/assets"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	rotationSpeedMin = -2
	rotationSpeedMax = 2
	baseSpeed        = 200
)

type Meteor struct {
	Entity
	rotation      float64
	rotationSpeed float64
	weight        int
}

func NewMeteor() *Meteor {
	bw := rand.Float64()
	m := &Meteor{}
	if bw < 0.5 {
		m.sprite = assets.SmallMeteorSprites[rand.Intn(len(assets.SmallMeteorSprites))]
		m.weight = 1
	} else {
		m.sprite = assets.LargeMeteorSprites[rand.Intn(len(assets.LargeMeteorSprites))]
		m.weight = 8
	}
	m.X = 800 - float64(m.sprite.Bounds().Dx()/2)
	m.Y = rand.Float64() * (600 - float64(m.sprite.Bounds().Dy()/2))
	m.rotationSpeed = rotationSpeedMin + rand.Float64()*(rotationSpeedMax-rotationSpeedMin)
	return m
}

func (m *Meteor) Update(speed float64) {
	m.X -= GetFrameValue(speed)
	m.rotation += GetFrameValue(m.rotationSpeed)
}

func (m *Meteor) Draw(screen *ebiten.Image) {
	halfW := m.getWidth() / 2
	halfH := m.getHeight() / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Rotate(m.rotation)

	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Translate(m.X, m.Y)

	screen.DrawImage(m.sprite, op)
}
