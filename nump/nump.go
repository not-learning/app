package nump

import (
	c "github.com/not-learning/app/clrs"
	f "github.com/not-learning/app/fonts"
	"strconv"

	l "github.com/not-learning/lmnts"
	"github.com/hajimehoshi/ebiten/v2"
)

type NumP struct {
	screen   *l.Lmnt
	npBtns   []*l.Lmnt
	eraseBtn *l.Lmnt
	checkBtn *l.Lmnt

	xact *xAct
	font *f.Font
}

func Init(x1, y1, x2, y2 float32) *NumP {
	top := l.New()
	top.SetRect(x1, y1, x2, y2)

	n := &NumP{}
	n.screen = l.New()
	n.screen.SetSize(0, y2*2/3)
	top.Add(n.screen)

	npRows := []*l.Lmnt{}
	for i := range 4 {
		npRows = append(npRows, l.New())
		npRows[i].SetRow()
	}

	n.npBtns = make([]*l.Lmnt, 10)
	for i := range n.npBtns {
		n.npBtns[i] = l.New()
		n.npBtns[i].SetSize(100, 75)
	}

	n.eraseBtn = l.New()
	n.eraseBtn.SetSize(100, 75)
	n.checkBtn = l.New()
	n.checkBtn.SetSize(100, 75)

	{
		var gg float32 = 0
		npRows[0].GapsAround(gg, n.npBtns[7:]...)
		npRows[1].GapsAround(gg, n.npBtns[4:7]...)
		npRows[2].GapsAround(gg, n.npBtns[1:4]...)
		npRows[3].GapsAround(gg, n.eraseBtn, n.npBtns[0], n.checkBtn)

		margin := l.New()
		margin.GapsAround(gg, npRows...)
		top.Add(margin)
	}

	top.DoAll()

	n.xact = initXact()
	for _, v := range n.npBtns { // todo func
		x1, y1, x2, y2 = v.Rect()
		n.xact.btnPos = append(n.xact.btnPos, &pos{
			float64(x1), float64(y1), float64(x2), float64(y2),
		})
	}

	x1, y1, x2, y2 = n.eraseBtn.Rect()
	n.xact.btnErase = &pos{float64(x1), float64(y1), float64(x2), float64(y2)}
	x1, y1, x2, y2 = n.checkBtn.Rect()
	n.xact.btnCheck = &pos{float64(x1), float64(y1), float64(x2), float64(y2)}

	n.font = f.InitFont()

	return n
}

func (n *NumP) Update() {
	n.xact.update()
}

func (n *NumP) Draw(scr *ebiten.Image) {
	//n.top.WalkUp(myDraw(scr))
	l.Draw(scr, n.npBtns...)
	l.Draw(scr, n.eraseBtn, n.checkBtn)

	var x, y float32
	for i, k := range n.npBtns {
		x, y = k.MidF32()
		n.font.Set(50, c.White)
		n.font.DrawCenter(scr, strconv.Itoa(i), x, y)
	}

	x, y = n.eraseBtn.MidF32()
	n.font.Set(33, c.White)
	n.font.DrawCenter(scr, "⌫", x, y)

	x, y = n.checkBtn.MidF32()
	n.font.Set(75, c.Green)
	n.font.DrawCenter(scr, "✓", x, y)
}

func (n *NumP) Num() int { return *n.xact.int }
func (n *NumP) Str() string { return strconv.Itoa(*n.xact.int) }
func (n *NumP) NumClear() { *n.xact.int = 0 }

func (n *NumP) Check(ans int) bool {
	if *n.xact.bool { return ans == *n.xact.int }
	return false
}

func (n *NumP) IsPlaying() bool { return *n.xact.play }

func (n *NumP) Play()  { *n.xact.play = true }
func (n *NumP) Stop()  { *n.xact.play = false }
func (n *NumP) Pause() { *n.xact.play = !*n.xact.play }
