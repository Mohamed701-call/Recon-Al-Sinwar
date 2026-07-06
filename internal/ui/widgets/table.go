package widgets

import (
	"strings"

	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/theme"
)

func Table(headers []string, rows [][]string) string {

	var b strings.Builder

	for _, h := range headers {
		b.WriteString(theme.Title.Render(h))
		b.WriteString("\t")
	}

	b.WriteString("\n")

	for _, row := range rows {

		for _, col := range row {

			b.WriteString(col)
			b.WriteString("\t")

		}

		b.WriteString("\n")

	}

	return b.String()

}
