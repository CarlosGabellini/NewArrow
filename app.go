package main

import (
	play "NewArrow/internalGo/player"
	"context"
)

// App struct
type App struct {
	ctx context.Context
	play.Gerador1
	
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		Gerador1: play.Gerador1{},
	}
}