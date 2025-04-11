package frame

import (
	c "github.com/not-learning/app/clrs"
	f "github.com/not-learning/app/fonts"
	"github.com/not-learning/app/nump"
	"github.com/not-learning/app/tracks"

	"log"
	"strconv"

	l "github.com/not-learning/lmnts"
	"github.com/hajimehoshi/ebiten/v2"
)

// TODO: text size
type Lecture struct {
	luma uint8

	upper   *l.Lmnt
	inner   *l.Lmnt
	content *l.Lmnt
	screen  []*l.Lmnt
	scrStr  []string

	play bool

	exercises [][]exercise
	curEx     int
	curExL    int
	funlist   []func() bool
	curFun    int

	script     [][]string
	curScript  int
	subL       *l.Lmnt
	subContent *l.Lmnt
	subCards   []*l.Lmnt
	subtitles  [][]*l.Lmnt

	btns []*l.Lmnt

	nump *nump.NumP
	font *f.Font
	tracks *tracks.Tracks
}

type exercise struct {
	text    string
	ans     int
	problem bool
}

func Init(x1, y1, x2, y2 float32) *Lecture {
	lc := &Lecture{}

	lc.upper = l.New()
	lc.upper.SetRect(x1, y1, x2, x2)

	lc.inner = l.New()
	lc.inner.Name = "in"
	lc.inner.SetRow()
	lc.inner.SetSize(0, 50)
	lc.upper.GapsAround(0, lc.inner)
	pad1 := l.New()
	pad1.SetSize(50, 0)
	lc.content = l.New()
	lc.inner.Add(pad1, lc.content)

	lc.upper.DoAll()

	lc.subL = l.New()
	lc.subL.SetRect(x1, x2, x2, y2)

	sub := l.New()
	sub.SetRow()
	//sub.SetSize(0, 17)
	lc.subL.Add(sub)
	padSub := l.New()
	padSub.SetSize(20, 0)
	lc.subContent = l.New()
	sub.Add(padSub, lc.subContent)

	btnRow := l.New()
	btnRow.SetRow()
	lc.btns = make([]*l.Lmnt, 3)
	for i := range 3 {
		lc.btns[i] = l.New()
		lc.btns[i].SetSize(150, 75)
	}
	lc.btns[1].SetSize(75, 75)
	btnRow.GapsAround(0, lc.btns...)
	btnMargin := l.New()
	btnMargin.SetSize(0, 20)
	lc.subL.Add(btnMargin, btnRow)

	//lc.subtitle.SetSize(0, 17)

	lc.subL.DoAll()

	

	lc.luma = 255

	lc.seqFunInit()
	lc.nump = nump.Init(x1, y1, x2, y2)

	lc.font = f.InitFont()

	lc.tracks = tracks.Init()

	lc.nump.Play()

	return lc
}

func (lc *Lecture) ScriptInit(script []string) {
	var res = [][]string{}

	lc.font.Set(15, c.YCC(lc.luma, 128, 128))
	x1, _, x2, _ := lc.subContent.Rect()
	w := x2 - x1
	for i, v := range script {
		res = append(res, []string{})
		lc.subCards = append(lc.subCards, l.New())
		lc.subtitles = append(lc.subtitles, []*l.Lmnt{})
		for j, w := range lc.font.Wrap(v, float64(w)) {
			lc.subtitles[i] = append(lc.subtitles[i], l.New())
			lc.subtitles[i][j].SetSize(0, 17)
			res[i] = append(res[i], w...)
		}
		lc.subCards[i].GapsBetween(5, lc.subtitles[i]...)
	}
	lc.subContent.GapsBetween(15, lc.subCards...)

	lc.subL.DoAll()
	lc.script = res
}

func (lc *Lecture) doEx(txt [][]string) {
	ex := []exercise{}
	for _, v := range txt {
		ans, problem := 0, false
		if v[1] > "" {
			a, e := strconv.Atoi(v[1])
			if e != nil {log.Println("doEx: ", e)}
			ans, problem = a, true
		}

		ex = append(ex, exercise{
			text: v[0],
			ans: ans,
			problem: problem,
		})
	}
	lc.exercises = append(lc.exercises, ex)
}

func (lc *Lecture) ExInit(txt ...[][]string) {
	for _, v := range txt { lc.doEx(v) }
}

const fs = 5 // TODO: move somewhere safe

func (lc *Lecture) fade() bool {
	lc.luma -= fs
	if lc.luma < fs {return true}
	return false
}

func (lc *Lecture) unfade() bool {
	lc.luma += fs
	if lc.luma > 255-fs {return true}
	return false
}

func (lc *Lecture) addScreen() {
	ln := len(lc.screen)
	lc.screen = append(lc.screen, l.New())
	lc.screen[ln].SetSize(0, 50)
	lc.content.Add(lc.screen[ln])
	lc.inner.SetSize(0, float32(50*(ln+1)))
	lc.upper.DoAll()
}

func (lc *Lecture) clearScreen() {
	lc.screen = []*l.Lmnt{}
	lc.content.Clear()
	lc.upper.DoAll()
}//*/

func (lc *Lecture) appendScrStr() {
	if lc.exercises[lc.curExL][lc.curEx].problem { lc.nump.NumClear() }
	lc.scrStr = append(lc.scrStr, lc.exercises[lc.curExL][lc.curEx].text)
}

func (lc *Lecture) nextScript() bool {
	lc.curScript++
	if lc.curScript >= len(lc.script) { lc.curScript = len(lc.script)-1 }
	return true
}

func (lc *Lecture) nextEx() bool {
	lc.curEx++
	if lc.curEx >= len(lc.exercises[lc.curExL]) {
		lc.curEx = 0
		lc.scrStr = []string{}
		lc.clearScreen()
		lc.curExL++
		if lc.curExL >= len(lc.exercises) { lc.curExL = len(lc.exercises)-1 }
	}

	if lc.exercises[lc.curExL][lc.curEx].text != "" {
		lc.addScreen()
		lc.appendScrStr()
	}
	return true
}

func (lc *Lecture) next() bool {
	if lc.tracks.IsPlaying() || !lc.nump.IsPlaying() { return false }
	lc.tracks.Next()
	lc.nextScript()
	lc.nextEx()
	return true
}

func (lc *Lecture) playStr() string {
	if lc.nump.IsPlaying() {
		return "▮▮"
	} else {
		return "▶"
	}
}
//▮▮▶▯▯▷

func (lc *Lecture) correct() {
	lc.tracks.PlayCorrect()
}

func (lc *Lecture) wrong() {
	lc.tracks.PlayWrong()
}

func (lc *Lecture) input() bool {
	if !lc.exercises[lc.curExL][lc.curEx].problem { return true }
	str := "?"
	if lc.nump.Num() != 0 { str = lc.nump.Str() }
	lc.scrStr[len(lc.scrStr)-1] = lc.exercises[lc.curExL][lc.curEx].text + str

	return lc.nump.Check(
		lc.exercises[lc.curExL][lc.curEx].ans,
		lc.correct,
		lc.wrong,
	)
}

func (lc *Lecture) end() bool {
	lc.scrStr = []string{`The End`}
	return false
}

func (lc *Lecture) seqFunInit() {
	lc.funlist = []func() bool{
		//lc.unfade,
		lc.input,
		//lc.fade,
		lc.next,
	}
}

func (lc *Lecture) Update() {
	lc.nump.Update()
	if lc.funlist[lc.curFun]() { lc.curFun++ }
	if lc.curFun >= len(lc.funlist) { lc.curFun = 0 }
}

func (lc *Lecture) Draw(scr *ebiten.Image) {
	//.WalkUp(l.TestDraw(scr))
	if lc.exercises[lc.curExL][lc.curEx].problem { lc.nump.Draw(scr) }

//lc.subL.WalkUp(l.TestDraw(scr))
	for i, v := range lc.subtitles {
	//l.Draw(scr, lc.subCards[i])
		for j, w := range v {
			x, y, _, _ := w.Rect()
			lc.font.Set(15, c.YCC(lc.luma, 128, 128))
			lc.font.Draw(scr, lc.script[i][j], x, y)
		}
	}//*/

	for i := range lc.screen {
		//l.Draw(scr, lc.screen[i])
		x, y, _, _ := lc.screen[i].Rect()
		lc.font.Set(50, c.YCC(lc.luma, 128, 128))
		lc.font.Draw(scr, lc.scrStr[i], x, y)
	}

	lc.font.Set(70, c.YCC(lc.luma, 128, 128))
	x, y := lc.btns[0].MidF32()
	lc.font.DrawCenter(scr, "◄", x, y-7) //TODO: center properly
	lc.font.Set(50, c.YCC(lc.luma, 128, 128))
	x, y = lc.btns[1].MidF32()
	lc.font.DrawCenter(scr, lc.playStr(), x, y-5) //TODO: center properly
	lc.font.Set(70, c.YCC(lc.luma, 128, 128))
	x, y = lc.btns[2].MidF32()
	lc.font.DrawCenter(scr, "►", x, y-7) //TODO: center properly
}
//☞☜☛☚
//◄► ◅▻

func (lc *Lecture) Play()  { lc.nump.Play() }
func (lc *Lecture) Stop()  { lc.nump.Stop() }
func (lc *Lecture) Pause() { lc.nump.Pause() }//*/

//font.Metrics()
//font.Advance()
