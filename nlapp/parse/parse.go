package parse

import (//"fmt"
	_ "embed"
	"log"
	"strconv"
	"strings"
)

// TODO colours

type Label struct {
	isEx bool
	str  string
	size, x, y float32
}

type Circle struct {
	x, y, r float32
}

type Poly struct {
	crds []float32
}

type Shape struct {
	isEx bool
	circles []Circle
	polys   []Poly
	labels  []Label
}

type Anims struct {
	startValues map[string]float32
	endValues map[string]float32
	speeds map[string]float32
}

type Xacts struct {
	isEx  bool
	ansF  float32
	isNum bool
	ansS  string
}

type Things struct {
	sub string
	Shape
	Anims
	Xacts
	//Zero
}

//var isEx = false

var curValues = make(map[string]float32)

//func IsEx() bool { return isEx }

func (th Things) Sub() string { return th.sub }

func (th Things) Shapes() Shape { return th.Shape }

func (sh Shape) Labels() []Label { return sh.labels }

func (sh Shape) IsEx() bool { return sh.isEx }

func (x Xacts) IsEx() bool { return x.isEx }

func (lb Label) IsEx() bool { return lb.isEx }

func (lb Label) Label() (str string, size, x, y float32) {
	return lb.str, lb.size, lb.x, lb.y
}

func (sh Shape) Circles() []Circle { return sh.circles }

func (c Circle) Circle() (cx, cy, cr float32) {
	if x, ok := prs("x"); ok { c.x = x }
	if y, ok := prs("y"); ok { c.y = y }
	if r, ok := prs("r"); ok { c.r = r }
	return c.x, c.y, c.r
}

func (sh Shape) Polys() []Poly { return sh.polys }

func (pl Poly) Poly() []float32 {
	return pl.crds
}

//func (th Things) Anims() Anims { return th.Anims }

func (an Anims) Anim() bool {
	if an.endValues == nil { // TODO not good
		return true
	}
	for k, v := range an.endValues {
		if curValues[k] < v {
			curValues[k] += an.speeds[k]
		} else {
			curValues[k] = v
			return true
		}
	}
	return false
}

func (th Things) Xact() Xacts { return th.Xacts }

func (x Xacts) IsNum() bool { return x.isNum }

func (x Xacts) AnsF() float32 { return x.ansF }

func (x Xacts) AnsS() string { return x.ansS }

func (an Anims) Zero() { // TODO BAD
	for k, v := range an.startValues {
		curValues[k] = v
	}
}

func indices(str string, chars string) []int {
	i := strings.IndexAny(str, chars)
	ii := []int{}
	k := 0
	for i >= 0 {
		ii = append(ii, i)
		k = i + 1
		i = strings.IndexAny(str[k:], chars)
		if i >= 0 { i += k }
	}
	return ii
}

func prs(str string) (float32, bool) {
	num, ok := curValues[str]
	if ok { return num, ok }

	n, err := strconv.ParseFloat(str, 32)
	if err != nil {
		//log.Println("parse.prs:", err) DEV
	} else {
		ok = true
	}
	num = float32(n)
	return num, ok
}

func Do(raw string) []Things {
	log.SetFlags(log.Lshortfile)
	tt := []Things{}
	ii := indices(raw, "[]")
	i, k := 0, 0
	for i < len(ii) {
		if raw[ii[i+1]] != ']' {
			log.Println("parse.Do: brackets mismatch")
		}
		switch raw[ii[i]-1] {
		case 't':
			tt = append(tt, Things{sub: raw[ii[i]+1:ii[i+1]]})
			k++
		case 'S':
			tt[k-1].Shape = doShapes(raw[ii[i]+1:ii[i+1]])
		case 'A':
			tt[k-1].Anims = doAnims(raw[ii[i]+1:ii[i+1]])
		case 'X':
			tt[k-1].Xacts = doXact(raw[ii[i]+1:ii[i+1]])
		}
		i += 2
	}
	return tt
}

func doShapes(str string) Shape {
	isEx := false
	str = strings.TrimSpace(str)
	if strings.Contains(str, "?") {
		isEx = true
	} else {
		isEx = false
	}
	ss := strings.Split(str, "\n")
	pp, cc, ll := []Poly{}, []Circle{}, []Label{}
	for _, s := range ss {
		s = strings.TrimSpace(s)
		switch s[0] {
		case 'P':
			pp = append(pp, doPoly(s))
		case 'C':
			cc = append(cc, doCircle(s))
		case 'L':
			ll = append(ll, doLabel(s))
		}
	}
	posLabels(ll)
	return Shape{isEx: isEx, polys: pp, circles: cc, labels: ll}
}

func doAnims(str string) Anims {
	an := Anims{
		startValues: make(map[string]float32),
		endValues: make(map[string]float32),
		speeds: make(map[string]float32),
	}
	str = strings.TrimSpace(str)
	ss := strings.Split(str, ",")
	for _, s := range ss {
		s = strings.TrimSpace(s)
		c := string(s[0])
		s = strings.TrimSpace(s[1:])
		strsN := strings.Split(s, "...")
		n1, err := strconv.ParseFloat(strsN[0], 32)
		if err != nil { log.Println("parse.doAnims:", err)}
		n2, err := strconv.ParseFloat(strsN[1], 32)
		if err != nil { log.Println("parse.doAnims:", err)}

		an.startValues[c] = float32(n1)
		curValues[c] = float32(n1)
		an.endValues[c] = float32(n2)
		an.speeds[c] = (an.endValues[c] - an.startValues[c]) / 60
	}
	return an
}

func doCircle(str string) Circle {
	var x, y, r float32
	str = strings.TrimSpace(strings.Trim(str, "C"))
	ss := strings.Split(str, ",")
	for _, s := range ss {
		s = strings.TrimSpace(s)
		switch s[0] {
		case 'x':
			tt := strings.TrimSpace(strings.TrimPrefix(s, "x"))
			x, _ = prs(tt)
		case 'y':
			tt := strings.TrimSpace(strings.TrimPrefix(s, "y"))
			y, _ = prs(tt)
		case 'r':
			tt := strings.TrimSpace(strings.TrimPrefix(s, "r"))
			r, _ = prs(tt)
		}
	}
	return Circle{x: x, y: y, r: r}
}

func doPoly(str string) Poly {
	var (
		n float64
		crds []float32
		err error
	)
	str = strings.TrimSpace(strings.Trim(str, "P"))
	ss := strings.Fields(str)
	crds = make([]float32, len(ss))
	for i, s := range ss {
		n, err = strconv.ParseFloat(s, 32)
		if err != nil { log.Print("parse.doPoly:", err)}
		crds[i] = float32(n)
	}
	return Poly{crds: crds}
}

func doLabel(text string) Label {
	var (
		str string
		size, x, y float64
		err error
	)
	lb := Label{}
	if strings.Contains(text, "?") {
		lb.isEx = true
	} else {
		lb.isEx = false
	}
	text = strings.TrimSpace(strings.Trim(text, "L"))
	ss := strings.Split(text, ",")
	for _, s := range ss {
		s = strings.TrimSpace(s)
		switch s[0] {
		case '"':
			str = strings.Trim(s, "\"")
		case 's':
			tt := strings.TrimSpace(strings.Trim(s, "s"))
			size, err = strconv.ParseFloat(tt, 32)
			if err != nil { log.Print("parse.doLabel:", err)}
		case 'x':
			tt := strings.TrimSpace(strings.Trim(s, "x"))
			x, err = strconv.ParseFloat(tt, 32)
			if err != nil { log.Print("parse.doLabel:", err)}
		case 'y':
			tt := strings.TrimSpace(strings.Trim(s, "y"))
			y, err = strconv.ParseFloat(tt, 32)
			if err != nil { log.Print("parse.doLabel:", err)}
		}
	}
	if size == 0 { size = 35 }
	lb.str = str
	lb.size = float32(size)
	lb.x = float32(x)
	lb.y = float32(y)
	return lb
}

func posLabels (ll []Label) {
	ln := len(ll)
	if ln < 2 { return }
	for _, lb := range ll {
		if lb.x != 0 || lb.y != 0 { return }
	}

	h := ll[0].size*float32(ln-1)
	for i := range ll {
		f := float32(i)
		ll[i].y = -f*ll[i].size + h/2
	}
}

func doXact(str string) Xacts {
	var (
		isEx bool
		ansF float32
		isNum bool
		ansS string
	)
	isEx = true
	str = strings.TrimSpace(strings.Trim(str, "I"))
	ss := strings.Split(str, ",")
	for _, s := range ss {
		s = strings.TrimSpace(s)
		switch s[0] {
		case 'f':
			tt := strings.TrimSpace(strings.TrimPrefix(s, "f"))
			ansF, isNum = prs(tt)
		case 's':
			ansS = strings.TrimSpace(strings.TrimPrefix(s, "s"))
		}
	}
	return Xacts{isEx: isEx, ansF: ansF, isNum: isNum, ansS: ansS}
}
