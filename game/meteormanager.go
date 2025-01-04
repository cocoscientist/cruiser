package game

import (
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type MeteorManager struct {
	Meteors              []*Meteor
	meteorSpawnTimer     *Timer
	baseSpeed            float64
	accelerationConstant float64
	score                int
}

func NewMeteors(baseSpeed float64, accelerationConstant float64) *MeteorManager {
	mm := &MeteorManager{}
	mm.meteorSpawnTimer = NewTimer(3000 * time.Millisecond)
	mm.baseSpeed = baseSpeed
	mm.accelerationConstant = accelerationConstant
	mm.score = 0
	return mm
}

func (mm *MeteorManager) UpdateAllMeteors() {
	mm.meteorSpawnTimer.Update()
	if mm.meteorSpawnTimer.IsReady() {
		mm.addMeteor()
		mm.meteorSpawnTimer.Reset()
	}
	for i, meteor := range mm.Meteors {
		meteor.Update(mm.baseSpeed)
		if meteor.X < -1*meteor.getWidth()/2 {
			mm.removeMeteor(i)
		}
	}
}

func (mm *MeteorManager) UpdateSpeed(meteor *Meteor, p *Player) {
	mm.baseSpeed += ((mm.accelerationConstant) / (meteor.GetDistance(&p.Entity))) * ((meteor.X - p.X) / math.Sqrt(meteor.GetDistance(&p.Entity)))
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
	mm.Meteors = append(mm.Meteors, NewMeteor())
}

func (mm *MeteorManager) removeMeteor(position int) {
	mm.Meteors = append(mm.Meteors[:position], mm.Meteors[position+1:]...)
	mm.score++
	if mm.score%5 == 0 {
		mm.meteorSpawnTimer.AdjustTicker(0.95 * mm.meteorSpawnTimer.targetTicks)
	}
}
