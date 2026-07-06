package crawl

import (
	"os/exec"

	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/common"
)

func RunKatana() error {

	tool, err := common.FindTool("katana")
	if err != nil {
		return err
	}

	cmd := exec.Command(
		tool,

		"-list",
		common.LiveHostsFile(),

		"-silent",

		"-jsonl",

		"-js-crawl",

		"-known-files",
		"all",

		"-o",
		common.KatanaJSON(),
	)

	return cmd.Run()
}
