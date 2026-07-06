package widgets

import "github.com/Mohamed701-call/Recon-Al-Sinwar/internal/theme"

func StatusBar(status string) string {

	return theme.Box.Render(

		theme.Status.Render("Status : " + status),
	)

}
