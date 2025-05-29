package game

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/not-learning/app/nlapp/home"
)

// TODO update sizes properly TODO
// TODO Check should work even when !play TODO
// TODO image.Bounds() and friends
// TODO relative sizes
// TODO: have lists prepared for drawing

//const scrW, scrH = 1000, 2000
//const scrW, scrH = 300, 600
const scrW, scrH = 500, 1000

type game struct{
	*home.Home
	scale float32
	ScreenW, ScreenH int
}

func Init() *game {
	g := &game{}
	g.ScreenW, g.ScreenH = scrW, scrH
	g.Home = home.Init(0, 0, scrW, scrH, 1)
	return g
}

func (g *game) Layout(outW, outH int) (int, int) {
	g.scale = float32(ebiten.Monitor().DeviceScaleFactor())
	g.ScreenW = int(float32(outW) * g.scale)
	g.ScreenH = int(float32(outH) * g.scale)
	return g.ScreenW, g.ScreenH
}

func (g *game) Update() error {
	g.Home.Update(g.ScreenW, g.ScreenH, g.scale)
	return nil
}

func (g *game) Draw(scr *ebiten.Image) {
	g.Home.Draw(scr)
}
