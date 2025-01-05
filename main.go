package main

import (
	"cruiser/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := game.NewGame()

	ebiten.SetWindowTitle("Cruiser")
	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
