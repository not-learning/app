package frame

import (
	//"errors"
	"log"
	"math"
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
	c, s := AnimLine(crds[i], crds[i+1], crds[i+2], crds[i+3])
	var dx, dy float32 = 0, 0
	ll := []float32{
		crds[i], crds[i+1],
		crds[i], crds[i+1],
	}
	return func(speed float32) []float32 {
		if i > len(crds)-4 { return crds }
		c, s = AnimLine(crds[i], crds[i+1], crds[i+2], crds[i+3])
		dx += c * speed
		dy += s * speed
		ll[i+2], ll[i+3] = dx + crds[i], dy + crds[i+1]
		if dx*dx >= (crds[i+2]-crds[i])*(crds[i+2]-crds[i]) &&
		   dy*dy >= (crds[i+3]-crds[i+1])*(crds[i+3]-crds[i+1]) {
		   	ll[i+2], ll[i+3] = crds[i+2], crds[i+3]
		   	dx, dy = 0, 0
		   	ll = append(ll, crds[i+2], crds[i+3])
		   	i += 2
		}
		return ll
	}
}

/*func XLinesPA(x1, y1, x2, y2, a, r float32) (x, y float32, err error) {
	var x3, y3 float32 = 0, 0
	s, c := math.Sincos(float64(a))
	x4, y4 := r*float32(c), r*float32(s)

	d := (x1-x2)*(y3-y4) - (y1-y2)*(x3-x4)
	if d == 0 {
		return 0, 0, errors.New("XLinesPA: lines are parallel or coincide")
	}

	t := ((x1-x3)*(y3-y4) - (y1-y3)*(x3-x4)) / d
	if t < 0 || t > 1 {
		return 0, 0, errors.New("XLinesPA: cross point out of bounds")
	}
	
	u := ((x1-x2)*(y1-y3) - (y1-y2)*(x1-x3)) / d
	if u < 0 || u > 1 {
		return 0, 0, errors.New("XLinesPA: cross point out of bounds")
	}

	x = x1 + t*(x2-x1)
	y = y1 + t*(y2-y1)
	return x, y, nil
}//*/
