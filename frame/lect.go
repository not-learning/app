package frame

import (
	"github.com/not-learning/app/clrs"
	"github.com/not-learning/app/fonts"
	"github.com/not-learning/app/graph"
	"github.com/not-learning/app/tracks"

	"github.com/hajimehoshi/ebiten/v2"
)

type ex struct {
	sub   func(*ebiten.Image)
	shape func(*ebiten.Image)
	anim  func() bool
	xact  func() bool
}

// TODO: think if move x0, y0 here
type Lect struct {
	play bool
	//Done bool
	exs  []ex
	n    int

	b *blocks
	*fonts.Font
	*graph.Graph
	*tracks.Tracks
}

func Init(x1, y1, x2, y2 float32) *Lect {
	l := &Lect{}
	l.Font = fonts.InitFont()
	l.b = initBlocks(x1, y1, x2, y2)
	l.Graph = graph.Init()
	l.Graph.SetOrigin(l.b.screen.MidF32())
	l.Tracks = tracks.Init()
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

func (l *Lect) next() {
	if l.n < len(l.exs)-1 { l.n++ }
	l.Tracks.Next()
}

func (l *Lect) SubWrap(sub string) []string {
	res := []string{}
	x1, _, x2, _ := l.b.subs[0].Rect()
	w := float64(x2 - x1) - 20 // TODO proper sizes
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

func (l *Lect) AddEx(sub, shape func(*ebiten.Image), anim, xact func() bool) {
	x := ex{sub: sub, shape: shape, anim: anim, xact: xact}
	l.exs = append(l.exs, x)
}

func (l *Lect) NumPadShow() { l.b.numPadShow() }

func (l *Lect) NumPadHide() {
	l.b.numPadHide()
}

/*func (l *Lect) Touch() {}//*/

func (l *Lect) Draw(screen *ebiten.Image) {
	l.exs[l.n].shape(screen)
	l.exs[l.n].sub(screen)
	graph.Draw(screen, l.b.npl...)
//l.b.top.WalkDown(graph.TestDraw(screen))

	x, y := l.b.pause.MidF32()
	l.Font.Set(50, clrs.White)
	if l.play {
		l.Font.DrawCenter(screen, "▮▮", x, y)
	} else {
		l.Font.DrawCenter(screen, "▶", x, y)
	}
}

func (l *Lect) Update() {
	var a, x bool
	if l.play { a = l.exs[l.n].anim() }
	x = l.exs[l.n].xact()
	if x && a && !l.Tracks.IsPlaying() { l.next() }
}
