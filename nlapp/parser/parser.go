package parser

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

// TODO colours

type Label struct {
	str string
	size, x, y float32
}

type shape struct {
	labels []Label
}

type Things struct {
	sub string
	shape
	/*anim
	xact
	zero*/
}

func (th Things) Sub() string { return th.sub }

func (th Things) Labels() []Label { return th.labels }

func (lb Label) Label() (str string, size, x, y float32) {
	return lb.str, lb.size, lb.x, lb.y
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

func Do(raw string) []Things {
	ii := indices(raw, "[]")
	tt := []Things{}
	i, k := 0, 0
	for i < len(ii) {
		if raw[ii[i+1]] != ']' {
			log.Println("parser.doRaw: square brackets mismatch")
		}
		switch raw[ii[i]-1] {
		case 't':
			tt = append(tt, Things{sub: raw[ii[i]+1:ii[i+1]]})
			k++
		case 'S':
			tt[k-1].shape = doShape(raw[ii[i]-1:ii[i+1]])
		}
		i += 2
	}
	return tt
}

func doShape(str string) shape {
	str = strings.TrimSpace(strings.Trim(str, "S[]"))
	ss := strings.Split(str, "\n")
	ll := []Label{}
	for _, s := range ss {
		s = strings.TrimSpace(s)
		switch s[0] {
		case 'L':
			ll = append(ll, doLabel(s))
		}
	}
	posLabels(ll)
	return shape{labels: ll}
}

func posLabels (ll []Label) {
	ln := len(ll)
	if ln < 2 { return }
	for _, lb := range ll {
		if lb.x != 0 || lb.y != 0 { return }
	}

	h := ll[0].size*float32(ln-1)
	println(ln)
	for i := range ll {
		f := float32(i)
		ll[i].y = -f*ll[i].size + h/2
		print(ll[i].y)
		print("	")
	}
	println()
}

func doLabel(text string) Label {
	var (
		str string
		size, x, y float64
		err error
	)
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
			if err != nil { log.Print("parser.doLabel: ", err)}
		case 'x':
			tt := strings.TrimSpace(strings.Trim(s, "x"))
			x, err = strconv.ParseFloat(tt, 32)
			if err != nil { log.Print("parser.doLabel: ", err)}
		case 'y':
			tt := strings.TrimSpace(strings.Trim(s, "y"))
			y, err = strconv.ParseFloat(tt, 32)
			if err != nil { log.Print("parser.doLabel: ", err)}
		}
	}
	if size == 0 { size = 35 }
	return Label{str: str, size: float32(size), x: float32(x), y: float32(y)}
}
