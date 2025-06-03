package maths

import (
	//"embed"
	//"math"
	// "strconv"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/not-learning/app/nlapp/clrs"
	"github.com/not-learning/app/nlapp/frame"
)

type Algebra struct {
	*frame.Frame
}

// ## Ex 1
func (a *Algebra) sub1(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Начнем с простого.`,
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
		`Представь весы,`,
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
		`на одной чаше арбуз и три гирьки по килограмму,`,
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
		`а на другой десять гирек по килограмму.`,
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
		`Сколько весит арбуз?`,
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
		`Возьмем посложнее.`,
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
		`Пусть на одной чаше весов будут`,
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

// ## Ex 8
func (a *Algebra) sub8(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`два одинаковых арбуза и гирька 5 кг.`,
	)
}

func (a *Algebra) shape8(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim8() bool {
	// TODO
	return true
}

func (a *Algebra) xact8() bool { return true }
func (a *Algebra) zero8() {}

// ## Ex 9
func (a *Algebra) sub9(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`а на другой - одна гиря на 27 килограмм.`,
	)
}

func (a *Algebra) shape9(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim9() bool {
	// TODO
	return true
}

func (a *Algebra) xact9() bool { return true }
func (a *Algebra) zero9() {}

// ## Ex 10
func (a *Algebra) sub10(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Сколько весит один арбуз?`,
	)
}

func (a *Algebra) shape10(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim10() bool {
	// TODO
	return true
}

func (a *Algebra) xact10() bool { return true }
func (a *Algebra) zero10() {}

/*// ## Ex 11
func (a *Algebra) sub11(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Задачки простые, потому что мы изучаем новый язык,`,
	)
}

func (a *Algebra) shape11(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim11() bool {
	// TODO
	return true
}

func (a *Algebra) xact11() bool { return true }
func (a *Algebra) zero11() {}

// ## Ex 12
func (a *Algebra) sub12(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`язык математики - алгебру.`,
	)
}

func (a *Algebra) shape12(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim12() bool {
	// TODO
	return true
}

func (a *Algebra) xact12() bool { return true }
func (a *Algebra) zero12() {}*/

// ## Ex 13
func (a *Algebra) sub13(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Попробуем описать эти задачи алгебраически.`,
	)
}

func (a *Algebra) shape13(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim13() bool {
	// TODO
	return true
}

func (a *Algebra) xact13() bool { return true }
func (a *Algebra) zero13() {}

// ## Ex 14
func (a *Algebra) sub14(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Обозначим вес арбуза буквой а,`,
	)
}

func (a *Algebra) shape14(scr *ebiten.Image) {
	a.Label(scr, "a", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim14() bool {
	// TODO
	return true
}

func (a *Algebra) xact14() bool { return true }
func (a *Algebra) zero14() {}

// ## Ex 15
func (a *Algebra) sub15(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`и просто запишем, что у нас находится`,
	)
}

func (a *Algebra) shape15(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim15() bool {
	// TODO
	return true
}

func (a *Algebra) xact15() bool { return true }
func (a *Algebra) zero15() {}

// ## Ex 16
func (a *Algebra) sub16(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`на каждой чаше весов.`,
	)
}

func (a *Algebra) shape16(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim16() bool {
	// TODO
	return true
}

func (a *Algebra) xact16() bool { return true }
func (a *Algebra) zero16() {}

// ## Ex 17
func (a *Algebra) sub17(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Начнем с первой задачи.`,
	)
}

func (a *Algebra) shape17(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim17() bool {
	// TODO
	return true
}

func (a *Algebra) xact17() bool { return true }
func (a *Algebra) zero17() {}

// ## Ex 18
func (a *Algebra) sub18(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Слева у нас будет а + 3,`,
	)
}

func (a *Algebra) shape18(scr *ebiten.Image) {
	a.Label(scr, "a + 3", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim18() bool {
	// TODO
	return true
}

func (a *Algebra) xact18() bool { return true }
func (a *Algebra) zero18() {}

// ## Ex 19
func (a *Algebra) sub19(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`а справа просто 10.`,
	)
}

func (a *Algebra) shape19(scr *ebiten.Image) {
	a.Label(scr, "a + 3   10", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim19() bool {
	// TODO
	return true
}

func (a *Algebra) xact19() bool { return true }
func (a *Algebra) zero19() {}

// ## Ex 20
func (a *Algebra) sub20(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Весы находятся в равновесии.`,
	)
}

func (a *Algebra) shape20(scr *ebiten.Image) {
	a.Label(scr, "a + 3   10", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim20() bool {
	// TODO
	return true
}

func (a *Algebra) xact20() bool { return true }
func (a *Algebra) zero20() {}

// ## Ex 21
func (a *Algebra) sub21(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Какой знак мы можем поставить между этими выражениями?`,
	)
}

func (a *Algebra) shape21(scr *ebiten.Image) {
	a.Label(scr, "a + 3   10", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim21() bool {
	// TODO
	return true
}

func (a *Algebra) xact21() bool { return true }
func (a *Algebra) zero21() {}

// ## Ex 22
func (a *Algebra) sub22(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Равно, конечно.`,
	)
}

func (a *Algebra) shape22(scr *ebiten.Image) {
	a.Label(scr, "a + 3 = 10", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim22() bool {
	// TODO
	return true
}

func (a *Algebra) xact22() bool { return true }
func (a *Algebra) zero22() {}

// ## Ex 23
func (a *Algebra) sub23(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Такое выражение называется уравнение.`,
	)
}

func (a *Algebra) shape23(scr *ebiten.Image) {
	a.Label(scr, "a + 3 = 10", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim23() bool {
	// TODO
	return true
}

func (a *Algebra) xact23() bool { return true }
func (a *Algebra) zero23() {}

// ## Ex 24
func (a *Algebra) sub24(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Любая буква в уравнении называется переменная,`,
	)
}

func (a *Algebra) shape24(scr *ebiten.Image) {
	a.Label(scr, "a + 3 = 10", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim24() bool {
	// TODO
	return true
}

func (a *Algebra) xact24() bool { return true }
func (a *Algebra) zero24() {}

// ## Ex 25
func (a *Algebra) sub25(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`а если мы ищем ее значение, называем неизвестная.`,
	)
}

func (a *Algebra) shape25(scr *ebiten.Image) {
	a.Label(scr, "a + 3 = 10", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim25() bool {
	// TODO
	return true
}

func (a *Algebra) xact25() bool { return true }
func (a *Algebra) zero25() {}

// ## Ex 26
func (a *Algebra) sub26(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Итак, нам нужно сделать, чтобы`,
	)
}

func (a *Algebra) shape26(scr *ebiten.Image) {
	a.Label(scr, "a + 3 = 10", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim26() bool {
	// TODO
	return true
}

func (a *Algebra) xact26() bool { return true }
func (a *Algebra) zero26() {}

// ## Ex 27
func (a *Algebra) sub27(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`с одной стороны от равно осталось только неизвестное - буква а.`,
	)
}

func (a *Algebra) shape27(scr *ebiten.Image) {
	a.Label(scr, "a + 3 = 10", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim27() bool {
	// TODO
	return true
}

func (a *Algebra) xact27() bool { return true }
func (a *Algebra) zero27() {}

// ## Ex 28
func (a *Algebra) sub28(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Ну то есть чтобы на одной чаше весов остался только арбуз.`,
	)
}

func (a *Algebra) shape28(scr *ebiten.Image) {
	a.Label(scr, "a + 3 = 10", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim28() bool {
	// TODO
	return true
}

func (a *Algebra) xact28() bool { return true }
func (a *Algebra) zero28() {}

// ## Ex 29
func (a *Algebra) sub29(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Теперь внимание: если мы одновременно`,
	)
}

func (a *Algebra) shape29(scr *ebiten.Image) {
	a.Label(scr, "a + 3 = 10", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim29() bool {
	// TODO
	return true
}

func (a *Algebra) xact29() bool { return true }
func (a *Algebra) zero29() {}


// ## Ex 30
func (a *Algebra) sub30(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`уберем с обеих чаш одинаковый вес,`,
	)
}

func (a *Algebra) shape30(scr *ebiten.Image) {
	a.Label(scr, "a + 3 = 10", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim30() bool {
	// TODO
	return true
}

func (a *Algebra) xact30() bool { return true }
func (a *Algebra) zero30() {}

// ## Ex 31
func (a *Algebra) sub31(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`как изменится равновесие?`,
	)
}

func (a *Algebra) shape31(scr *ebiten.Image) {
	a.Label(scr, "a + 3 = 10", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim31() bool {
	// TODO
	return true
}

func (a *Algebra) xact31() bool { return true }
func (a *Algebra) zero31() {}

// ## Ex 32
func (a *Algebra) sub32(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Никак, верно!`,
	)
}

func (a *Algebra) shape32(scr *ebiten.Image) {
	a.Label(scr, "a + 3 = 10", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim32() bool {
	// TODO
	return true
}

func (a *Algebra) xact32() bool { return true }
func (a *Algebra) zero32() {}

// ## Ex 33
func (a *Algebra) sub33(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Сколько гирек нужно убрать слева,`,
	)
}

func (a *Algebra) shape33(scr *ebiten.Image) {
	a.Label(scr, "a + 3 = 10", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim33() bool {
	// TODO
	return true
}

func (a *Algebra) xact33() bool { return true }
func (a *Algebra) zero33() {}

// ## Ex 34
func (a *Algebra) sub34(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`чтобы остался только арбуз?`,
	)
}

func (a *Algebra) shape34(scr *ebiten.Image) {
	a.Label(scr, "a + 3 = 10", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim34() bool {
	// TODO
	return true
}

func (a *Algebra) xact34() bool { return true }
func (a *Algebra) zero34() {}

// ## Ex 35
func (a *Algebra) sub35(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Три, верно!`,
	)
}

func (a *Algebra) shape35(scr *ebiten.Image) {
	a.Label(scr, "a + 3 = 10", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim35() bool {
	// TODO
	return true
}

func (a *Algebra) xact35() bool { return true }
func (a *Algebra) zero35() {}

// ## Ex 36
func (a *Algebra) sub36(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Значит, столько же уберем и справа.`,
	)
}

func (a *Algebra) shape36(scr *ebiten.Image) {
	a.Label(scr, "a + 3 = 10", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim36() bool {
	// TODO
	return true
}

func (a *Algebra) xact36() bool { return true }
func (a *Algebra) zero36() {}

// ## Ex 37
func (a *Algebra) sub37(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Обозначим это действие в уравнении.`,
	)
}

func (a *Algebra) shape37(scr *ebiten.Image) {
	a.Label(scr, "a + 3 = 10", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim37() bool {
	// TODO
	return true
}

func (a *Algebra) xact37() bool { return true }
func (a *Algebra) zero37() {}


// ## Ex 38
func (a *Algebra) sub38(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Проведи справа от уравнения вертикальную черту`,
	)
}

func (a *Algebra) shape38(scr *ebiten.Image) {
	a.Label(scr, "a + 3 = 10 |", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim38() bool {
	// TODO
	return true
}

func (a *Algebra) xact38() bool { return true }
func (a *Algebra) zero38() {}

// ## Ex 39
func (a *Algebra) sub39(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`и запиши, что делаешь.`,
	)
}

func (a *Algebra) shape39(scr *ebiten.Image) {
	a.Label(scr, "a + 3 = 10 |", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim39() bool {
	// TODO
	return true
}

func (a *Algebra) xact39() bool { return true }
func (a *Algebra) zero39() {}

// ## Ex 40
func (a *Algebra) sub40(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Убираем 3 гирьки, значит минус 3.`,
	)
}

func (a *Algebra) shape40(scr *ebiten.Image) {
	a.Label(scr, "a + 3 = 10 | - 3", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim40() bool {
	// TODO
	return true
}

func (a *Algebra) xact40() bool { return true }
func (a *Algebra) zero40() {}


// ## Ex 41
func (a *Algebra) sub41(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Теперь перепиши уравнение, но`,
	)
}

func (a *Algebra) shape41(scr *ebiten.Image) {
	a.Label(scr, "a + 3 - 3 = 10 - 3 | - 3", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim41() bool {
	// TODO
	return true
}

func (a *Algebra) xact41() bool { return true }
func (a *Algebra) zero41() {}

// ## Ex 42
func (a *Algebra) sub42(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`добавь слева и справа минус 3.`,
	)
}

func (a *Algebra) shape42(scr *ebiten.Image) {
	a.Label(scr, "a + 3 - 3 = 10 - 3 | - 3", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim42() bool {
	// TODO
	return true
}

func (a *Algebra) xact42() bool { return true }
func (a *Algebra) zero42() {}

// ## Ex 43
func (a *Algebra) sub43(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`a плюс три минус три равно десять минус три.`,
	)
}

func (a *Algebra) shape43(scr *ebiten.Image) {
	a.Label(scr, "a + 3 - 3 = 10 - 3 | - 3", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim43() bool {
	// TODO
	return true
}

func (a *Algebra) xact43() bool { return true }
func (a *Algebra) zero43() {}

// ## Ex 44
func (a *Algebra) sub44(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Теперь считаем.`,
	)
}

func (a *Algebra) shape44(scr *ebiten.Image) {
	a.Label(scr, "a + 3 - 3 = 10 - 3 | - 3", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim44() bool {
	// TODO
	return true
}

func (a *Algebra) xact44() bool { return true }
func (a *Algebra) zero44() {}

// ## Ex 45
func (a *Algebra) sub45(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Слева три минус три будет равно...`,
	)
}

func (a *Algebra) shape45(scr *ebiten.Image) {
	a.Label(scr, "a + 3 - 3 = 10 - 3 | - 3", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim45() bool {
	// TODO
	return true
}

func (a *Algebra) xact45() bool { return true }
func (a *Algebra) zero45() {}

// ## Ex 46
func (a *Algebra) sub46(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`все правильно, 0.`,
	)
}

func (a *Algebra) shape46(scr *ebiten.Image) {
	a.Label(scr, "a + 0 = 10 - 3 | - 3", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim46() bool {
	// TODO
	return true
}

func (a *Algebra) xact46() bool { return true }
func (a *Algebra) zero46() {}

// ## Ex 47
func (a *Algebra) sub47(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`А справа десять минус три...`,
	)
}

func (a *Algebra) shape47(scr *ebiten.Image) {
	a.Label(scr, "a = 10 - 3 | - 3", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim47() bool {
	// TODO
	return true
}

func (a *Algebra) xact47() bool { return true }
func (a *Algebra) zero47() {}

// ## Ex 48
func (a *Algebra) sub48(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`семь.`,
	)
}

func (a *Algebra) shape48(scr *ebiten.Image) {
	a.Label(scr, "a = 10 - 3 | - 3", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim48() bool {
	// TODO
	return true
}

func (a *Algebra) xact48() bool { return true }
func (a *Algebra) zero48() {}

// ## Ex 49
func (a *Algebra) sub49(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Получилось а = 7`,
	)
}

func (a *Algebra) shape49(scr *ebiten.Image) {
	a.Label(scr, "a = 7 | - 3", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim49() bool {
	// TODO
	return true
}

func (a *Algebra) xact49() bool { return true }
func (a *Algebra) zero49() {}

// ## Ex 50
func (a *Algebra) sub50(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Проверка!`,
	)
}

func (a *Algebra) shape50(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim50() bool {
	// TODO
	return true
}

func (a *Algebra) xact50() bool { return true }
func (a *Algebra) zero50() {}

// ## Ex 51
func (a *Algebra) sub51(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Напиши наше первое уравнение, но`,
	)
}

func (a *Algebra) shape51(scr *ebiten.Image) {
	a.Label(scr, "a + 3 = 10", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim51() bool {
	// TODO
	return true
}

func (a *Algebra) xact51() bool { return true }
func (a *Algebra) zero51() {}

// ## Ex 52
func (a *Algebra) sub52(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`подставь вместо а его значение - семь.`,
	)
}

func (a *Algebra) shape52(scr *ebiten.Image) {
	a.Label(scr, "7 + 3 = 10", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim52() bool {
	// TODO
	return true
}

func (a *Algebra) xact52() bool { return true }
func (a *Algebra) zero52() {}

// ## Ex 53
func (a *Algebra) sub53(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Семь плюс три равно десять.`,
	)
}

func (a *Algebra) shape53(scr *ebiten.Image) {
	a.Label(scr, "7 + 3 = 10", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim53() bool {
	// TODO
	return true
}

func (a *Algebra) xact53() bool { return true }
func (a *Algebra) zero53() {}

// ## Ex 54
func (a *Algebra) sub54(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Складываем, семь плюс три - десять,`,
	)
}

func (a *Algebra) shape54(scr *ebiten.Image) {
	a.Label(scr, "10 = 10", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim54() bool {
	// TODO
	return true
}

func (a *Algebra) xact54() bool { return true }
func (a *Algebra) zero54() {}

// ## Ex 55
func (a *Algebra) sub55(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`десять равно десять, все сходится.`,
	)
}

func (a *Algebra) shape55(scr *ebiten.Image) {
	a.Label(scr, "10 = 10", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim55() bool {
	// TODO
	return true
}

func (a *Algebra) xact55() bool { return true }
func (a *Algebra) zero55() {}

// ## Ex 56
func (a *Algebra) sub56(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Решили, поздравляю!`,
	)
}

func (a *Algebra) shape56(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim56() bool {
	// TODO
	return true
}

func (a *Algebra) xact56() bool { return true }
func (a *Algebra) zero56() {}

// ## Ex 57
func (a *Algebra) sub57(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Вторая задача.`,
	)
}

func (a *Algebra) shape57(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim57() bool {
	// TODO
	return true
}

func (a *Algebra) xact57() bool { return true }
func (a *Algebra) zero57() {}

// ## Ex 58
func (a *Algebra) sub58(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Как запишем два арбуза?`,
	)
}

func (a *Algebra) shape58(scr *ebiten.Image) {
	a.Label(scr, "a", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim58() bool {
	// TODO
	return true
}

func (a *Algebra) xact58() bool { return true }
func (a *Algebra) zero58() {}

// ## Ex 59
func (a *Algebra) sub59(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`а умножить на два, верно.`,
	)
}

func (a *Algebra) shape59(scr *ebiten.Image) {
	a.Label(scr, "a×2", 30, -100, 0, clrs.White)
}

func (a *Algebra) anim59() bool {
	// TODO
	return true
}

func (a *Algebra) xact59() bool { return true }
func (a *Algebra) zero59() {}

// ## Ex 60
func (a *Algebra) sub60(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Значит, что запишешь в левой части?`,
	)
}

func (a *Algebra) shape60(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim60() bool {
	// TODO
	return true
}

func (a *Algebra) xact60() bool { return true }
func (a *Algebra) zero60() {}

// ## Ex 61
func (a *Algebra) sub61(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`a умножить на два плюс пять. Да!`,
	)
}

func (a *Algebra) shape61(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim61() bool {
	// TODO
	return true
}

func (a *Algebra) xact61() bool { return true }
func (a *Algebra) zero61() {}

// ## Ex 62
func (a *Algebra) sub62(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`А в правой?`,
	)
}

func (a *Algebra) shape62(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim62() bool {
	// TODO
	return true
}

func (a *Algebra) xact62() bool { return true }
func (a *Algebra) zero62() {}

// ## Ex 63
func (a *Algebra) sub63(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Просто 27, верно.`,
	)
}

func (a *Algebra) shape63(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim63() bool {
	// TODO
	return true
}

func (a *Algebra) xact63() bool { return true }
func (a *Algebra) zero63() {}

// ## Ex 64
func (a *Algebra) sub64(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Что хочешь сделать дальше?`,
	)
}

func (a *Algebra) shape64(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim64() bool {
	// TODO
	return true
}

func (a *Algebra) xact64() bool { return true }
func (a *Algebra) zero64() {}

// ## Ex 65
func (a *Algebra) sub65(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Убрать лишний вес, отлично!`,
	)
}

func (a *Algebra) shape65(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim65() bool {
	// TODO
	return true
}

func (a *Algebra) xact65() bool { return true }
func (a *Algebra) zero65() {}

// ## Ex 66
func (a *Algebra) sub66(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Убираем.`,
	)
}

func (a *Algebra) shape66(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim66() bool {
	// TODO
	return true
}

func (a *Algebra) xact66() bool { return true }
func (a *Algebra) zero66() {}

// ## Ex 67
func (a *Algebra) sub67(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Проводи вертикальную черту`,
	)
}

func (a *Algebra) shape67(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim67() bool {
	// TODO
	return true
}

func (a *Algebra) xact67() bool { return true }
func (a *Algebra) zero67() {}

// ## Ex 68
func (a *Algebra) sub68(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`и пиши, какой вес убираешь.`,
	)
}

func (a *Algebra) shape68(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim68() bool {
	// TODO
	return true
}

func (a *Algebra) xact68() bool { return true }
func (a *Algebra) zero68() {}

// ## Ex 69
func (a *Algebra) sub69(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Минус пять.`,
	)
}

func (a *Algebra) shape69(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim69() bool {
	// TODO
	return true
}

func (a *Algebra) xact69() bool { return true }
func (a *Algebra) zero69() {}

// ## Ex 70
func (a *Algebra) sub70(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Получается а умножить на два плюс пять минус пять`,	
	)
}

func (a *Algebra) shape70(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim70() bool {
	// TODO
	return true
}

func (a *Algebra) xact70() bool { return true }
func (a *Algebra) zero70() {}

// ## Ex 71
func (a *Algebra) sub71(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`равно двадцать семь минус пять.`,
	)
}

func (a *Algebra) shape71(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim71() bool {
	// TODO
	return true
}

func (a *Algebra) xact71() bool { return true }
func (a *Algebra) zero71() {}

// ## Ex 72
func (a *Algebra) sub72(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Отлично! Считаем.`,
	)
}

func (a *Algebra) shape72(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim72() bool {
	// TODO
	return true
}

func (a *Algebra) xact72() bool { return true }
func (a *Algebra) zero72() {}

// ## Ex 73
func (a *Algebra) sub73(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Пять минус пять - ноль,`,
	)
}

func (a *Algebra) shape73(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim73() bool {
	// TODO
	return true
}

func (a *Algebra) xact73() bool { return true }
func (a *Algebra) zero73() {}

// ## Ex 74
func (a *Algebra) sub74(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`двадцать семь минус пять - двадцать два.`,
	)
}

func (a *Algebra) shape74(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim74() bool {
	// TODO
	return true
}

func (a *Algebra) xact74() bool { return true }
func (a *Algebra) zero74() {}

// ## Ex 75
func (a *Algebra) sub75(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Что остается?`,
	)
}

func (a *Algebra) shape75(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim75() bool {
	// TODO
	return true
}

func (a *Algebra) xact75() bool { return true }
func (a *Algebra) zero75() {}

// ## Ex 76
func (a *Algebra) sub76(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`а умножить на два равно двадцать два.`,
	)
}

func (a *Algebra) shape76(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim76() bool {
	// TODO
	return true
}

func (a *Algebra) xact76() bool { return true }
func (a *Algebra) zero76() {}

// ## Ex 77
func (a *Algebra) sub77(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Мы нашли вес двух арбузов, ура!`,
	)
}

func (a *Algebra) shape77(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim77() bool {
	// TODO
	return true
}

func (a *Algebra) xact77() bool { return true }
func (a *Algebra) zero77() {}

// ## Ex 78
func (a *Algebra) sub78(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Но как из двух арбузов сделать один?`,
	)
}

func (a *Algebra) shape78(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim78() bool {
	// TODO
	return true
}

func (a *Algebra) xact78() bool { return true }
func (a *Algebra) zero78() {}

// ## Ex 79
func (a *Algebra) sub79(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Разделить на два, конечно!`,
	)
}

func (a *Algebra) shape79(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim79() bool {
	// TODO
	return true
}

func (a *Algebra) xact79() bool { return true }
func (a *Algebra) zero79() {}

// ## Ex 80
func (a *Algebra) sub80(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Теперь смотри.`,
	)
}

func (a *Algebra) shape80(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim80() bool {
	// TODO
	return true
}

func (a *Algebra) xact80() bool { return true }
func (a *Algebra) zero80() {}

// ## Ex 81
func (a *Algebra) sub81(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Вес на обеих чашах одинаковый, так?`,
	)
}

func (a *Algebra) shape81(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim81() bool {
	// TODO
	return true
}

func (a *Algebra) xact81() bool { return true }
func (a *Algebra) zero81() {}

// ## Ex 82
func (a *Algebra) sub82(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Значит, если мы одновременно с обеих чаш`,
	)
}

func (a *Algebra) shape82(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim82() bool {
	// TODO
	return true
}

func (a *Algebra) xact82() bool { return true }
func (a *Algebra) zero82() {}

// ## Ex 83
func (a *Algebra) sub83(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`уберем одинаковую часть (половину),`,
	)
}

func (a *Algebra) shape83(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim83() bool {
	// TODO
	return true
}

func (a *Algebra) xact83() bool { return true }
func (a *Algebra) zero83() {}

// ## Ex 84
func (a *Algebra) sub84(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`что будет с равновесием?`,
	)
}

func (a *Algebra) shape84(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim84() bool {
	// TODO
	return true
}

func (a *Algebra) xact84() bool { return true }
func (a *Algebra) zero84() {}

// ## Ex 85
func (a *Algebra) sub85(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Сохранится!`,
	)
}

func (a *Algebra) shape85(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim85() bool {
	// TODO
	return true
}

func (a *Algebra) xact85() bool { return true }
func (a *Algebra) zero85() {}

// ## Ex 86
func (a *Algebra) sub86(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Запишем это все.`,
	)
}

func (a *Algebra) shape86(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim86() bool {
	// TODO
	return true
}

func (a *Algebra) xact86() bool { return true }
func (a *Algebra) zero86() {}

// ## Ex 87
func (a *Algebra) sub87(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Снова вертикальная черта.`,
	)
}

func (a *Algebra) shape87(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim87() bool {
	// TODO
	return true
}

func (a *Algebra) xact87() bool { return true }
func (a *Algebra) zero87() {}

// ## Ex 88
func (a *Algebra) sub88(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Что напишешь?`,
	)
}

func (a *Algebra) shape88(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim88() bool {
	// TODO
	return true
}

func (a *Algebra) xact88() bool { return true }
func (a *Algebra) zero88() {}

// ## Ex 89
func (a *Algebra) sub89(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Разделить на 2.`,
	)
}

func (a *Algebra) shape89(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim89() bool {
	// TODO
	return true
}

func (a *Algebra) xact89() bool { return true }
func (a *Algebra) zero89() {}

// ## Ex 90
func (a *Algebra) sub90(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Снова перепиши уравнение с новым действием.`,
	)
}

func (a *Algebra) shape90(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim90() bool {
	// TODO
	return true
}

func (a *Algebra) xact90() bool { return true }
func (a *Algebra) zero90() {}

// ## Ex 91
func (a *Algebra) sub91(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Что получается?`,
	)
}

func (a *Algebra) shape91(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim91() bool {
	// TODO
	return true
}

func (a *Algebra) xact91() bool { return true }
func (a *Algebra) zero91() {}

// ## Ex 92
func (a *Algebra) sub92(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`а умножить на два разделить на два`,
	)
}

func (a *Algebra) shape92(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim92() bool {
	// TODO
	return true
}

func (a *Algebra) xact92() bool { return true }
func (a *Algebra) zero92() {}

// ## Ex 93
func (a *Algebra) sub93(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`равно двадцать два разделить на два.`,
	)
}

func (a *Algebra) shape93(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim93() bool {
	// TODO
	return true
}

func (a *Algebra) xact93() bool { return true }
func (a *Algebra) zero93() {}

// ## Ex 94
func (a *Algebra) sub94(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Считаем.`,
	)
}

func (a *Algebra) shape94(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim94() bool {
	// TODO
	return true
}

func (a *Algebra) xact94() bool { return true }
func (a *Algebra) zero94() {}

// ## Ex 95
func (a *Algebra) sub95(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`а умножить на два да разделить на два это просто а,`,
	)
}

func (a *Algebra) shape95(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim95() bool {
	// TODO
	return true
}

func (a *Algebra) xact95() bool { return true }
func (a *Algebra) zero95() {}

// ## Ex 96
func (a *Algebra) sub96(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`двадцать два разделить на два - одиннадцать,`,
	)
}

func (a *Algebra) shape96(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim96() bool {
	// TODO
	return true
}

func (a *Algebra) xact96() bool { return true }
func (a *Algebra) zero96() {}

// ## Ex 97
func (a *Algebra) sub97(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`а равно одиннадцать.`,
	)
}

func (a *Algebra) shape97(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim97() bool {
	// TODO
	return true
}

func (a *Algebra) xact97() bool { return true }
func (a *Algebra) zero97() {}

// ## Ex 98
func (a *Algebra) sub98(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Проверка!`,
	)
}

func (a *Algebra) shape98(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim98() bool {
	// TODO
	return true
}

func (a *Algebra) xact98() bool { return true }
func (a *Algebra) zero98() {}

// ## Ex 99
func (a *Algebra) sub99(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Перепиши первое (в этой задаче) уравнение, но`,	
	)
}

func (a *Algebra) shape99(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim99() bool {
	// TODO
	return true
}

func (a *Algebra) xact99() bool { return true }
func (a *Algebra) zero99() {}

// ## Ex 100
func (a *Algebra) sub100(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`подставь вместо а его значение.`,
	)
}

func (a *Algebra) shape100(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim100() bool {
	// TODO
	return true
}

func (a *Algebra) xact100() bool { return true }
func (a *Algebra) zero100() {}

// ## Ex 101
func (a *Algebra) sub101(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Одиннадцать умножить на два плюс пять`,
	)
}

func (a *Algebra) shape101(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim101() bool {
	// TODO
	return true
}

func (a *Algebra) xact101() bool { return true }
func (a *Algebra) zero101() {}

// ## Ex 102
func (a *Algebra) sub102(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`равно двадцать семь.`,
	)
}

func (a *Algebra) shape102(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim102() bool {
	// TODO
	return true
}

func (a *Algebra) xact102() bool { return true }
func (a *Algebra) zero102() {}

// ## Ex 103
func (a *Algebra) sub103(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Двадцать два плюс пять равно двадцать семь.`,
	)
}

func (a *Algebra) shape103(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim103() bool {
	// TODO
	return true
}

func (a *Algebra) xact103() bool { return true }
func (a *Algebra) zero103() {}

// ## Ex 104
func (a *Algebra) sub104(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Двадцать семь равно двадцать семь.`,
	)
}

func (a *Algebra) shape104(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim104() bool {
	// TODO
	return true
}

func (a *Algebra) xact104() bool { return true }
func (a *Algebra) zero104() {}

// ## Ex 105
func (a *Algebra) sub105(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Сходится!`,
	)
}

func (a *Algebra) shape105(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim105() bool {
	// TODO
	return true
}

func (a *Algebra) xact105() bool { return true }
func (a *Algebra) zero105() {}

// ## Ex 106
func (a *Algebra) sub106(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Ура, решили!`,
	)
}

func (a *Algebra) shape106(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim106() bool {
	// TODO
	return true
}

func (a *Algebra) xact106() bool { return true }
func (a *Algebra) zero106() {}

// ## Ex 107
func (a *Algebra) sub107(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Итак, мы увидели, что можно делать с уравнением:`,
	)
}

func (a *Algebra) shape107(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim107() bool {
	// TODO
	return true
}

func (a *Algebra) xact107() bool { return true }
func (a *Algebra) zero107() {}

// ## Ex 108
func (a *Algebra) sub108(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`вычитать из обеих частей одинаковое число,`,
	)
}

func (a *Algebra) shape108(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim108() bool {
	// TODO
	return true
}

func (a *Algebra) xact108() bool { return true }
func (a *Algebra) zero108() {}

// ## Ex 109
func (a *Algebra) sub109(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`делить обе части на одно и же число,`,
	)
}

func (a *Algebra) shape109(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim109() bool {
	// TODO
	return true
}

func (a *Algebra) xact109() bool { return true }
func (a *Algebra) zero109() {}

// ## Ex 110
func (a *Algebra) sub110(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`равенство сохраняется.`,
	)
}

func (a *Algebra) shape110(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim110() bool {
	// TODO
	return true
}

func (a *Algebra) xact110() bool { return true }
func (a *Algebra) zero110() {}

// ## Ex 111
func (a *Algebra) sub111(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`А как думаешь, прибавлять одинаковое число можно?`,
	)
}

func (a *Algebra) shape111(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim111() bool {
	// TODO
	return true
}

func (a *Algebra) xact111() bool { return true }
func (a *Algebra) zero111() {}

// ## Ex 112
func (a *Algebra) sub112(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`А умножать?`,
	)
}

func (a *Algebra) shape112(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim112() bool {
	// TODO
	return true
}

func (a *Algebra) xact112() bool { return true }
func (a *Algebra) zero112() {}

// ## Ex 113
func (a *Algebra) sub113(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Конечно да!`,
	)
}

func (a *Algebra) shape113(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim113() bool {
	// TODO
	return true
}

func (a *Algebra) xact113() bool { return true }
func (a *Algebra) zero113() {}

// ## Ex 114
func (a *Algebra) sub114(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Но самое главное - мы увидели, что можно`,
	)
}

func (a *Algebra) shape114(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim114() bool {
	// TODO
	return true
}

func (a *Algebra) xact114() bool { return true }
func (a *Algebra) zero114() {}

// ## Ex 115
func (a *Algebra) sub115(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`записать любое число буквой и с этим работать.`,
	)
}

func (a *Algebra) shape115(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim115() bool {
	// TODO
	return true
}

func (a *Algebra) xact115() bool { return true }
func (a *Algebra) zero115() {}

// ## Ex 116
func (a *Algebra) sub116(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Вот когда числа обозначают буквами - это и есть алгебра!`,
	)
}

func (a *Algebra) shape116(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim116() bool {
	// TODO
	return true
}

func (a *Algebra) xact116() bool { return true }
func (a *Algebra) zero116() {}

// ## Ex 117
func (a *Algebra) sub117(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Не так уж и сложно, правда?`,
	)
}

func (a *Algebra) shape117(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim117() bool {
	// TODO
	return true
}

func (a *Algebra) xact117() bool { return true }
func (a *Algebra) zero117() {}

// ## Ex 118
func (a *Algebra) sub118(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Поздравляю!`,
	)
}

func (a *Algebra) shape118(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim118() bool {
	// TODO
	return true
}

func (a *Algebra) xact118() bool { return true }
func (a *Algebra) zero118() {}

// ## Ex 119
func (a *Algebra) sub119(*ebiten.Image) func(*ebiten.Image) {
	return a.Sub(
		`Дальше узнаем, как это все можно упростить.`,
	)
}

func (a *Algebra) shape119(scr *ebiten.Image) {
	// TODO
}

func (a *Algebra) anim119() bool {
	// TODO
	return true
}

func (a *Algebra) xact119() bool { return true }
func (a *Algebra) zero119() {}
