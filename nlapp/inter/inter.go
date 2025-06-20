package inter

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Inter struct {
	ids []ebiten.TouchID
}

/*func Init(x1, y1, x2, y2 float32) *Inter {
	in := &Inter{}
	return in
}//*/

func (in Inter) Escape() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyEscape)
}

func (in Inter) Enter() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyEnter) ||
	       inpututil.IsKeyJustPressed(ebiten.KeyNumpadEnter)
}

func (in Inter) Backspace() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyBackspace)
}

func (in Inter) Space() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeySpace)
}

func (in Inter) ArrowUp() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyArrowUp)
}

func (in Inter) ArrowDown() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyArrowDown)
}

func (in Inter) ArrowLeft() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft)
}

func (in Inter) ArrowRight() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyArrowRight)
}

func (in Inter) MouseL() bool {
	return inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
}

func (in Inter) MouseR() bool {
	return inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight)
}

func (in Inter) MousePos() (x, y int) {
	return ebiten.CursorPosition()
}

// TODO: check
func (in Inter) MouseLIn(x1, y1, x2, y2 float32) bool {
	a, b := in.MousePos()
	x, y := float32(a), float32(b)
	/*x *= x
	y *= y
	x1 *= x1
	y1 *= y1
	x2 *= x2
	y2 *= y2//*/
	if in.MouseL() {
		return x1 < x && x < x2 && y1 < y && y < y2 ||
		       x1 > x && x > x2 && y1 > y && y > y2
		//return x1*x1 < x*x && x*x < x2*x2 && y1*y1 < y*y && y*y < y2*y2
	}
	return false
}

func (in Inter) TouchIn(x1, y1, x2, y2 float32) bool {
	in.ids = inpututil.AppendJustPressedTouchIDs(in.ids[:0])
	for _, id := range in.ids { // todo use in.ids[0] instead?
		a, b := ebiten.TouchPosition(id)
		x, y := float32(a), float32(b)
		if x1 < x && x < x2 && y1 < y && y < y2 ||
       x1 > x && x > x2 && y1 > y && y > y2 {
				return true
		}
	}
	return false
}

func (in Inter) TapIn(x1, y1, x2, y2 float32) bool {
	return in.TouchIn(x1, y1, x2, y2) || in.MouseLIn(x1, y1, x2, y2)
}

func (in Inter) Number() (int, bool) {
	if inpututil.IsKeyJustPressed(ebiten.KeyDigit0) {return 0, true}
	if inpututil.IsKeyJustPressed(ebiten.KeyDigit1) {return 1, true}
	if inpututil.IsKeyJustPressed(ebiten.KeyDigit2) {return 2, true}
	if inpututil.IsKeyJustPressed(ebiten.KeyDigit3) {return 3, true}
	if inpututil.IsKeyJustPressed(ebiten.KeyDigit4) {return 4, true}
	if inpututil.IsKeyJustPressed(ebiten.KeyDigit5) {return 5, true}
	if inpututil.IsKeyJustPressed(ebiten.KeyDigit6) {return 6, true}
	if inpututil.IsKeyJustPressed(ebiten.KeyDigit7) {return 7, true}
	if inpututil.IsKeyJustPressed(ebiten.KeyDigit8) {return 8, true}
	if inpututil.IsKeyJustPressed(ebiten.KeyDigit9) {return 9, true}

	if inpututil.IsKeyJustPressed(ebiten.KeyNumpad0) {return 0, true}
	if inpututil.IsKeyJustPressed(ebiten.KeyNumpad1) {return 1, true}
	if inpututil.IsKeyJustPressed(ebiten.KeyNumpad2) {return 2, true}
	if inpututil.IsKeyJustPressed(ebiten.KeyNumpad3) {return 3, true}
	if inpututil.IsKeyJustPressed(ebiten.KeyNumpad4) {return 4, true}
	if inpututil.IsKeyJustPressed(ebiten.KeyNumpad5) {return 5, true}
	if inpututil.IsKeyJustPressed(ebiten.KeyNumpad6) {return 6, true}
	if inpututil.IsKeyJustPressed(ebiten.KeyNumpad7) {return 7, true}
	if inpututil.IsKeyJustPressed(ebiten.KeyNumpad8) {return 8, true}
	if inpututil.IsKeyJustPressed(ebiten.KeyNumpad9) {return 9, true}

	return 0, false
}

//func (in *Inter) NumPad() {}
