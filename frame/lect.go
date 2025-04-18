package frame

import (
	"github.com/not-learning/app/tracks"
	"github.com/not-learning/app/draw"
	"github.com/hajimehoshi/ebiten/v2"
)

type lects struct {
	tracks  *tracks.Tracks
	subs    [][]string
	n int
}

type ex struct {
	texts   []string
	shapes  []func(*ebiten.Image)
	anims   []func()
	xacts   []func()
	n int
}

type Lect struct {
	lects
	ex
	*draw.Img
}

func Init(x1, y1, x2, y2 float32) *Lect {
	l := &Lect{}
	l.Img = draw.Init(x1, y1, x2, y2)
	return l
}//*/

func (l *Lect) AddShape(fn func(*ebiten.Image)) {
	l.shapes = append(l.shapes, fn)
}

func (l *Lect) Next() {
	if l.lects.n < len(l.lects.subs)-1 { l.lects.n++ }
	if l.ex.n < len(l.ex.texts)-1 { l.ex.n++ }
	l.tracks.Next()
}

func run(screen *ebiten.Image, fn func(*ebiten.Image)) {
	if fn != nil { fn(screen) }
}

func (l *Lect) Draw(screen *ebiten.Image) {
	run(screen, l.shapes[l.ex.n])
}
