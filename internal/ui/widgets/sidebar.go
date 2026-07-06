package widgets

import (
	"strings"

	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/theme"
)

func Sidebar(items []string, selected int) string {

	var b strings.Builder

	for i, item := range items {

		if i == selected {

			b.WriteString(

				theme.Selected.Render("> " + item),
			)

		} else {

			b.WriteString(

				theme.Info.Render("  " + item),
			)

		}

		b.WriteString("\n")

	}

	return b.String()

}
