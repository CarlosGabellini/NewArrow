package player

var PlayerMusic PlayingMusic

func ToqueMusica() error {

	PlayerMusic = CarregarMusica("interno/diretorioMusicas/Tu Pai.mp3")

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