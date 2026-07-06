package ui

import tea "github.com/charmbracelet/bubbletea"

type App struct {
	program *tea.Program
}

func New() *App {
	model := NewModel()

	return &App{
		program: tea.NewProgram(
			model,
			tea.WithAltScreen(),
		),
	}
}

func (a *App) Run() error {
	_, err := a.program.Run()
	return err
}
