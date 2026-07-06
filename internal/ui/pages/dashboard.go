package pages

import (
	"strings"

	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/ui/widgets"
)

func Dashboard(menu []string, selected int) string {

	var b strings.Builder

	b.WriteString(widgets.Banner())

	b.WriteString("\n")

	b.WriteString(

		widgets.Box(

			"Main Menu",

			widgets.Sidebar(menu, selected),
		),
	)

	b.WriteString("\n")

	b.WriteString(

		widgets.StatusBar("Ready"),
	)

	b.WriteString("\n")

	b.WriteString(

		widgets.Footer(),
	)

	return b.String()

}
