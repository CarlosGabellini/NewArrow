package main

import (
	"context"
	play "NewArrow/interno/player"
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