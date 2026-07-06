package widgets

import "github.com/Mohamed701-call/Recon-Al-Sinwar/internal/theme"

func Input(label string, value string) string {

	return theme.Box.Render(

		theme.Title.Render(label) +

			"\n\n" +

			value,
	)

}
