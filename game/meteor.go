package game

import (
	"cruiser/assets"
	"math/rand"
)

const (
	rotationSpeedMin = -2
	rotationSpeedMax = 2
)

type Meteor struct {
	Entity
	weight int
}

func FirstMeteor() *Meteor {
	m := NewMeteor()
	m.Y = 100
	return m
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
	m.X = 800 + float64(m.getWidth()/2)
	m.Y = m.getWidth()/2 + rand.Float64()*(600-float64(m.getWidth()))
	return m
}

func (m *Meteor) Update(speed float64) {
	m.X -= GetFrameValue(speed)
}
