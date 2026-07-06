package widgets

import (
	"fmt"
	"strings"

	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/theme"
)

func Progress(percent float64) string {

	if percent < 0 {
		percent = 0
	}

	if percent > 100 {
		percent = 100
	}

	width := 40

	filled := int((percent / 100) * float64(width))

	bar := strings.Repeat("█", filled) +
		strings.Repeat("░", width-filled)

	return theme.Status.Render(
		fmt.Sprintf("[%s] %.0f%%", bar, percent),
	)
}
