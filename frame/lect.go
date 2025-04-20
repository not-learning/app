package frame

import (
	"github.com/not-learning/app/draw"
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
}

func Init(x1, y1, x2, y2 float32) *Lect {
	l := &Lect{}
	l.Img = draw.Init(x1, y1, x2, y2)
	return l
} //*/

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

func (l *Lect) Draw(screen *ebiten.Image) {
	l.shapes[l.ex.n](screen)
}

func (l *Lect) Space() bool {
	return ebiten.IsKeyPressed(ebiten.KeySpace)
}

func (l *Lect) MouseL() bool {
	return ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
}

func (l *Lect) MouseR() bool {
	return ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight)
}

/*func (l *Lect) Touch() {

}//*/

func (l *Lect) Update() {
	l.anims[l.ex.n]()
	l.xacts[l.ex.n]()
}
