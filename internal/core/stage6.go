package core

import (
	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/logger"
	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/modules/screenshots"
)

func Stage6() error {

	logger.Info("===================================")
	logger.Info("Stage 6 : GoWitness")
	logger.Info("===================================")

	if err := screenshots.RunGoWitness(); err != nil {
		return err
	}

	logger.Info("GoWitness Finished")

	return nil
}
