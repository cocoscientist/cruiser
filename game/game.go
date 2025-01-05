package game

import (
	"cruiser/assets"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	screenWidth          = 800
	screenHeight         = 600
	baseSpeed            = 100
	accelerationConstant = 15
	playerX              = 85
)

type Game struct {
	player        *Player
	bg            *Background
	meteorManager *MeteorManager
	gameOver      bool
	toggleEngine  bool
}

func NewGame() *Game {
	g := &Game{}
	g.bg = NewBackground()
	g.player = NewPlayer(accelerationConstant, playerX, screenHeight/2)
	g.meteorManager = NewMeteors(baseSpeed, accelerationConstant)
	g.gameOver = false
	g.toggleEngine = false
	return g
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.bg.Draw(screen)

	if !g.gameOver {
		g.player.Draw(screen, g.toggleEngine)
		g.meteorManager.DrawAllMeteors(screen)
		text.Draw(screen, fmt.Sprintf("SCORE:%04d", g.meteorManager.score), assets.ScoreFont, 20, 60, color.White)
	} else {
		text.Draw(screen, fmt.Sprintf("FINAL SCORE: %04d", g.meteorManager.score), assets.GameOverFont, screenWidth/2-200, 90, color.White)
		text.Draw(screen, fmt.Sprintf("GAME OVER!"), assets.GameOverFont, screenWidth/2-114, screenHeight/2-100, color.White)
		text.Draw(screen, fmt.Sprintf("Press Space to Play Again!"), assets.GameOverFont, screenWidth/2-302, screenHeight/2, color.White)
	}
}

func (g *Game) Update() error {
	g.bg.Update()
	if g.gameOver {
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.resetGame()
		}
	} else {
		g.player.Update()
		g.meteorManager.UpdateAllMeteors()
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.toggleEngine = !g.toggleEngine
		}
		if g.toggleEngine {
			meteor := g.meteorManager.GetClosestMeteor(playerX)
			g.player.UpdateVerticalVelocity(meteor)
			g.meteorManager.UpdateSpeed(meteor, g.player)
		} else {
			g.player.ResetVerticalVelocity()
		}
	}
	return nil
}

func (g *Game) resetGame() {
	g.gameOver = false
	g.player = NewPlayer(accelerationConstant, playerX, screenHeight/2)
	g.meteorManager = NewMeteors(baseSpeed, accelerationConstant)
	g.toggleEngine = false
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
