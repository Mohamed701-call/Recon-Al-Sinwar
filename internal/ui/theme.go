package ui

import "github.com/charmbracelet/lipgloss"

var (
	PrimaryColor = lipgloss.Color("#39FF14")

	SecondaryColor = lipgloss.Color("#00DD55")

	SuccessColor = lipgloss.Color("#00FF7F")

	WarningColor = lipgloss.Color("#FFD700")

	ErrorColor = lipgloss.Color("#FF4C4C")

	BackgroundColor = lipgloss.Color("#111111")

	BorderColor = lipgloss.Color("#00AA44")

	TextColor = lipgloss.Color("#F5F5F5")

	MutedColor = lipgloss.Color("#777777")
)

var (
	TitleStyle = lipgloss.NewStyle().
			Foreground(PrimaryColor).
			Bold(true)

	SubTitleStyle = lipgloss.NewStyle().
			Foreground(SecondaryColor)

	BoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(BorderColor).
			Padding(1, 2)

	SuccessStyle = lipgloss.NewStyle().
			Foreground(SuccessColor).
			Bold(true)

	ErrorStyle = lipgloss.NewStyle().
			Foreground(ErrorColor).
			Bold(true)

	WarningStyle = lipgloss.NewStyle().
			Foreground(WarningColor).
			Bold(true)

	InfoStyle = lipgloss.NewStyle().
			Foreground(TextColor)

	FooterStyle = lipgloss.NewStyle().
			Foreground(MutedColor)

	StatusStyle = lipgloss.NewStyle().
			Foreground(PrimaryColor)

	ProgressStyle = lipgloss.NewStyle().
			Foreground(PrimaryColor)

	InputStyle = lipgloss.NewStyle().
			Foreground(PrimaryColor)

	HighlightStyle = lipgloss.NewStyle().
			Foreground(PrimaryColor).
			Bold(true)

	SelectedStyle = lipgloss.NewStyle().
			Foreground(BackgroundColor).
			Background(PrimaryColor).
			Bold(true)
)
