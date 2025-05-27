package game

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/not-learning/app/nlapp/home"
)

// TODO package tools <- clrs, fonts, inter...
// TODO Check should work even when !play TODO
// TODO image.Bounds() and friends TODO
// TODO relative sizes
// TODO: have lists prepared for drawing

//const scrW, scrH = 1000, 2000
//const scrW, scrH = 300, 600
const scrW, scrH = 500, 1000

type game struct{
	*home.Home
	ScreenW, ScreenH int
}

func (g *game) Update() error {
	g.Home.Update(g.ScreenW, g.ScreenH)
	return nil
}

func (g *game) Draw(scr *ebiten.Image) {
	g.Home.Draw(scr)
}

func (g *game) Layout(outW, outH int) (int, int) {
	scale := ebiten.Monitor().DeviceScaleFactor()
	g.ScreenW = int(float64(outW) * scale)
	g.ScreenH = int(float64(outH) * scale)
	return g.ScreenW, g.ScreenH
}

func Init() *game {
	g := &game{}
	g.ScreenW, g.ScreenH = scrW, scrH
	g.Home = home.Init(0, 0, scrW, scrH)
	return g
}
