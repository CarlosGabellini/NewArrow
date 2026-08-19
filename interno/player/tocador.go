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
	f, err := os.Open("interno/diretorioMusicas/FRONT LINE ASSEMBLY (Feat. Jimmy Urine) - Rock Me Amadeus.mp3")

	//Outra pegadinha, na hora de fazer o binario, nao pode colocar estes caminhos, o binario altera como funciona
	//o caminho das pastas adicionadas, isso tambem acontece na hora de localizar as outras paginas no html.

	if err != nil {
		return err
	}

	//O que faz a magica acontecer aqui;
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