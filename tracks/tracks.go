package tracks

import (
	"bytes"
	"log"
	"embed"
	"path"
	//"strconv"

	//"github.com/hajimehoshi/ebiten"
	//"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

const sr = 48000

// //go:embed pow
// var files embed.FS
//var fnames []string

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
	t.context = audio.NewContext(sr)
	t.initOwnTracks()
	//t.initFiles(dir, files)
	return t
}

func (t *Tracks) InitFiles(dir string, files embed.FS) {
	fnames, err := files.ReadDir(dir)
	if err != nil { log.Fatal("initFiles: ", err) }
	t.tracks = make([]*audio.Player, len(fnames))

	for i, v := range fnames {
		f, err := files.ReadFile(path.Join(dir, v.Name()))
		if err != nil { log.Fatal("tracks.initFiles() ReadFile: ", err) }

		d, err := vorbis.DecodeF32(bytes.NewReader(f))
		if err != nil { log.Fatal("tracks.initFiles() vorbis.DecodeF32: ", err) }

		tr, err := t.context.NewPlayerF32(d)
		if err != nil { log.Fatal("tracks.initFiles() NewPlayerF32: ", err) }
		t.tracks[i] = tr
	}
}//*/

func (t *Tracks) initOwnTracks() {
	// t.initFiles("pow")

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

func (t *Tracks) PlayCorrect() {
	if !t.IsPlaying() {
		t.correct.Rewind()
		t.correct.Play()
	}
}

func (t *Tracks) PlayWrong() {
	if !t.IsPlaying() {
		t.wrong.Rewind()
		t.wrong.Play()
	}
}

func (t *Tracks) Play() {
	if !t.IsPlaying() { t.tracks[t.curTrack].Play() }
}

func (t *Tracks) Pause() {
	if t.IsPlaying() {
		t.tracks[t.curTrack].Pause()
	} else {
		t.tracks[t.curTrack].Play()
	}
}

func (t *Tracks) nextTrack() {
	if t.curTrack < len(t.tracks)-1 { t.curTrack++ }
}

func (t *Tracks) prevTrack() {
	if t.curTrack > 0 { t.curTrack-- }
}

func (t *Tracks) Next() {
	if !t.IsPlaying() { t.tracks[t.curTrack].Play() }
	if !t.IsPlaying() {
		t.nextTrack()
		t.tracks[t.curTrack].Play()
	}
}

func (t *Tracks) IsPlaying() bool {
	current := t.tracks[t.curTrack].IsPlaying()
	correct := t.correct.IsPlaying()
	wrong := t.wrong.IsPlaying()
	return current || correct || wrong
}
