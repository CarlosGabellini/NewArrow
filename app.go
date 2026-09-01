package main

import (
	"NewArrow/interno/dir1org"
	play "NewArrow/interno/player"
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func (a *App) PlayTrack() error {
	return play.ToqueMusica()
}

func (a *App) TogglePause() {
	play.PararMusica()
}


//Funcao importante que vai escanear o JSON e retornar a lista de musicas diretamente para
//o JavaScript.

func (a *App) RetornandoLista() ([]dir1org.ListaMusicas, error) {
	var CaminhoJSON string = "./myJSONfile.json"
	lista, err := dir1org.EscanearJSON(CaminhoJSON)

	if err != nil {
		return nil, err
	}

	return lista, nil
}

func (a *App) AtualizarCache() (string, error) {
	diretorioHome, err := dir1org.EncontreHome()

	if err != nil {
		return "", err
	}

	lista2, err := dir1org.Liste_a_pasta(diretorioHome)
	if err != nil {
		return "", err
	}

	var sucessoMensagem string = fmt.Sprintf("O arquivo foi um sucesso! primeira musica %s", 
		lista2[0].Nome_da_musica)

	return sucessoMensagem, nil
}

func (a *App) ListarDiretorios() ([]string, error) {
	caminhoHOME, err := dir1org.EncontreHome()

	if err != nil {
		return nil, err
	}
	
	diretoriosEncontrados, err := dir1org.VerOsdiretorios_musics(caminhoHOME)

	if err != nil {
		return nil, err
	}
	
	return diretoriosEncontrados, err
}