package main

import (
	"context"
	player "NewArrow/interno/player"
)

// App struct
type App struct {
	ctx context.Context
	player.Tocador
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		Tocador: player.Tocador{},
	}
}