package app

import (
	"log"
)

type App struct {
}

func New() *App {
	return &App{}
}

func (a *App) Run() error {
	log.Println("Lumina-Mesh Node is running.")
	// Keep the process alive or start servers here
	return nil
}
