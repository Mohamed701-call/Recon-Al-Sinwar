package widgets

import "github.com/Mohamed701-call/Recon-Al-Sinwar/internal/theme"

func Footer() string {

	return theme.Footer.Render(

		"↑ ↓ Move   Enter Select   ESC Back   Ctrl+C Exit",
	)

}
