package frame
// TODO move to graph??
import (
	//"errors"
	"log"
	"math"
	//"slices"
)

func Polygon(n int, x, y, rad, ang float32) []float32 {
	if n < 3 {
		log.Println("frame.Polygon: less than three sides")
		return []float32{}
	}

	a := 2*math.Pi / float64(n)
	l := make([]float32, 2*n+2)
	for i := 0; i < 2*n+2; i += 2 {
		s, c := math.Sincos(float64(i/2)*a + float64(ang))
		l[i] = rad * float32(c) + x
		l[i+1] = rad * float32(s) + y
	}
	return l
}

func AnimLine(x1, y1, x2, y2 float32) (c, s float32) {
	dx, dy := x2-x1, y2-y1
	h := float32(math.Hypot(float64(dx), float64(dy)))
	if h == 0 { return 0, 0 }
	c, s = dx/h, dy/h
	return
}

func AnimPoly(crds []float32) func(float32) []float32 {
	if len(crds) < 4 { log.Fatal("frame.AnimPoly: not enough coordinates") }
	i := 0
	x1, y1, x2, y2 := crds[i], crds[i+1], crds[i+2], crds[i+3]
	c, s := AnimLine(x1, y1, x2, y2)
	var dx, dy float32 = 0, 0
	ll := []float32{x1, y1, x1, y1}

	return func(speed float32) []float32 {
		if speed < 0 { // todo what is this??
			i, dx, dy = 0, 0, 0
			ll = []float32{crds[i], crds[i+1], crds[i], crds[i+1]}
			return ll
		}

		if i > len(crds)-4 { return crds }
		x1, y1, x2, y2 = crds[i], crds[i+1], crds[i+2], crds[i+3]
		c, s = AnimLine(x1, y1, x2, y2)
		dx += c * speed
		dy += s * speed
		x, y := dx + x1, dy + y1
		ll[i+2], ll[i+3] = x, y

		if !(x1 < x2 && x < x2 ||
	    	 x1 > x2 && x > x2 ||
	    	 y1 < y2 && y < y2 ||
	    	 y1 > y2 && y > y2) {

			   	ll[i+2], ll[i+3] = x2, y2
			   	dx, dy = 0, 0
			   	ll = append(ll, x2, y2)
			   	i += 2
		}
		return ll
	}
}
