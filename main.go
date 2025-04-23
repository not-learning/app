package main

import (
	"log"
	"math"
	"image/color"
	"github.com/not-learning/app/lectures/maths"

	"github.com/hajimehoshi/ebiten/v2"
)

// TODO image.Bounds() and friends TODO
// TODO relative sizes
// TODO: have lists prepared for drawing

// const scrW, scrH = 900, 2000
const scrW, scrH = 450, 1000
//const scrW, scrH = 270, 600

type game struct {
	//pow *maths.Power
	trig *maths.Trig
}

func (g *game) Update() error {
	//g.pow.Update()
	g.trig.Update()
	return nil
}

func (g *game) Draw(scr *ebiten.Image) {
	s := scr
	s.Fill(color.RGBA{20, 20, 20, 255})
	//g.pow.Draw(scr)
	g.trig.Draw(scr)
}

func (g *game) Layout(outW, outH int) (int, int) {
	scale := ebiten.DeviceScaleFactor()
	canvasW := int(math.Ceil(float64(outW) * scale))
	canvasH := int(math.Ceil(float64(outH) * scale))
	return canvasW, canvasH
}

func initGame() *game {
	g := &game{}
	//g.pow = maths.InitPow(0, 0, scrW, scrH)
	g.trig = maths.InitTrig(0, 0, scrW, scrH)
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
