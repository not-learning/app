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

type ex struct {
	sub   []string
	shape func(*ebiten.Image)
	anim  func()
	xact  func()
}

// TODO: think if move x0, y0 here
type Lect struct {
	play bool
	Done bool
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
	l.b = InitBlocks(x1, y1, x2, y2)
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

func (l *Lect) Next() {
	if l.n < len(l.exs)-1 { l.n++ }
	//l.Tracks.Next() //TODO uncomment (developing)
}

func (l *Lect) doSub(sub string) []string {
	res := []string{}
	x1, _, x2, _ := l.b.subs[0].Rect()
	w := float64(x2 - x1) - 20 // TODO proper sizes
	for _, v := range strings.Split(sub, "\n") {
		l.Font.Set(15, clrs.White)
		//res = append(res, []string{})
		for _, k := range l.Font.Wrap(v, w) {
			res = append(res, k...)
		}
	}
	return res
}

func (l *Lect) AddEx(sub string, shape func(*ebiten.Image), anim, xact func()) {
	x := ex{
		sub: l.doSub(sub),
		shape: shape, anim: anim, xact: xact,
	}
	l.exs = append(l.exs, x)
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
	l.exs[l.n].shape(screen)

	for i, v := range l.exs[l.n].sub {
		x, y, _, _ := l.b.subs[i].Rect()
		l.Font.Set(15, clrs.White)
		l.Font.Draw(screen, v, x, y)
	}

	x, y := l.b.pause.MidF32()
	l.Font.Set(50, clrs.White)
	if l.play {
		l.Font.DrawCenter(screen, "▮▮", x, y)
	} else {
		l.Font.DrawCenter(screen, "▶", x, y)
	}

	//l.b.top.WalkUp(graph.TestDraw(screen))
}

func (l *Lect) Update() {
	if l.play { l.exs[l.n].anim() }
	l.exs[l.n].xact()
	if l.Done && !l.Tracks.IsPlaying() { l.Next() }
}
