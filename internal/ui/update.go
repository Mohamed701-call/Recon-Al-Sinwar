package ui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type splashFinishedMsg struct{}

func waitSplash() tea.Cmd {
	return tea.Tick(
		2*time.Second,
		func(time.Time) tea.Msg {
			return splashFinishedMsg{}
		},
	)
}

func (m Model) Init() tea.Cmd {
	return waitSplash()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case splashFinishedMsg:

		if m.CurrentScreen == SplashScreen {
			m.CurrentScreen = DashboardScreen
		}

	case tea.WindowSizeMsg:

		m.Width = msg.Width
		m.Height = msg.Height

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c":
			return m, tea.Quit

		case "up":

			if m.CurrentScreen == DashboardScreen {

				if m.Selected > 0 {
					m.Selected--
				}

			}

		case "down":

			if m.CurrentScreen == DashboardScreen {

				if m.Selected < len(m.MenuItems)-1 {
					m.Selected++
				}

			}

		case "enter":

			if m.CurrentScreen == DashboardScreen {

				switch m.Selected {

				case 0:
					m.CurrentScreen = ReconScreen

				case 1:
					m.CurrentScreen = ScopeScreen

				case 3:
					m.CurrentScreen = DoctorScreen

				case 4:
					m.CurrentScreen = SettingsScreen

				case 5:
					return m, tea.Quit

				}

			}

		case "esc":

			m.CurrentScreen = DashboardScreen

		}

	}

	return m, nil
}
