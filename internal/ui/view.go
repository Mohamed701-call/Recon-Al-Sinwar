package ui

import (
	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/ui/pages"
)

func (m Model) View() string {

	switch m.CurrentScreen {

	case SplashScreen:
		return pages.Splash()

	case DashboardScreen:
		return pages.Dashboard(
			m.MenuItems,
			m.Selected,
		)

	case ScopeScreen:
		return "Scope Manager"

	case ReconScreen:
		return "Recon"

	case DoctorScreen:
		return "Doctor"

	case SettingsScreen:
		return "Settings"

	default:
		return "Unknown Screen"

	}

}
