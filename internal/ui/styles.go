package ui

func Box(title string, content string) string {

	return BoxStyle.Render(

		TitleStyle.Render(title) +

			"\n\n" +

			content,
	)
}

func Success(text string) string {

	return SuccessStyle.Render("✔ " + text)

}

func Error(text string) string {

	return ErrorStyle.Render("✖ " + text)

}

func Warning(text string) string {

	return WarningStyle.Render("⚠ " + text)

}

func Info(text string) string {

	return InfoStyle.Render(text)

}
