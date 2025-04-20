package frame

import (
	"strings"

	"github.com/not-learning/app/clrs"
	"github.com/not-learning/app/draw"
	"github.com/not-learning/app/fonts"
	"github.com/not-learning/app/tracks"

	"github.com/hajimehoshi/ebiten/v2"
)

type lects struct {
	tracks *tracks.Tracks
	subs   [][]string
	n      int
}

type ex struct {
	texts  []string
	shapes []func(*ebiten.Image)
	anims  []func()
	xacts  []func()
	n      int
}

type Lect struct {
	lects
	ex
	*draw.Img
	*fonts.Font
	b *blocks
}

func Init(x1, y1, x2, y2 float32) *Lect {
	l := &Lect{}
	l.Img = draw.Init(x1, y1, x2, y2)
	l.Font = fonts.InitFont()
	l.b = InitBlocks(x1, y1, x2, y2)
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

func (l *Lect) AddShape(fn func(*ebiten.Image)) {
	l.shapes = append(l.shapes, fn)
}

func (l *Lect) AddAnim(fn func()) {
	l.anims = append(l.anims, fn)
}

func (l *Lect) AddXact(fn func()) {
	l.xacts = append(l.xacts, fn)
}

func (l *Lect) Next() {
	if l.lects.n <= len(l.lects.subs) {
		l.lects.n++
	}
	if l.ex.n <= len(l.ex.texts) {
		l.ex.n++
	}
	l.tracks.Next()
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
}

func (l *Lect) Update() {
	l.anims[l.ex.n]()
	l.xacts[l.ex.n]()
}
