package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type MeteorManager struct {
	Meteors          []*Meteor
	meteorSpawnTimer *Timer
	baseSpeed        float64
}

func NewMeteors(baseSpeed float64) *MeteorManager {
	mm := &MeteorManager{}
	mm.meteorSpawnTimer = NewTimer(2500 * time.Millisecond)
	mm.baseSpeed = baseSpeed
	return mm
}

func (mm *MeteorManager) UpdateAllMeteors() {
	mm.meteorSpawnTimer.Update()
	if mm.meteorSpawnTimer.IsReady() {
		mm.addMeteor()
		mm.meteorSpawnTimer.Reset()
	}
	for i, meteor := range mm.Meteors {
		meteor.Update(baseSpeed)
		if meteor.X < -1*meteor.getWidth()/2 {
			mm.removeMeteor(i)
		}
	}
}

func (mm *MeteorManager) UpdateSpeed(p *Player) {

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
}
