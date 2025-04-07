package frame

import (
	c "eb/clrs"
	f "eb/fonts"
	"eb/nump"
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

	exercises [][]exercise
	curEx     int
	curExL    int
	funlist   []func() bool
	curFun    int

	script    []string
	curScript int
	subtitle *l.Lmnt

	nump *nump.NumP
	font *f.Font
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

	lower := l.New()
	lower.SetRect(x1, x2, x2, y2)

	sub := l.New()
	sub.SetRow()
	lower.Add(sub)
	padSub := l.New()
	lc.subtitle = l.New()
	padSub.SetSize(20, 0)
	sub.Add(padSub, lc.subtitle)
	//w, h := lc.subtitles[i].TextSize(s) //TODO: wrap
	lc.subtitle.SetSize(0, 17)

	lower.DoAll()

	lc.luma = 255

	lc.seqFunInit()
	lc.nump = nump.Init(x1, y1, x2, y2)

	lc.font = f.InitFont()

	return lc
}

func (lc *Lecture) ScriptInit(script []string) {
	lc.script = script
}

func (lc *Lecture) doEx(txt [][]string) {
	ex := []exercise{}
	for _, v := range txt {
		ans, problem := 0, false
		if v[1] > "" {
			a, e := strconv.Atoi(v[1])
			if e != nil {log.Println("doEx:", e)}
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

func (lc *Lecture) input() bool {
	if !lc.exercises[lc.curExL][lc.curEx].problem { return true }
	str := "?"
	if lc.nump.Num() != 0 { str = lc.nump.Str() }
	lc.scrStr[len(lc.scrStr)-1] = lc.exercises[lc.curExL][lc.curEx].text + str

	if lc.nump.Check(lc.exercises[lc.curExL][lc.curEx].ans) { return true }
	return false
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
		lc.nextScript,
		lc.nextEx,
	}
}

func (lc *Lecture) Update() {
	lc.nump.Update()
	if lc.funlist[lc.curFun]() { lc.curFun++ }
	if lc.curFun >= len(lc.funlist) { lc.curFun = 0 }
}

func (lc *Lecture) Draw(scr *ebiten.Image) {
	//lc.upper.WalkUp(l.TestDraw(scr))
	if lc.exercises[lc.curExL][lc.curEx].problem { lc.nump.Draw(scr) }

	x, y, _, _ := lc.subtitle.Rect()
	lc.font.Set(15, c.YCC(lc.luma, 128, 128))
	lc.font.Draw(scr, lc.script[lc.curScript], x, y)

	for i := range lc.screen {
		//l.Draw(scr, lc.screen[i])
		
		x, y, _, _ := lc.screen[i].Rect()
		lc.font.Set(50, c.YCC(lc.luma, 128, 128))
		lc.font.Draw(scr, lc.scrStr[i], x, y)
	}
}

//font.Metrics()
//font.Advance()
