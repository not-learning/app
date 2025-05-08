package tracks

import (
	"bytes"
	"embed"
	"log"
	"path"
	"strings"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

const sr = 48000

//go:embed answers
var ansFiles embed.FS

type Tracks struct {
	context *audio.Context
	//player  *audio.Player
	correct *audio.Player
	wrong   *audio.Player

	tracks   []*audio.Player
	curTrack int
}

func Init() *Tracks {
	t := &Tracks{}
	if audio.CurrentContext() == nil {
		t.context = audio.NewContext(sr)
	} else {
		t.context = audio.CurrentContext()
	}
	t.initOwnTracks()
	return t
}

func (t *Tracks) InitFiles(dir string, files embed.FS) {
	fnames, err := files.ReadDir(dir)
	if err != nil { log.Fatal("initFiles: ", err) }

	for _, v := range fnames {
		if !strings.HasSuffix(v.Name(), "ogg") {
			return
		}
		f, err := files.ReadFile(path.Join(dir, v.Name()))
		if err != nil { log.Fatal("tracks.initFiles() ReadFile: ", err) }

		d, err := vorbis.DecodeF32(bytes.NewReader(f))
		if err != nil { log.Fatal("tracks.initFiles() vorbis.DecodeF32: ", err) }

		tr, err := t.context.NewPlayerF32(d)
		if err != nil { log.Fatal("tracks.initFiles() NewPlayerF32: ", err) }
		t.tracks = append(t.tracks, tr)
	}
} //*/

func (t *Tracks) initOwnTracks() {
	f, err := ansFiles.ReadFile("answers/correct/1.ogg")
	if err != nil { log.Fatal("Audio.Init() ReadFile: ", err) }
	d, err := vorbis.DecodeF32(bytes.NewReader(f))
	if err != nil { log.Fatal("Audio.Init() vorbis.DecodeF32: ", err) }
	t.correct, err = t.context.NewPlayerF32(d)
	if err != nil { log.Fatal("Audio.Init() NewPlayerF32: ", err) }

	f, err = ansFiles.ReadFile("answers/wrong/1.ogg")
	if err != nil { log.Fatal("Audio.Init() ReadFile: ", err) }
	d, err = vorbis.DecodeF32(bytes.NewReader(f))
	if err != nil { log.Fatal("Audio.Init() vorbis.DecodeF32: ", err) }
	t.wrong, err = t.context.NewPlayerF32(d)
	if err != nil { log.Fatal("Audio.Init() NewPlayerF32: ", err) }
}

func (t *Tracks) PlayCorrect(ans bool) {
	if pl, ok := t.IsPlaying(); !pl && ok {
		if e := t.correct.Rewind(); e != nil { log.Println("PlayCorrect: ", e) }
		if e := t.wrong.Rewind(); e != nil { log.Println("PlayWrong: ", e) }

		if ans {
			t.correct.Play()
		} else {
			t.wrong.Play()
		}
	}
}

/*func (t *Tracks) PlayWrong() {
	if !t.IsPlaying() {
	}
}//*/

func (t *Tracks) Play() {
	if pl, ok := t.IsPlaying(); !pl && ok {
		t.tracks[t.curTrack].Play()
	}
}

func (t *Tracks) Pause() {
	pl, ok := t.IsPlaying()
	if !ok { return }
	if pl {
		t.tracks[t.curTrack].Pause()
	} else {
		t.tracks[t.curTrack].Play()
	}
}

func (t *Tracks) nextTrack() {
	if t.curTrack < len(t.tracks)-1 {
		t.curTrack++
	}
}

func (t *Tracks) switchTrack(n int) {
	if n >= 0 && n < len(t.tracks) {
		t.curTrack = n
	}
}

func (t *Tracks) prevTrack() {
	if t.curTrack > 0 {
		t.curTrack--
	}
}

func (t *Tracks) Proceed() {
	if pl, ok := t.IsPlaying(); !pl && ok {
		t.tracks[t.curTrack].Play()
	}
	if pl, ok := t.IsPlaying(); !pl && ok {
		t.nextTrack()
		t.tracks[t.curTrack].Play()
	}
}

func (t *Tracks) Next() {
	if t.tracks == nil || len(t.tracks) == 0 { return }

	t.tracks[t.curTrack].Pause()
	if e := t.tracks[t.curTrack].Rewind(); e != nil {
		log.Println("Prev: ", e)
	}
	t.nextTrack()
}

func (t *Tracks) Switch(n int) {
	if t.tracks == nil || len(t.tracks) == 0 { return }

	t.tracks[t.curTrack].Pause()
	if e := t.tracks[t.curTrack].Rewind(); e != nil {
		log.Println("Prev: ", e)
	}
	t.switchTrack(n)
}

func (t *Tracks) Prev() {
	if t.tracks == nil || len(t.tracks) == 0 { return }

	t.tracks[t.curTrack].Pause()
	if e := t.tracks[t.curTrack].Rewind(); e != nil {
		log.Println("Prev: ", e)
	}
	t.prevTrack()
}

func (t *Tracks) IsPlaying() (pl, ok bool) {
	if t.tracks == nil || len(t.tracks) == 0 {
		return false, false
	}
	current := t.tracks[t.curTrack].IsPlaying()
	correct := t.correct.IsPlaying()
	wrong := t.wrong.IsPlaying()
	return current || correct || wrong, true
}
