package player

var PlayerMusic PlayingMusic

func ToqueMusica(pathWay string) error {

	PlayerMusic = CarregarMusica(pathWay)

	if PlayerMusic.Erro != nil {
		return PlayerMusic.Erro
	}

	PlayerMusic.Erro = PlayerMusic.IniciarSpeaker()
	PlayerMusic.TocandoMusica()

	return nil
}

func PararMusica() {
	PlayerMusic.AlternarPausa()
}