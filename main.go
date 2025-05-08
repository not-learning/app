package main

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"

	lobby "github.com/not-learning/app/lectures"
)

// TODO Check should work even when !play TODO
// TODO image.Bounds() and friends TODO
// TODO relative sizes
// TODO: have lists prepared for drawing

// const scrW, scrH = 900, 2000
// const scrW, scrH = 270, 600
const scrW, scrH = 450, 1000

type game struct{
	*lobby.Lobby
	screenW, screenH int
}

func (g *game) Update() error {
	g.Lobby.Update(g.screenW, g.screenH)
	return nil
}

func (g *game) Draw(scr *ebiten.Image) {
	g.Lobby.Draw(scr)
}

func (g *game) Layout(outW, outH int) (int, int) {
	scale := ebiten.Monitor().DeviceScaleFactor()
	screenW := int(math.Ceil(float64(outW) * scale))
	screenH := int(math.Ceil(float64(outH) * scale))
	g.screenW, g.screenH = screenW, screenH
	return screenW, screenH
}

func initGame() *game {
	g := &game{}
	g.Lobby = lobby.Init(0, 0, scrW, scrH)
	return g
}

func main() {
	g := initGame()
	ebiten.SetWindowSize(scrW, scrH)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	//ebiten.SetWindowDecorated(false)
	if e := ebiten.RunGame(g); e != nil {
		log.Fatal(e)
	}
}
