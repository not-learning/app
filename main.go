package main

import (
	"log"

	"github.com/not-learning/app/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := game.Init()
	ebiten.SetWindowSize(g.ScreenW, g.ScreenH)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	//ebiten.SetWindowDecorated(false)
	if e := ebiten.RunGame(g); e != nil {
		log.Fatal(e)
	}
}
