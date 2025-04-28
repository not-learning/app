package frame

import (
	//"slices"
	"strings"

	"github.com/not-learning/app/clrs"
	"github.com/not-learning/app/fonts"
	"github.com/not-learning/app/graph"
	"github.com/not-learning/app/tracks"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type lects struct {
	Tracks *tracks.Tracks
	subs   [][]string
	n      int
}

type ex struct {
	shapes []func(*ebiten.Image)
	anims  []func()
	xacts  []func()
	n      int
}

// TODO: think if move x0, y0 here
type Lect struct {
	play bool
	lects
	ex

	*graph.Graph
	*fonts.Font
	b *blocks
}

func Init(x1, y1, x2, y2 float32) *Lect {
	l := &Lect{}
	l.Font = fonts.InitFont()
	l.b = InitBlocks(x1, y1, x2, y2)
	l.Graph = graph.Init()
	l.Graph.SetOrigin(l.b.screen.MidF32())
	l.lects.Tracks = tracks.Init()
	// l.play = true
	return l
} //*/

func (l *Lect) doSub(sub string) {
	res := [][]string{}
	x1, _, x2, _ := l.b.subs[0].Rect()
	w := float64(x2 - x1)
	for i, v := range strings.Split(sub, "\n") {
		l.Font.Set(15, clrs.White)
		res = append(res, []string{})
		for _, k := range l.Font.Wrap(v, w) {
			res[i] = append(res[i], k...)
		}
	}
	l.subs = append(l.subs, res...)
}

func (l *Lect) AddSubs(subs ...string) {
	for i := range subs {
		l.doSub(subs[i])
	}
}

func (l *Lect) AddShapes(fns ...func(*ebiten.Image)) {
	l.shapes = append(l.shapes, fns...)
}

func (l *Lect) AddAnims(fns ...func()) {
	l.anims = append(l.anims, fns...)
}

func (l *Lect) AddXacts(fns ...func()) {
	l.xacts = append(l.xacts, fns...)
}

func (l *Lect) Play() {
	l.play = true
	l.Tracks.Play()
}

func (l *Lect) Pause() {
	l.play = !l.play
	l.Tracks.Pause()
}

func (l *Lect) Next() {
	if l.lects.n <= len(l.lects.subs) { l.lects.n++ }
	if l.ex.n <= len(l.ex.anims) { l.ex.n++ }
	l.Tracks.Next()
}

func (l *Lect) Space() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeySpace)
}

func (l *Lect) MouseL() bool {
	return ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
}

func (l *Lect) MouseR() bool {
	return ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight)
}

/*func (l *Lect) Touch() {

}//*/

func (l *Lect) Draw(screen *ebiten.Image) {
	l.shapes[l.ex.n](screen)

	for i, v := range l.subs[l.lects.n] {
		x, y, _, _ := l.b.subs[i].Rect()
		l.Font.Set(15, clrs.White)
		l.Font.Draw(screen, v, x, y)
	}

	x, y := l.b.pause.MidF32()
	l.Font.Set(50, clrs.White)
	l.Font.DrawCenter(screen, "▶", x, y)

	//l.b.top.WalkUp(graph.TestDraw(screen))
}

func (l *Lect) Update() {
	if l.play { l.anims[l.ex.n]() }
	l.xacts[l.ex.n]()
}
