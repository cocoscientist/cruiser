package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// x and y are co-ordinates of the center

type Entity struct {
	X      float64
	Y      float64
	width  float64
	height float64
	sprite *ebiten.Image
}

func (e *Entity) getWidth() float64 {
	return float64(e.sprite.Bounds().Dx())
}

func (e *Entity) getHeight() float64 {
	return float64(e.sprite.Bounds().Dy())
}

func (e *Entity) Draw(screen *ebiten.Image) {
	bounds := e.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Translate(e.X, e.Y)
	screen.DrawImage(e.sprite, op)
}

func (e *Entity) GetDistance(e2 *Entity) float64 {
	return math.Sqrt((e.X-e2.X)*(e.X-e2.X) + (e.Y-e2.Y)*(e.Y-e2.Y))
}
