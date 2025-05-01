package maths

import (
	"embed"
	//"math"
	"strconv"

	"github.com/not-learning/app/clrs"
	"github.com/not-learning/app/frame"

	"github.com/hajimehoshi/ebiten/v2"
)

type Pow struct {
	*frame.Lect
	clrs.Clr

	ans int
	input int
}

// ## Ex1
func (p *Pow) sub1() func(*ebiten.Image) {
	return p.SubDraw(p.SubWrap(
		`Ты конечно помнишь, что умножение - это сложение одинаковых чисел.`,
	))
}

func (p *Pow) shape1(scr *ebiten.Image) {
	p.Label(scr, "Степень", 30, 0, 0, p.Clr)
}

func (p *Pow) anim1() bool { return true }

func (p *Pow) xact1() bool {
	frame.Escape() //TODO: move basic xacts to frame
	if frame.Space() { p.Pause() }
	return true
}

// ## Ex2
func (p *Pow) sub2() func(*ebiten.Image) {
	return p.SubDraw(p.SubWrap(
		`Чему равно, скажем, пять умножить на три?`,
	))
}

func (p *Pow) shape2(scr *ebiten.Image) {
	s := "?"
	if p.input > 0 { s = strconv.Itoa(p.input) }
	p.Label(scr, "5·3 = " + s, 30, 0, 0, p.Clr)
}

func (p *Pow) anim2() bool { return true }

func (p *Pow) xact2cl() func() bool {
	ret := false

	return func() bool {
		_ = p.xact1()
		p.ans = 15
		if i, ok := frame.Number(); ok {
			p.input = i + 10*p.input
		}
		if frame.Enter() {
			if p.input != p.ans {
				p.PlayWrong()
				return ret
			}
			p.PlayCorrect()
			ret = true
		}
		return ret
	}
}

// ## Ex3
func (p *Pow) sub3() func(*ebiten.Image) {
	return p.SubDraw(p.SubWrap(
		`Так вот, степень - это умножение одинаковых чисел.`,
	))
}

func (p *Pow) shape3(scr *ebiten.Image) { p.shape1(scr) }
func (p *Pow) anim3() bool { return true }
func (p *Pow) xact3() bool {
	_ = p.xact1()
	return true
}

// ## Ex4
func (p *Pow) sub4() func(*ebiten.Image) {
	return p.SubDraw(p.SubWrap(
		`Пять в третьей это пять на пять на пять.`,
	))
}

func (p *Pow) shape4(scr *ebiten.Image) {
	p.Label(scr, "5³ = 5·5·5", 30, 0, 0, p.Clr)
}

func (p *Pow) anim4() bool { return true }

func (p *Pow) xact4() bool {
	_ = p.xact1()
	return true
}
//*/


// TODO proper tracks
//go:embed tracks/pow
var powF embed.FS

func InitPow(x1, y1, x2, y2 float32) *Pow {
	p := &Pow{}
	p.Lect = frame.Init(x1, y1, x2, y2)
	p.Lect.Tracks.InitFiles("tracks/pow", powF)
	p.Clr = clrs.White

	p.AddEx(p.sub1(), p.shape1, p.anim1, p.xact1)
	p.AddEx(p.sub2(), p.shape2, p.anim2, p.xact2cl())
	p.AddEx(p.sub3(), p.shape3, p.anim3, p.xact3)
	p.AddEx(p.sub4(), p.shape4, p.anim4, p.xact4)

	return p
}

/*func (p *Pow) shape1(scr *ebiten.Image) {
	p.Coords(scr, clrs.Green)

	p.Arc(scr, 0, 0, p.r, 0, p.a1, p.Clr)
	//p.Arrow(scr, 0, 0, 1.2*p.x, 1.2*p.y, clrs.White)
	p.PolyEmp(scr, []float32{0, 0, p.x, p.y}, clrs.White)

	//p.PolyEmp(scr, []float32{0, 0, x*1.5, y*1.5}, clrs.White)
	//p.PolyEmp(scr, []float32{0, 0, r*1.5, 0}, clrs.White)

	p.PolyEmp(scr, []float32{p.x, p.y, p.x, 0}, clrs.Blue)
	p.PolyEmp(scr, []float32{p.x, p.y, 0, p.y}, clrs.Blue)

	p.CirFull(scr, p.x, p.y, 4, clrs.White)
	p.CirFull(scr, p.x,   0, 4, clrs.White)
	p.CirFull(scr, 0,   p.y, 4, clrs.White)

	p.Label(scr, "x", 15, p.x, -10, clrs.White)
	p.Label(scr, "y", 15, 10, p.y, clrs.White)
}//*/

/*func (p *Pow) anim1() {
		p.a2 += 0.02
		p.xy()
	if p.a2 >= 4*math.Pi { p.a2 = 2*math.Pi }
	//if p.a2 >= 2.33*math.Pi { p.a2 = 2.33*math.Pi } else { p.a2 += 0.02 }
}//*/
