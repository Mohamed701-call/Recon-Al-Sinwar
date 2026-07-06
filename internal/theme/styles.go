package theme

import "github.com/charmbracelet/lipgloss"

var (
	PrimaryColor    = lipgloss.Color("#39FF14")
	SecondaryColor  = lipgloss.Color("#00DD55")
	SuccessColor    = lipgloss.Color("#00FF7F")
	WarningColor    = lipgloss.Color("#FFD700")
	ErrorColor      = lipgloss.Color("#FF4C4C")
	BackgroundColor = lipgloss.Color("#111111")
	BorderColor     = lipgloss.Color("#00AA44")
	TextColor       = lipgloss.Color("#F5F5F5")
	MutedColor      = lipgloss.Color("#777777")
)

var (
	Title = lipgloss.NewStyle().
		Foreground(PrimaryColor).
		Bold(true)

	Subtitle = lipgloss.NewStyle().
			Foreground(SecondaryColor)

	Box = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(BorderColor).
		Padding(1, 2)

	Success = lipgloss.NewStyle().
		Foreground(SuccessColor).
		Bold(true)

	Error = lipgloss.NewStyle().
		Foreground(ErrorColor).
		Bold(true)

	Warning = lipgloss.NewStyle().
		Foreground(WarningColor).
		Bold(true)

	Info = lipgloss.NewStyle().
		Foreground(TextColor)

	Footer = lipgloss.NewStyle().
		Foreground(MutedColor)

	Status = lipgloss.NewStyle().
		Foreground(PrimaryColor)

	Selected = lipgloss.NewStyle().
			Foreground(BackgroundColor).
			Background(PrimaryColor).
			Bold(true)

	Input = lipgloss.NewStyle().
		Foreground(PrimaryColor)
)
