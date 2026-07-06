package pages

import (
	"strings"

	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/theme"
	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/ui/widgets"
)

const Version = "v1.0.0"

func Splash() string {

	var b strings.Builder

	b.WriteString("\n")

	// Banner
	b.WriteString(widgets.Banner())

	b.WriteString("\n")

	// Title
	b.WriteString(theme.Title.Render("Recon Al-Sinwar"))

	b.WriteString("\n")

	// Description
	b.WriteString(theme.Subtitle.Render("Bug Hunting Recon Framework"))

	b.WriteString("\n\n")

	// Version
	b.WriteString(theme.Info.Render("Version : " + Version))

	b.WriteString("\n\n")

	// Loading
	b.WriteString(theme.Success.Render("Loading..."))

	b.WriteString("\n")

	return b.String()
}
