package fonts

import (
	"bufio"
	"embed"
	"image/color"
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// //go:embed DejaVuSans.ttf
//go:embed JuliaMono-Light.ttf
var ff embed.FS

type Font struct {
	face *text.GoTextFace
	op *text.DrawOptions

	scale float64
	//SizeL, SizeM, SizeS float64
}

func loadFont() *text.GoTextFaceSource {
	var face *text.GoTextFaceSource
	if face != nil { return face }

	file, err := ff.Open("JuliaMono-Light.ttf")
	//file, err := ff.Open("DejaVuSans.ttf")
	if err != nil { log.Fatal(err) }
	defer file.Close()//*/

	face, err = text.NewGoTextFaceSource(bufio.NewReader(file))
	if err != nil { log.Fatal(err) }
	return face
}

func InitFont(scale float32) *Font {
	f := &Font{ face: &text.GoTextFace{} }
	f.scale = float64(scale)
	f.face.Source = loadFont()
	return f
}

func (f *Font) Update(scale float32) {
	f.scale = float64(scale)
}

//func (f *Font) Set(size float64, clr color.Color) {}

func (f *Font) TextSize(str string) (w, h float32) {
	w = float32(text.Advance(str, f.face))
	m := f.face.Metrics()
	h = float32(m.HAscent + m.HDescent)
	return
}

func (f *Font) strLen(size float64, str string) float64 {
	f.op = &text.DrawOptions{}
	f.face.Size = size * f.scale // TODO proper size
	return text.Advance(str, f.face)
}

func (f *Font) Wrap(size, width float32, str string) []string {
	ss, ww := float64(size), float64(width)
	words := strings.Fields(str)
	if str == "" { return []string{str} }
	res := []string{words[0]}
	rem := ww - f.strLen(ss, res[0])
	i := 0
	for _, w := range words[1:] {
		if f.strLen(ss, w) > rem {
			res = append(res, w)
			i++
			rem = ww - f.strLen(ss, w)
		} else {
			res[i] += " " + w
			rem = rem - f.strLen(ss, " "+w)
		}
	}
	return res
}

/*func (f *Font) Super(str, super string) {
	w, h := f.TextSize(str)
}//*/

func (f *Font) Draw(scr *ebiten.Image, str string, size, x, y float32, clr color.Color) {
	f.op = &text.DrawOptions{}
	f.op.ColorScale.ScaleWithColor(clr)
	f.face.Size = float64(size) * f.scale
	f.op.GeoM.Translate(float64(x), float64(y))
	text.Draw(scr, str, f.face, f.op)
}

func (f *Font) DrawCenter(scr *ebiten.Image, str string, size, x, y float32, clr color.Color) {
	f.op = &text.DrawOptions{}
	f.op.ColorScale.ScaleWithColor(clr)
	f.face.Size = float64(size) * f.scale
	f.op.GeoM.Translate(float64(x), float64(y))
	f.op.PrimaryAlign = text.AlignCenter
	f.op.SecondaryAlign = text.AlignCenter
	text.Draw(scr, str, f.face, f.op)
}
