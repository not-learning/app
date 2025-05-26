package mobile

import (
	"github.com/hajimehoshi/ebiten/v2/mobile"

	"github.com/not-learning/app/nlapp/game"
)

func init() {
	g := game.Init()
	mobile.SetGame(g)
}

// Dummy is a dummy exported function.
//
// gomobile doesn't compile a package that doesn't include any exported function.
// Dummy forces gomobile to compile this package.
func Dummy() {}
