package main

import (
	"log"

	"github.com/not-learning/app/home"

	"github.com/hajimehoshi/ebiten/v2"
)

// TODO package tools <- clrs, fonts, inter...
// TODO Check should work even when !play TODO
// TODO image.Bounds() and friends TODO
// TODO relative sizes
// TODO: have lists prepared for drawing

// const scrW, scrH = 900, 2000
//const scrW, scrH = 270, 600
const scrW, scrH = 450, 1000

type game struct{
	*home.Home
	screenW, screenH int
	ratW, ratH float32
}

func (g *game) Update() error {
	g.ratW, g.ratH = float32(g.screenW)/scrW, float32(g.screenH)/scrH
	g.Home.Update(g.screenW, g.screenH, g.ratW, g.ratH)
	return nil
}

func (g *game) Draw(scr *ebiten.Image) {
	g.Home.Draw(scr)
}

func (g *game) Layout(outW, outH int) (int, int) {
	scale := ebiten.Monitor().DeviceScaleFactor()
	g.screenW = int(float64(outW) * scale)
	g.screenH = int(float64(outH) * scale)
	return g.screenW, g.screenH
}

func initGame() *game {
	g := &game{}
	g.Home = home.Init(0, 0, scrW, scrH)
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
