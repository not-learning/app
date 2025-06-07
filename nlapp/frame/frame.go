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

	subH, pbcH, npH float32

	SetChapter func(int)

	*numpad

	*blocks
	*fonts.Font
	*graph.Graph
	inter.Inter
	*tracks.Tracks
}

func Init(x1, y1, x2, y2, scale float32) *Frame {
	f := &Frame{}
	f.numpad = &numpad{}
	f.blocks = initBlocks(x1, y1, x2, y2)
	f.Font = fonts.InitFont(scale)
	f.Graph = graph.Init()
	f.SetEdges(f.blocks.screen.Rect())
	f.Tracks = tracks.Init()

	f.subH = 15
	f.pbcH = 40
	f.npH = 40

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

/*func (f *Frame) SetSizes(subH, pbcH, npH float32) {
	f.subH = float64(subH)
	f.pbcH = pbcH
	f.npH = npH
}//*/

func (f *Frame) Update(scrW, scrH int, scale float32) {
	if f.Escape() { os.Exit(0) }

	if f.Space() || f.TapIn(f.blocks.pause.Rect()) {
		f.Pause()
	}

	if f.ArrowLeft() || f.TapIn(f.blocks.prev.Rect()) {
		f.Prev()
		return
	}
	if f.ArrowRight() || f.TapIn(f.blocks.next.Rect()) {
		f.Next()
		return
	}

	f.Font.Update(scale)
	i := scale*1.5
	f.blocks.subL.SetSize(0, f.subH*i*2)
	f.blocks.pbc.SetSize(0, f.pbcH*i)
	f.blocks.np.SetSize(0, f.npH*i*4)
	f.blocks.update(scrW, scrH)
	f.Graph.Update(f.blocks.screen.Rect())
	f.Proceed()
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
	x1, x2 = x-w/2, x+w/2
	y1, y2 = y-h/2, y+h/2
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
	w := x2-x1 - 20 // TODO proper sizes
	res = append(res, f.Font.Wrap(f.subH, w, sub)...)
	return res
}

func (f *Frame) subDraw(sub string) func(*ebiten.Image) {
	return func(scr *ebiten.Image) {
		sl := f.subWrap(sub)
		for i, v := range sl {
			x, y, _, _ := f.blocks.subs[i].Rect()
			f.Font.Draw(scr, v, float32(f.subH), x, y, clrs.White)
		}
	}
}

func (f *Frame) Sub(sub string) func(*ebiten.Image) {
	return f.subDraw(sub)
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

func (f *Frame) Input() int {
	if n, ok := f.Number(); ok {
		f.numpad.input = n + f.numpad.input*10
	}

	for i, v := range f.blocks.npl {
		if i == 9 || i == 11 { continue }
		if f.TapIn(v.Rect()) {
			f.numpad.input = f.numpad.input*10 + f.numpad.num[i]
		}
	}
	return f.numpad.input
}

func (f *Frame) Erase() {
	if f.TapIn(f.blocks.npl[9].Rect()) || f.Backspace() {
		f.numpad.input /= 10
	}
}

// todo: not very good
func (f *Frame) Check(solution int) (correct, ok bool) {
	correct = f.numpad.input == solution
	ok = f.TapIn(f.blocks.npl[11].Rect()) || f.Enter()
	//if ok { f.PlayCorrect(correct) }
	return
}

func (f *Frame) Draw(screen *ebiten.Image) {
	f.exs[f.exn].shape(screen) // TODO fix: no shapes -> panic
	f.exs[f.exn].sub(screen)
	//graph.Draw(screen, f.blocks.prev, f.blocks.pause, f.blocks.next)
	//f.blocks.top.WalkDown(graph.TestDraw(screen))

	if f.blocks.pbcshow {
		x, y := f.blocks.prev.MidF32()
		f.Font.DrawCenter(screen, "⮜", f.pbcH, x, y, clrs.White)
		x, y = f.blocks.next.MidF32()
		f.Font.DrawCenter(screen, "⮞", f.pbcH, x, y, clrs.White)

		x, y = f.blocks.pause.MidF32()
		if f.play {
			f.Font.DrawCenter(screen, "▮▮", f.pbcH, x, y, clrs.White)
		} else {
			f.Font.DrawCenter(screen, "▶", f.pbcH, x, y, clrs.White)
		}
	}

	if f.blocks.npshow {
		//graph.Draw(screen, f.blocks.npf...)
		for i, v := range f.blocks.npl {
			x, y := v.MidF32()
			f.Font.DrawCenter(screen, f.numpad.str[i], f.npH, x, y, clrs.White)
		}
	}
}

// ←→
//⬅➡⮕ ⇦ ⇨ 🡐 🡒 ⇽ ⇾ ⟵ ⟶ ⟸ ⟹ ⟻ ⟼ ⇜ ⇝ ⬿ ⤳ ↫ ↬ ⮜ ⮞ ⮪ ⮫ ◄ ► ◅ ▻
