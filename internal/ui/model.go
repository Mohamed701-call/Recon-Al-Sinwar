package ui

type Model struct {
	Width  int
	Height int

	CurrentScreen Screen

	MenuItems []string
	Selected  int
}

func NewModel() Model {

	return Model{

		CurrentScreen: SplashScreen,

		MenuItems: []string{
			"Start Recon",
			"Scope Manager",
			"Reports",
			"Doctor",
			"Settings",
			"Exit",
		},

		Selected: 0,
	}
}
