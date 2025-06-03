package maths

import ("fmt"
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
	/*subs := make([]string, len(pp))
	for i, p := range pp {
		subs[i] = p.Sub()
	}*/

	for _, p := range pp {
		//ll := p.Labels()
		a.AddEx(a.doSub(p.Sub()), a.shape1, a.anim1, a.xact1, a.zero1)
	}

ll := pp[13].Labels()
fmt.Printf("%T\n", ll[0])
str, x, y, size := ll[0].Label()
fmt.Println("'"+str+"'", x+1, y+1, size-12)

	//a.AddEx(a.sub1(), a.shape1, a.anim1, a.xact1, a.zero1)

	return a
}

func (a *Alg) doSub(str string) func(*ebiten.Image) {
	return a.Sub(str)
}

/*func (a *Alg) doShape(ll []parser.Label) func(*ebiten.Image) {
	return func(scr *ebiten.Image) {
		a.PlayConShow()
		for _, lb := range ll {
			str, x, y, size := lb.Label()
			a.Label(scr, str, x, y, size, clrs.White)
		}
	}
}*/

func (a *Alg) sub1() func(*ebiten.Image) {
	return a.Sub(
		`Ты конечно помнишь, что умножение - это сложение одинаковых чисел.`,
	)
}

func (a *Alg) shape1(scr *ebiten.Image) {
	a.PlayConShow()
	a.Label(scr, "Степень", 30, 0, 0, clrs.White)
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

