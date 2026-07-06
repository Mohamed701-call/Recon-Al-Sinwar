package core

import (
	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/logger"
	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/modules/ports"
	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/report"
)

func Stage3() error {

	logger.Info("===================================")
	logger.Info("Stage 3 : Naabu")
	logger.Info("===================================")

	if err := ports.RunNaabu(); err != nil {
		return err
	}

	if err := report.ParsePorts(); err != nil {
		return err
	}

	logger.Info("Naabu Finished")

	return nil
}
