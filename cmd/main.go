package main

import (
	"log"

	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/core"
)

func main() {
	app := core.New()

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
