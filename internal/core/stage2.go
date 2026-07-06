package core

import (
	httpmodule "github.com/Mohamed701-call/Recon-Al-Sinwar/internal/modules/http"

	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/logger"
	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/report"
)

func Stage2() error {

	logger.Info("===================================")
	logger.Info("Stage 2")
	logger.Info("Running HTTPX")
	logger.Info("===================================")

	if err := httpmodule.RunHTTPX(); err != nil {
		return err
	}

	logger.Info("Parsing HTTP Results")

	if err := report.ParseHTTPReport(); err != nil {
		return err
	}

	logger.Info("HTTP Stage Finished")

	return nil
}
