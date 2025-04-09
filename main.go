package main

import (
	"log"
	"math"
	"github.com/not-learning/app/lectures/maths"
	//f "github.com/not-learning/app/fonts"

	"github.com/hajimehoshi/ebiten/v2"
)

// TODO relative sizes
// TODO: have lists prepared for drawing

//const scrW, scrH = 900, 2000
const scrW, scrH = 450, 1000

type game struct {
	pow *maths.Power
	//Font *f.Font
}

func (g *game) Update() error {
	g.pow.Update()
	return nil
}

func (g *game) Draw(scr *ebiten.Image) {
	g.pow.Draw(scr)
}

func (g *game) Layout(outW, outH int) (int, int) {
	scale := ebiten.DeviceScaleFactor()
	canvasW := int(math.Ceil(float64(outW) * scale))
	canvasH := int(math.Ceil(float64(outH) * scale))
	return canvasW, canvasH
}

func initGame() *game {
	g := &game{}
	g.pow = maths.InitPow(0, 0, scrW, scrH)
	//g.Font = f.InitFont()
	return g
}

func main() {
	g := initGame()
	ebiten.SetWindowSize(scrW, scrH)
	ebiten.SetWindowDecorated(false)
	if e := ebiten.RunGame(g); e != nil {
		log.Fatal(e)
	}
}
