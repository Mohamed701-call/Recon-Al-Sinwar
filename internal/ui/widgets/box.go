package widgets

import "github.com/Mohamed701-call/Recon-Al-Sinwar/internal/theme"

func Box(title string, body string) string {

	content := theme.Title.Render(title) + "\n\n" + body

	return theme.Box.Render(content)

}
