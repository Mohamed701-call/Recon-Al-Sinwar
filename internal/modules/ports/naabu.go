package ports

import (
	"os/exec"

	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/common"
)

func RunNaabu() error {

	tool, err := common.FindTool("naabu")
	if err != nil {
		return err
	}

	cmd := exec.Command(
		tool,

		"-list",
		common.LiveHostsFile(),

		"-top-ports",
		"1000",

		"-rate",
		"1000",

		"-silent",

		"-json",

		"-o",
		common.NaabuJSON(),
	)

	return cmd.Run()
}
