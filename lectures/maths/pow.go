package maths

import (
	//c "github.com/not-learning/app/clrs"
	"github.com/not-learning/app/frame"
	//"github.com/not-learning/app/tracks"
	//_ "embed"
	"strings"
	"encoding/csv"

	"github.com/hajimehoshi/ebiten/v2"
)

type Power struct {
	*frame.Lecture
	//tracks *tracks.Tracks
}

/*//go:embed powScript.txt
var script string

////go:embed powText.txt
var text1 = `5×2 = 5+5`

////go:embed powProblem.txt//*/

// text|ans
var ex0 = `|`
var ex1 = `|
5×3 = 5+5+5|`

var ex2 = `|
5² = 5×5|
5³ = 5×5×5|
5⁴ = 5×5×5×5|
5⁵ = 5×5×5×5×5|`

var ex3 = `
10² = |100
10³ = |1000
10⁴ = |10000
10ˣ = 1000000|
x = |6`

var t1 = `Ты конечно знаешь, что умножение - это сложение одинаковых чисел.
Пять на три - это пять плюс пять плюс пять.
Так вот, степень - это просто умножение одинаковых чисел.
Пять во второй это пять умножить на пять.
Пять в третьей это пять на пять на пять.
И так
далее.
Сможешь найти десять во второй?
А в третьей?
В четвертой?
В какую степень надо возвести десять, чтобы получить миллион?`

func myRead(text string) [][]string {
	r := csv.NewReader(strings.NewReader(text))
	r.Comma = '|'
	recs, err := r.ReadAll()
	if err != nil { panic(err) }
	return recs
}

func InitPow(x1, y1, x2, y2 float32) *Power {
	p := &Power{}
	p.Lecture = frame.Init(x1, y1, x2, y2)

	p.Lecture.ExInit(myRead(ex0), myRead(ex1), myRead(ex2), myRead(ex3))
	p.Lecture.ScriptInit(strings.Split(t1, "\n"))
	//p.Lecture.Play()

	//p.tracks = tracks.Init()
	//p.tracks.Play()
	return p
}

func (p *Power) Draw(scr *ebiten.Image) {
	p.Lecture.Draw(scr)
}

func (p *Power) Update() {
	//p.tracks.Update()
	p.Lecture.Update()
}

/*
hi there
ˣ
⁰¹²³⁴⁵⁶⁷⁸⁹
⁺⁻⁼⁽⁾ⁿ
𝟝
⅕
5×5×5
*/
