package frame

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func Escape() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) { os.Exit(0) }
}

func Space() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeySpace)
}

func Enter() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyEnter) ||
	       inpututil.IsKeyJustPressed(ebiten.KeyNumpadEnter)
}

func MouseL() bool {
	return ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
}

func MouseR() bool {
	return ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight)
}

func Number() (int, bool) {
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

