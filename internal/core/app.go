package core

import "github.com/Mohamed701-call/Recon-Al-Sinwar/internal/ui"

type App struct {
	ui *ui.App
}

func New() *App {
	return &App{
		ui: ui.New(),
	}
}

func (a *App) Run() error {
	return a.ui.Run()
}
