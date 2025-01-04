package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func GetFrameValue(s float64) float64 {
	return s / float64(ebiten.TPS())
}
