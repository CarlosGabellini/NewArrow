package player

import (
	"os"
	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/mp3"
	"github.com/gopxl/beep/v2/speaker"
	"time"
)

type Tocador struct {
	ctrl *beep.Ctrl
	stream beep.StreamSeekCloser
}

func (t *Tocador) ToqueMusica() error {

	//Para tocar a musica, o caminho nao pode ser relativo, no Wails, o caminho sempre comeca
	// a partir da raiz do projeto, e entao selecionamos a pasta a partir disso, nao posso me
	// esquecer de fazer isso.
	f, err := os.Open("interno/diretorioMusicas/ABBA - Dancing Queen (Official Music Video).mp3")

	if err != nil {
		return err
	}

	stream, format, err := mp3.Decode(f)

	if err != nil {
		return err
	}

	t.stream = stream

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	t.ctrl = &beep.Ctrl{Streamer: stream, Paused: false}
	speaker.Play(t.ctrl)

	return nil
}