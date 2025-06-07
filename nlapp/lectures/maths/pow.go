package maths

import (//"fmt"
	"embed"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/not-learning/app/nlapp/clrs"
	"github.com/not-learning/app/nlapp/frame"
	"github.com/not-learning/app/nlapp/parse"
)

//go:embed pow.txt
var powRaw string

// TODO proper tracks
//go:embed tracks/pow
var powFiles embed.FS

type Pow struct {
	clrs.Clr
	*frame.Frame

	clrY uint8

	input int
	check bool
}

func InitPow(x1, y1, x2, y2, scale float32) *Pow {
	p := &Pow{}
	p.Frame = frame.Init(x1, y1, x2, y2, scale)
	p.Tracks.InitFiles("tracks/pow", powFiles)
	p.Clr = clrs.White

	exs := parse.Do(powRaw)
	for _, ex := range exs {
		p.AddEx(
			p.Subs(ex.Sub()),
			p.Shapes(ex.Shapes()),
			ex.Anim,
			p.Xacts(ex.Xact()),
			ex.Zero,
		)
	}
	return p
}

func (p *Pow) Subs(str string) func(*ebiten.Image) {
	return p.Sub(str)
}

func (p *Pow) Shapes(sh parse.Shape) func(*ebiten.Image) {
	return func(scr *ebiten.Image) {
		p.PlayConShow()
		if sh.IsEx() {
			p.NumPadShow()
		} else {
			p.NumPadHide()
		}
		for _, pl := range sh.Polys() {
			crds := pl.Poly()
			p.PolyEmp(scr, crds, clrs.Blue)
		}

		for _, cr := range sh.Circles() {
			x, y, r := cr.Circle()
			p.CirEmp(scr, x, y, r, clrs.Green)
		}

		for _, lb := range sh.Labels() {
			str, size, x, y := lb.Label()
			if lb.IsEx() { str += strconv.Itoa(p.input) }
			p.Label(scr, str, size, x, y, clrs.White)
		}
	}
}

func (p *Pow) Xacts(x parse.Xacts) func() bool {
	return func () bool {
		if x.IsEx() {
			p.check = false
		} else {
			p.check = true
			return p.check
		}
		p.input = p.Input()
		p.Erase()
		correct, ok := p.Check(int(x.AnsF()))
		if correct && ok {
			p.check = true
			p.Next()
			p.Play()
		}
		return p.check
	}
}
