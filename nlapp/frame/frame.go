package frame

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/not-learning/app/nlapp/clrs"
	"github.com/not-learning/app/nlapp/fonts"
	"github.com/not-learning/app/nlapp/graph"
	"github.com/not-learning/app/nlapp/inter"
	"github.com/not-learning/app/nlapp/tracks"
)

type ex struct {
	sub, shape func(*ebiten.Image)
	anim, xact func() bool
	zero       func()
}

// TODO numpad out
type numpad struct {
	str []string
	num []int

	input int
}

type Frame struct {
	play bool
	exs  []ex
	exn    int

	SetChapter func(int)

	X1, Y1, X2, Y2 float32 // DEV

	*numpad

	*blocks
	*fonts.Font
	*graph.Graph
	inter.Inter
	*tracks.Tracks
}

func Init(x1, y1, x2, y2 float32) *Frame {
	f := &Frame{}
	f.numpad = &numpad{}
	f.blocks = initBlocks(x1, y1, x2, y2)
	f.Font = fonts.InitFont()
	f.Graph = graph.Init()
	f.Graph.SetOrigin(f.blocks.screen.MidF32())
	f.Tracks = tracks.Init()

	f.numpad.str = []string{
		"7", "8", "9",
		"4", "5", "6",
		"1", "2", "3",
		"⌫", "0", "✓",
	}

	f.numpad.num = []int{
		7, 8, 9,
		4, 5, 6,
		1, 2, 3,
		0, 0,
	}

	// f.play = true // TODO DEV
	//f.Pause() // TODO DEV
	return f
}

func (f *Frame) ChapterFn(fn func(int)) { f.SetChapter = fn }

func (f *Frame) Label(
	scr *ebiten.Image,
	text string,
	size, x, y float32,
	clr clrs.Clr,
) {
	x, y = f.Coords(x, y)
	f.DrawCenter(scr, text, size, x, y, clr)
	return
}

func (f *Frame) LabelRect(
	scr *ebiten.Image,
	text string,
	size, x, y float32,
	clr clrs.Clr,
) (x1, y1, x2, y2 float32) {
	x, y = f.Coords(x, y)
	f.DrawCenter(scr, text, size, x, y, clr)
	w, h := f.TextSize(text)
	x1, x2 = x-float32(w/2), x+float32(w/2)
	y1, y2 = y-float32(h/2), y+float32(h/2)
	return
}

func (f *Frame) Play() {
	f.play = true
	f.Tracks.Play()
}

func (f *Frame) Pause() {
	f.play = !f.play
	f.Tracks.Pause()
}

func (f *Frame) Proceed() {
	var anim, xact bool
	if f.play { anim = f.exs[f.exn].anim() }
	xact = f.exs[f.exn].xact()

	if pl, ok := f.Tracks.IsPlaying()
	xact && anim && f.play && !pl && ok {
		if f.exn < len(f.exs)-1 {
			f.exs[f.exn].zero()
			f.exn++
		}
		f.Tracks.Proceed()
	}
}

func (f *Frame) Next() {
	f.play = false
	f.exs[f.exn].zero()
	if f.exn < len(f.exs)-1 { f.exn++ }
	f.exs[f.exn].zero()
	f.Tracks.Switch(f.exn)
}

func (f *Frame) Prev() {
	f.play = false
	f.exs[f.exn].zero()
	if f.exn > 0 {
		f.exn--
	} else {
		f.SetChapter(0)
	}
	f.exs[f.exn].zero()
	f.Tracks.Switch(f.exn)
}

func (f *Frame) subWrap(sub string) []string {
	res := []string{}
	x1, _, x2, _ := f.blocks.subs[0].Rect()
	w := float64(x2-x1) - 20 // TODO proper sizes
	res = append(res, f.Font.Wrap(15, w, sub)...)
	return res
}

func (f *Frame) subDraw(sub []string) func(*ebiten.Image) {
	return func(scr *ebiten.Image) {
		for i, v := range sub {
			x, y, _, _ := f.blocks.subs[i].Rect()
			f.Font.Draw(scr, v, 15, x, y, clrs.White)
		}
	}
}

func (f *Frame) Sub(sub string) func(*ebiten.Image) {
	return f.subDraw(f.subWrap(sub))
}

func (f *Frame) AddEx(
	sub, shape func(*ebiten.Image), 
	anim, xact func() bool,
	zero func(),
) {
	x := ex{sub: sub, shape: shape, anim: anim, xact: xact, zero: zero}
	f.exs = append(f.exs, x)
}

func (f *Frame) NumPadShow() { f.blocks.numPadShow() }

func (f *Frame) NumPadHide() { f.blocks.numPadHide() }

func (f *Frame) PlayConShow() { f.blocks.playConShow() }

func (f *Frame) PlayConHide() { f.blocks.playConHide() }

/*func (f *Frame) Touch() {}//*/

func (f *Frame) Input() int {
	if n, ok := f.Number(); ok {
		f.numpad.input = n + f.numpad.input*10
	}

	for i, v := range f.blocks.npl {
		if i == 9 || i == 11 { continue }
		if f.MouseLIn(v.Rect()) {
			f.numpad.input = f.numpad.input*10 + f.numpad.num[i]
		}
	}
	return f.numpad.input
}

func (f *Frame) Erase() {
	if f.MouseLIn(f.blocks.npl[9].Rect()) || f.Backspace() {
		f.numpad.input /= 10
	}
}

// todo: not very good
func (f *Frame) Check(solution int) (correct, ok bool) {
	correct = f.numpad.input == solution
	ok = f.MouseLIn(f.blocks.npl[11].Rect()) || f.Enter()
	if ok { f.PlayCorrect(correct) }
	return
}

func (f *Frame) Update(scrW, scrH int, ratW, ratH float32) {
	f.X2, f.Y2 = float32(scrW), float32(scrH) // DEV
	f.blocks.update(scrW, scrH, ratW, ratH)
	f.Font.Update(scrW, scrH, ratW, ratH)

	if f.Escape() { os.Exit(0) }
	if f.Space() || f.MouseLIn(f.blocks.pause.Rect()) { f.Pause() }

	if f.ArrowLeft() || f.MouseLIn(f.blocks.prev.Rect()) {
		f.Prev()
		return
	}
	if f.ArrowRight() || f.MouseLIn(f.blocks.next.Rect()) {
		f.Next()
		return
	}

	f.Proceed()
}

func (f *Frame) Draw(screen *ebiten.Image) {
	f.exs[f.exn].shape(screen)
	f.exs[f.exn].sub(screen)
	//graph.Draw(screen, f.blocks.prev, f.blocks.pause, f.blocks.next)
	//f.blocks.top.WalkDown(graph.TestDraw(screen))

	if f.blocks.pbcshow {
		x, y := f.blocks.prev.MidF32()
		f.Font.DrawCenter(screen, "←", 75, x, y, clrs.White)
		x, y = f.blocks.next.MidF32()
		f.Font.DrawCenter(screen, "→", 75, x, y, clrs.White)

		x, y = f.blocks.pause.MidF32()
		if f.play {
			f.Font.DrawCenter(screen, "▮▮", 50, x, y, clrs.White)
		} else {
			f.Font.DrawCenter(screen, "▶", 50, x, y, clrs.White)
		}
	}

	if f.blocks.npshow {
		//graph.Draw(screen, f.blocks.npf...)
		for i, v := range f.blocks.npl {
			x, y := v.MidF32()
			f.Font.DrawCenter(screen, f.numpad.str[i], 50, x, y, clrs.White)
		}
	}
}

// ←→
//⬅➡⮕ ⇦ ⇨ 🡐 🡒 ⇽ ⇾ ⟵ ⟶ ⟸ ⟹ ⟻ ⟼ ⇜ ⇝ ⬿ ⤳ ↫ ↬ ⮜ ⮞ ⮪ ⮫ ◄ ► ◅ ▻
