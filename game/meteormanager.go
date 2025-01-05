package game

import (
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type MeteorManager struct {
	Meteors              []*Meteor
	baseSpeed            float64
	accelerationConstant float64
	score                int
}

func NewMeteors(baseSpeed float64, accelerationConstant float64) *MeteorManager {
	mm := &MeteorManager{}
	mm.Meteors = append(mm.Meteors, FirstMeteor())
	for i := 1; i < 7; i++ {
		mm.Meteors = append(mm.Meteors, MeteorCreator(float64(300+i*400)))
	}
	mm.baseSpeed = baseSpeed
	mm.accelerationConstant = accelerationConstant * 0.2
	mm.score = 0
	return mm
}

func (mm *MeteorManager) UpdateAllMeteors() {
	for i, meteor := range mm.Meteors {
		meteor.Update(mm.baseSpeed)
		if meteor.X < -1*meteor.getWidth()/2 {
			mm.removeMeteor(i)
			mm.addMeteor()
		}
	}
}

func (mm *MeteorManager) UpdateSpeed(meteor *Meteor, p *Player) {
	if (&Meteor{}) != meteor {
		mm.baseSpeed += ((mm.accelerationConstant * float64(meteor.weight)) / (meteor.GetDistance(&p.Entity))) * ((meteor.X - p.X) / math.Sqrt(meteor.GetDistance(&p.Entity)))
	}
}

func (mm *MeteorManager) GetClosestMeteor(x float64) *Meteor {
	m := &Meteor{}
	for _, meteor := range mm.Meteors {
		if meteor.X > x {
			m = meteor
			break
		}
	}
	return m
}

func (mm *MeteorManager) DrawAllMeteors(screen *ebiten.Image) {
	for _, meteor := range mm.Meteors {
		meteor.Draw(screen)
	}
}

func (mm *MeteorManager) addMeteor() {
	m := mm.Meteors[len(mm.Meteors)-1]
	mm.Meteors = append(mm.Meteors, MeteorCreator(m.X+360+rand.Float64()*40))
}

func (mm *MeteorManager) removeMeteor(position int) {
	mm.Meteors = append(mm.Meteors[:position], mm.Meteors[position+1:]...)
	mm.score++
}

func (mm *MeteorManager) CheckCollision(p *Player) bool {
	hasCollided := false
	for _, meteor := range mm.Meteors {
		if meteor.CollisionDetection(p) {
			hasCollided = true
			break
		}
	}
	return hasCollided
}
