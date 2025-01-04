package game

import (
	"cruiser/assets"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

type Game struct {
	player   *Player
	bg       *Background
	ax       float64
	ay       float64
	score    int
	gameOver bool
}

func NewGame() *Game {
	g := &Game{}
	g.bg = NewBackground()
	g.ax = 0
	g.ay = 0
	g.gameOver = false
	return g
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.bg.Draw(screen)
	if !g.gameOver {
		text.Draw(screen, fmt.Sprintf("SCORE:%04d", g.score), assets.ScoreFont, 20, 60, color.White)
	} else {
		text.Draw(screen, fmt.Sprintf("FINAL SCORE: %04d", g.score), assets.GameOverFont, screenWidth/2-200, 90, color.White)
		text.Draw(screen, fmt.Sprintf("GAME OVER!"), assets.GameOverFont, screenWidth/2-114, screenHeight/2-100, color.White)
		text.Draw(screen, fmt.Sprintf("Press Space to Play Again!"), assets.GameOverFont, screenWidth/2-302, screenHeight/2, color.White)
	}
}

func (g *Game) Update() error {
	g.bg.Update()
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
