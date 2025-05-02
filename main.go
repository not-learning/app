package main

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"

	lobby "github.com/not-learning/app/lectures"
)

// TODO image.Bounds() and friends TODO
// TODO relative sizes
// TODO: have lists prepared for drawing

// const scrW, scrH = 900, 2000
// const scrW, scrH = 270, 600
const scrW, scrH = 450, 1000

type game struct{ *lobby.Lobby }

func (g *game) Update() error {
	g.Lobby.Update()
	return nil
}

func (g *game) Draw(scr *ebiten.Image) { g.Lobby.Draw(scr) }

func (g *game) Layout(outW, outH int) (int, int) {
	scale := ebiten.DeviceScaleFactor()
	canvasW := int(math.Ceil(float64(outW) * scale))
	canvasH := int(math.Ceil(float64(outH) * scale))
	return canvasW, canvasH
}

func initGame() *game {
	g := &game{}
	g.Lobby = lobby.Init(0, 0, scrW, scrH)
	return g
}

func main() {
	g := initGame()
	ebiten.SetWindowSize(scrW, scrH)
	//ebiten.SetWindowDecorated(false)
	if e := ebiten.RunGame(g); e != nil {
		log.Fatal(e)
	}
}
