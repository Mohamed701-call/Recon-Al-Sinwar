package http

import (
	"os/exec"

	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/common"
)

func RunHTTPX() error {

	tool, err := common.FindTool("httpx")
	if err != nil {
		return err
	}

	cmd := exec.Command(
		tool,

		"-l",
		common.SubfinderOutput(),

		"-silent",

		"-json",

		"-title",

		"-tech-detect",

		"-status-code",

		"-ip",

		"-cdn",

		"-o",
		common.HTTPJSON(),
	)

	return cmd.Run()
}
