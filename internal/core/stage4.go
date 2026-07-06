package core

import (
	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/logger"
	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/modules/crawl"
	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/report"
)

func Stage4() error {

	logger.Info("===================================")
	logger.Info("Stage 4 : Katana")
	logger.Info("===================================")

	if err := crawl.RunKatana(); err != nil {
		return err
	}

	if err := report.ParseKatana(); err != nil {
		return err
	}

	logger.Info("Katana Finished")

	return nil
}
