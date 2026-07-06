package widgets

import "github.com/Mohamed701-call/Recon-Al-Sinwar/internal/theme"

func Success(text string) string {

	return theme.Success.Render("✔ " + text)

}

func Error(text string) string {

	return theme.Error.Render("✖ " + text)

}

func Warning(text string) string {

	return theme.Warning.Render("⚠ " + text)

}
