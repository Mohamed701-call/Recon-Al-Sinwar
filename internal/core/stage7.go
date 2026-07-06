package core

import (
	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/logger"
	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/modules/vulnerability"
)

func Stage7() error {

	logger.Info("===================================")
	logger.Info("Stage 7 : Nuclei")
	logger.Info("===================================")

	err := vulnerability.RunNuclei(vulnerability.Normal)
	if err != nil {
		return err
	}

	logger.Info("Nuclei Finished")

	findings, err := vulnerability.ParseResults()
	if err != nil {
		return err
	}

	err = vulnerability.SaveFindings(findings)
	if err != nil {
		return err
	}

	groups, err := vulnerability.BuildGroups(findings)
	if err != nil {
		return err
	}

	err = vulnerability.SaveGroups(groups)
	if err != nil {
		return err
	}

	logger.Info("Nuclei Results Parsed")

	return nil
}
