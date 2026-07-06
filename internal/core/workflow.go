package core

import (
	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/common"
	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/logger"
	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/modules/discovery"
)

func StartWorkflow() error {

	if err := common.EnsureWorkspace(); err != nil {
		return err
	}

	logger.Info("========================================")
	logger.Info("Recon Al-Sinwar Started")
	logger.Info("========================================")

	result, err := discovery.RunSubfinder()
	if err != nil {
		return err
	}

	logger.Info("Subdomains Found:", result.Count)

	if err := Stage2(); err != nil {
		return err
	}

	if err := Stage3(); err != nil {
		return err
	}

	if err := Stage4(); err != nil {
		return err
	}
	if err := Stage4(); err != nil {
		return err
	}

	if err := Stage5(); err != nil {
		return err
	}

	if err := Stage6(); err != nil {
		return err
	}

	if err := Stage7(); err != nil {
		return err
	}

	logger.Info("========================================")
	logger.Info("Workflow Finished")
	logger.Info("========================================")

	return nil
}
