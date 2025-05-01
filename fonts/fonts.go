package fonts

import (
	"bufio"
	"image/color"
	"log"
	"os"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// TODO: go:embed the font

type Font struct {
	face *text.GoTextFace
	//size float64
	op *text.DrawOptions
}

func loadFont() *text.GoTextFaceSource {
	var face *text.GoTextFaceSource
	if face != nil { return face }

	//file, err := os.Open("assets/fonts/JuliaMono-Light.ttf")
	file, err := os.Open("assets/fonts/DejaVuSans.ttf")
	if err != nil { log.Fatal(err) }
	defer file.Close()

	face, err = text.NewGoTextFaceSource(bufio.NewReader(file))
	if err != nil { log.Fatal(err) }
	return face
}

func InitFont() *Font {
	f := &Font{face: &text.GoTextFace{}}
	f.face.Source = loadFont()
	return f
} //*/

func (f *Font) Set(size float64, clr color.Color) {
	f.op = &text.DrawOptions{}
	f.op.ColorScale.ScaleWithColor(clr)
	f.face.Size = size
}

func (f *Font) TextSize(str string) (w, h float64) {
	w = text.Advance(str, f.face)
	m := f.face.Metrics()
	h = m.HAscent + m.HDescent
	return
}

func (f *Font) strLen(str string) float64 {
	return text.Advance(str, f.face)
}

func (f *Font) Wrap(str string, width float64) []string {
	words := strings.Fields(str)
	if str == "" { return []string{str} }
	res := []string{words[0]}
	rem := width - f.strLen(res[0])
	i := 0
	for _, w := range words[1:] {
		if f.strLen(w) > rem {
			res = append(res, w)
			i++
			rem = width - f.strLen(w)
		} else {
			res[i] += " " + w
			rem = rem - f.strLen(" " + w)
		}
	}
	return res
}

/*func (f *Font) Super(str, super string) {
	w, h := f.TextSize(str)
}//*/

func (f *Font) Draw(scr *ebiten.Image, str string, x, y float32) {
	f.op.GeoM.Translate(float64(x), float64(y))
	text.Draw(scr, str, f.face, f.op)
}

func (f *Font) DrawCenter(scr *ebiten.Image, str string, x, y float32) {
	f.op.GeoM.Translate(float64(x), float64(y))
	f.op.PrimaryAlign = text.AlignCenter
	f.op.SecondaryAlign = text.AlignCenter
	text.Draw(scr, str, f.face, f.op)
}
