package player

import (
	"os"
	"time"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/mp3"
	"github.com/gopxl/beep/v2/speaker"
)

type PlayingMusic struct {
	Erro error
	stremando beep.StreamSeekCloser
	formato beep.Format
	Controle *beep.Ctrl
}

func CarregarMusica(caminho string) (P PlayingMusic) {
	f, err := os.Open(caminho)

	if err != nil {
		P.Erro = err
		return P
	}

	P.stremando, P.formato, err = mp3.Decode(f)
	P.Erro = err

	return P
}

func (P PlayingMusic) IniciarSpeaker() error {
	speaker.Init(P.formato.SampleRate, P.formato.SampleRate.N(time.Second / 20))
	return nil
}