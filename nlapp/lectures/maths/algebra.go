package maths

import (
	//"embed"
	//"math"
	// "strconv"

	"github.com/hajimehoshi/ebiten/v2"

	//"github.com/not-learning/app/clrs"
	"github.com/not-learning/app/nlapp/frame"
)

type Algebra struct {
	*frame.Frame
}

// ## Ex 1
func (a *Algebra) sub1(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Вот ящик.`,
	)
}

func (a *Algebra) shape1(scr *ebiten.Image) {
	a.PlayConShow()
	// TODO
}

func (a *Algebra) anim1() bool {
	// TODO
	return true
}

func (a *Algebra) xact1() bool { return true }
func (a *Algebra) zero1() {}

// ## Ex 2
func (a *Algebra) sub2(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Положим в него три банана.`,
	)
}

func (a *Algebra) shape2(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim2() bool {
	// TODO
	return true
}

func (a *Algebra) xact2() bool { return true }
func (a *Algebra) zero2() {}

// ## Ex 3
func (a *Algebra) sub3(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Откроем.`,
	)
}

func (a *Algebra) shape3(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim3() bool {
	// TODO
	return true
}

func (a *Algebra) xact3() bool { return true }
func (a *Algebra) zero3() {}

// ## Ex 4
func (a *Algebra) sub4(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`А там пять бананов!`,
	)
}

func (a *Algebra) shape4(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim4() bool {
	// TODO
	return true
}

func (a *Algebra) xact4() bool { return true }
func (a *Algebra) zero4() {}

// ## Ex 5
func (a *Algebra) sub5(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Внимание, вопрос.`,
	)
}

func (a *Algebra) shape5(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim5() bool {
	// TODO
	return true
}

func (a *Algebra) xact5() bool { return true }
func (a *Algebra) zero5() {}

// ## Ex 6
func (a *Algebra) sub6(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Сколько было бананов в ящике сначала?`,
	)
}

func (a *Algebra) shape6(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim6() bool {
	// TODO
	return true
}

func (a *Algebra) xact6() bool { return true }
func (a *Algebra) zero6() {}

// ## Ex 7
func (a *Algebra) sub7(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Так выглядела математика тысячи лет, а потом люди придумали для нее свой язык.`,
	)
}

func (a *Algebra) shape7(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim7() bool {
	// TODO
	return true
}

func (a *Algebra) xact7() bool { return true }
func (a *Algebra) zero7() {}

/*
// ## Ex 1
func (a *Algebra) sub1(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Возьмем ящик.`,
	)
}

func (a *Algebra) shape1(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim1() bool {
	// TODO
	return true
}

func (a *Algebra) xact1() bool { return true }
func (a *Algebra) zero1() {}
*/
