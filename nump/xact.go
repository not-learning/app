package nump

import (
	"os"

	input "github.com/quasilyte/ebitengine-input"
)

type xAct struct {
	inSys       input.System
	inHand     *input.Handler
	btnPos   []*pos
	btnErase   *pos
	btnCheck   *pos

	play *bool

	*int
	confirm *bool
}

type pos struct{ x1, y1, x2, y2 float64 }

const ( // todo: make it a list maybe?
	actionInsert0 input.Action = iota
	actionInsert1
	actionInsert2
	actionInsert3
	actionInsert4
	actionInsert5
	actionInsert6
	actionInsert7
	actionInsert8
	actionInsert9
	actionBackspace
	actionEnter
	actionPause

	actionClick

	actionQuit
)

var keymap = input.Keymap{
	actionInsert0:   {input.Key0, input.KeyNum0},
	actionInsert1:   {input.Key1, input.KeyNum1},
	actionInsert2:   {input.Key2, input.KeyNum2},
	actionInsert3:   {input.Key3, input.KeyNum3},
	actionInsert4:   {input.Key4, input.KeyNum4},
	actionInsert5:   {input.Key5, input.KeyNum5},
	actionInsert6:   {input.Key6, input.KeyNum6},
	actionInsert7:   {input.Key7, input.KeyNum7},
	actionInsert8:   {input.Key8, input.KeyNum8},
	actionInsert9:   {input.Key9, input.KeyNum9},
	actionBackspace: {input.KeyBackspace},

	actionPause: {input.KeySpace},

	actionClick: {input.KeyMouseLeft},
	actionEnter: {input.KeyEnter, input.KeyNumEnter},
	actionQuit:  {input.KeyEscape},
}

func (x *xAct) keyRepeat(action input.Action) bool {
	ev, ok := x.inHand.PressedActionInfo(action)
	if !ok {return false}

	dur := ev.Duration
	delay, interval := 20, 3
	if dur == 1 { return true }
	if dur > delay && (dur-delay)%interval == 0 { return true }
	return false
}

func (x *xAct) clickInside(p *pos) bool {
	if info, ok := x.inHand.JustPressedActionInfo(actionClick); ok {
		xpos, ypos := info.Pos.X, info.Pos.Y
		return p.x1 < xpos && xpos < p.x2 &&
					 p.y1 < ypos && ypos < p.y2
	}
	return false
}

// TODO something about overflow
func (x *xAct) update() {
	x.inSys.Update()

	if x.keyRepeat(actionInsert0) {*x.int = *x.int*10 + 0}
	if x.keyRepeat(actionInsert1) {*x.int = *x.int*10 + 1}
	if x.keyRepeat(actionInsert2) {*x.int = *x.int*10 + 2}
	if x.keyRepeat(actionInsert3) {*x.int = *x.int*10 + 3}
	if x.keyRepeat(actionInsert4) {*x.int = *x.int*10 + 4}
	if x.keyRepeat(actionInsert5) {*x.int = *x.int*10 + 5}
	if x.keyRepeat(actionInsert6) {*x.int = *x.int*10 + 6}
	if x.keyRepeat(actionInsert7) {*x.int = *x.int*10 + 7}
	if x.keyRepeat(actionInsert8) {*x.int = *x.int*10 + 8}
	if x.keyRepeat(actionInsert9) {*x.int = *x.int*10 + 9}

	for i, v := range x.btnPos {
			if x.clickInside(v) {
					*x.int = *x.int*10 + i
			}
	}

	*x.confirm = x.inHand.ActionIsJustPressed(actionEnter) || x.clickInside(x.btnCheck)

	if x.inHand.ActionIsJustPressed(actionPause) { *x.play = !*x.play }

	if x.clickInside(x.btnErase) { *x.int /= 10 }
	if x.keyRepeat(actionBackspace) { *x.int /= 10 }

	if x.inHand.ActionIsPressed(actionQuit) { os.Exit(0) }
}

func initXact() *xAct {
	x := &xAct{}
	x.inSys.Init(input.SystemConfig{DevicesEnabled: input.AnyDevice})
	x.inHand = x.inSys.NewHandler(0, keymap)

	i := 0
	x.int = &i
	b := false
	x.confirm = &b
	p := false
	x.play = &p
	return x
}

func (x *xAct) isPlaying() bool { return *x.play }
