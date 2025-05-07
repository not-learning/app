package frame

import (
	"github.com/not-learning/app/clrs"
	"github.com/not-learning/app/fonts"
	"github.com/not-learning/app/graph"
	"github.com/not-learning/app/inter"
	"github.com/not-learning/app/tracks"

	"github.com/hajimehoshi/ebiten/v2"
)

type ex struct {
	sub   func(*ebiten.Image)
	shape func(*ebiten.Image)
	anim  func() bool
	xact  func() bool
	zero  func()
}

// TODO: think if move x0, y0 here
type Lect struct {
	play bool
	exs  []ex
	n    int

	np *numpad

	b *blocks
	*fonts.Font
	*graph.Graph
	*tracks.Tracks
}

type numpad struct {
	str []string
	num []int

	input int
}

func Init(x1, y1, x2, y2 float32) *Lect {
	l := &Lect{}
	l.np = &numpad{}
	l.b = initBlocks(x1, y1, x2, y2)
	l.Font = fonts.InitFont()
	l.Graph = graph.Init()
	l.Graph.SetOrigin(l.b.screen.MidF32())
	l.Tracks = tracks.Init()

	l.np.str = []string{
		"7", "8", "9",
		"4", "5", "6",
		"1", "2", "3",
		"⌫", "0", "✓",
	}

	l.np.num = []int{
		7, 8, 9,
		4, 5, 6,
		1, 2, 3,
		0,
	}

	// l.play = true
	return l
}

func (l *Lect) Play() {
	l.play = true
	l.Tracks.Play()
}

func (l *Lect) Pause() {
	l.play = !l.play
	l.Tracks.Pause()
}

func (l *Lect) Proceed() {
	var a, x bool
	if l.play { a = l.exs[l.n].anim() }
	x = l.exs[l.n].xact()

	if pl, ok := l.Tracks.IsPlaying(); x && a && l.play && !pl && ok {
		if l.n < len(l.exs)-1 {
			l.exs[l.n].zero()
			l.n++
		}
		l.Tracks.Proceed()
	}
}

func (l *Lect) Next() {
	l.play = false
	l.exs[l.n].zero()
	if l.n < len(l.exs)-1 { l.n++ }
	l.exs[l.n].zero()
	l.Tracks.Switch(l.n)
}

func (l *Lect) Prev() {
	l.play = false
	l.exs[l.n].zero()
	if l.n > 0 { l.n-- }
	l.exs[l.n].zero()
	l.Tracks.Switch(l.n)
}

func (l *Lect) SubWrap(sub string) []string {
	res := []string{}
	x1, _, x2, _ := l.b.subs[0].Rect()
	w := float64(x2-x1) - 20 // TODO proper sizes
	l.Font.Set(15, clrs.White)
	res = append(res, l.Font.Wrap(sub, w)...)
	return res
}

func (l *Lect) SubDraw(sub []string) func(*ebiten.Image) {
	return func(scr *ebiten.Image) {
		for i, v := range sub {
			x, y, _, _ := l.b.subs[i].Rect()
			l.Font.Set(15, clrs.White)
			l.Font.Draw(scr, v, x, y)
		}
	}
}

func (l *Lect) AddEx(sub, shape func(*ebiten.Image), anim, xact func() bool, zero func()) {
	x := ex{sub: sub, shape: shape, anim: anim, xact: xact, zero: zero}
	l.exs = append(l.exs, x)
}

func (l *Lect) NumPadShow() { l.b.numPadShow() }

func (l *Lect) NumPadHide() { l.b.numPadHide() }

/*func (l *Lect) Touch() {}//*/

func (l *Lect) Input() int {
	if n, ok := inter.Number(); ok {
		l.np.input = n + l.np.input*10
	}

	for i, v := range l.b.npl {
		if i == 9 || i == 11 {
			continue
		}
		if inter.MouseInL(v.Rect()) {
			l.np.input = l.np.input*10 + l.np.num[i]
		}
	}
	return l.np.input
}

func (l *Lect) Erase() {
	if inter.MouseInL(l.b.npl[9].Rect()) || inter.Backspace() {
		l.np.input /= 10
	}
}

// todo: not very good
func (l *Lect) Check(solution int) (correct, ok bool) {
	correct = l.np.input == solution
	ok = inter.MouseInL(l.b.npl[11].Rect()) || inter.Enter()
	if ok { l.PlayCorrect(correct) }
	return
}

/*func (l *Lect) Check(solution int) bool {
	if inter.MouseInL(l.b.npl[11].Rect()) || inter.Enter() {
		if l.np.input == solution {
			l.PlayCorrect()
			return true
		}
		l.PlayWrong()
	}
	return false
}//*/

func (l *Lect) Update() {
	inter.Escape()
	if inter.Space() || inter.MouseInL(l.b.pause.Rect()) { l.Pause() }

	if inter.ArrowLeft() || inter.MouseInL(l.b.prev.Rect()) {
		l.Prev()
		return
	}
	if inter.ArrowRight() || inter.MouseInL(l.b.next.Rect()) {
		l.Next()
		return
	}

	l.Proceed()
}

func (l *Lect) Draw(screen *ebiten.Image) {
	l.exs[l.n].shape(screen)
	l.exs[l.n].sub(screen)
	//graph.Draw(screen, l.b.prev, l.b.pause, l.b.next)
	//l.b.top.WalkDown(graph.TestDraw(screen))

	l.Font.Set(75, clrs.White)
	x, y := l.b.prev.MidF32()
	l.Font.DrawCenter(screen, "←", x, y)
	l.Font.Set(75, clrs.White)
	x, y = l.b.next.MidF32()
	l.Font.DrawCenter(screen, "→", x, y)

	l.Font.Set(50, clrs.White)
	x, y = l.b.pause.MidF32()
	if l.play {
		l.Font.DrawCenter(screen, "▮▮", x, y)
	} else {
		l.Font.DrawCenter(screen, "▶", x, y)
	}

	if l.b.npshow {
		//graph.Draw(screen, l.b.npl...)
		for i, v := range l.b.npl {
			x, y = v.MidF32()
			l.Font.Set(50, clrs.White)
			l.Font.DrawCenter(screen, l.np.str[i], x, y)
		}
	}
}

// ←→
//⬅➡⮕ ⇦ ⇨ 🡐 🡒 ⇽ ⇾ ⟵ ⟶ ⟸ ⟹ ⟻ ⟼ ⇜ ⇝ ⬿ ⤳ ↫ ↬ ⮜ ⮞ ⮪ ⮫ ◄ ► ◅ ▻
