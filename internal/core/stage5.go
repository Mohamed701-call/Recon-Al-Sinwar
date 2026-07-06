package core

import (
	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/logger"
	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/modules/fuzz"
	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/report"
)

func Stage5() error {

	logger.Info("===================================")
	logger.Info("Stage 5 : Smart FFUF")
	logger.Info("===================================")

	if err := report.BuildFuzzTargets(30); err != nil {
		return err
	}

	logger.Info("Targets Ready")

	if err := fuzz.RunFFUF(); err != nil {
		return err
	}

	logger.Info("FFUF Finished")

	if err := report.ParseFFUF(); err != nil {
		return err
	}

	logger.Info("FFUF Results Parsed")

	return nil
}
