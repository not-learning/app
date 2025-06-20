package maths

import (//"fmt"
	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/not-learning/app/nlapp/clrs"
	"github.com/not-learning/app/nlapp/frame"
	"github.com/not-learning/app/nlapp/parse"
)

//go:embed trig.txt
var trigRaw string

type Trig struct {
	clrs.Clr
	*frame.Frame

	input int
	check bool
}

func InitTrig(x1, y1, x2, y2, scale float32) *Trig {
	t := &Trig{}
	t.Frame = frame.Init(x1, y1, x2, y2, scale)
	t.Tracks.InitFiles("tracks/pow", powFiles) // TODO
	t.Clr = clrs.White

	exs := parse.Do(trigRaw)
	for _, ex := range exs {
		t.AddEx(
			t.Subs(ex.Sub()),
			t.Shapes(ex.Shapes()),
			//t.Anim(p.Anims()),
			ex.Anim,
			t.xact1,
			//t.zero1,
			ex.Zero,
		)
	}
	return t
}

func (t *Trig) Subs(str string) func(*ebiten.Image) {
	return t.Sub(str)
}

func (t *Trig) Shapes(sh parse.Shape) func(*ebiten.Image) {
	return func(scr *ebiten.Image) {
		t.PlayConShow()

		for _, pl := range sh.Polys() {
			crds := pl.Poly()
			t.PolyEmp(scr, crds, clrs.Blue)
		}

		for _, cr := range sh.Circles() {
			x, y, r := cr.Circle()
			t.CirEmp(scr, x, y, r, clrs.Green)
		}

		for _, lb := range sh.Labels() {
			str, size, x, y := lb.Label()
			t.Label(scr, str, size, x, y, clrs.White)
		}
	}
}

/*func (t *Trig) Anim(an parse.Anims) func() bool {
	return an.Anim
}*/

/*func (t *Trig) Zero(z parse.Zero) func() {
	return z.Zero
}*/

func (t *Trig) anim1() bool {
	// s := uint8(7)
	// if t.clrY < 255-s {
	// 	t.clrY += s
	// 	return false
	// }
	return true
}

func (t *Trig) xact1() bool {
	// ans := 15
	// p.input = p.Input()
	// p.Erase()
	// correct, ok := p.Check(ans)
	// if correct && ok { p.check = true }
	// if p.check { p.NumPadHide() }
	// return p.check
	return true
}

func (t *Trig) zero1() {}
