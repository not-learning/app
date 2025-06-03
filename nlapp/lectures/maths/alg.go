package maths

import (//"fmt"
	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/not-learning/app/nlapp/clrs"
	"github.com/not-learning/app/nlapp/frame"
	"github.com/not-learning/app/nlapp/parser"
)

//go:embed alg.txt
var algRaw string

type Alg struct {
	*frame.Frame
	clrs.Clr

	input int
	check bool
}

func InitAlg(x1, y1, x2, y2, scale float32) *Alg {
	a := &Alg{}
	a.Frame = frame.Init(x1, y1, x2, y2, scale)
	a.Tracks.InitFiles("tracks/pow", powFiles) // TODO
	a.Clr = clrs.White

	pp := parser.Do(algRaw)
	for _, p := range pp {
		// ll := p.Labels()
		a.AddEx(a.doSub(p.Sub()), a.doShape(p.Labels()), a.anim1, a.xact1, a.zero1)
	}

	return a
}

func (a *Alg) doSub(str string) func(*ebiten.Image) {
	return a.Sub(str)
}

func (a *Alg) doShape(ll []parser.Label) func(*ebiten.Image) {
	return func(scr *ebiten.Image) {
		a.PlayConShow()
		for _, lb := range ll {
			str, size, x, y := lb.Label()
			a.Label(scr, str, size, x, y, clrs.White)
		}
	}
}

func (a *Alg) anim1() bool {
	// s := uint8(7)
	// if a.clrY < 255-s {
	// 	a.clrY += s
	// 	return false
	// }
	return true
}

func (a *Alg) xact1() bool {
	// ans := 15
	// p.input = p.Input()
	// p.Erase()
	// correct, ok := p.Check(ans)
	// if correct && ok { p.check = true }
	// if p.check { p.NumPadHide() }
	// return p.check
	return true
}
func (a *Alg) zero1() {}

