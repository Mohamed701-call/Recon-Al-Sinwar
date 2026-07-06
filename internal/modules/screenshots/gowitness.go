package screenshots

import (
	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/common"
)

func RunGoWitness() error {

	tool, err := common.FindTool("gowitness")
	if err != nil {
		return err
	}

	runner := common.NewRunner()

	_, err = runner.Run(

		tool,

		"file",

		"-f",
		common.InterestingDirectoriesFile(),

		"--threads",
		"8",

		"--write-db",

		"--screenshot-path",
		common.ScreenshotsDir,
	)

	return err
}
