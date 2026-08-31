package main

import (
	"NewArrow/interno/dir1org"
	play "NewArrow/interno/player"
	"context"
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

func (a *App) Creates_list() ([]dir1org.ListaMusicas, error) {
	DiretorioMusica, err := dir1org.Pasta_de_musica()
	if err != nil {
		return nil, err
	}

	lista, err := dir1org.Liste_a_pasta(DiretorioMusica)
	if err != nil {
		return lista, err
	}

	return lista, nil
}