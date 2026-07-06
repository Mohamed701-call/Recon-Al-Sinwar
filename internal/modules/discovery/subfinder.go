package discovery

import (
	"bufio"
	"os"
	"os/exec"
	"time"

	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/common"
	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/types"
)

func RunSubfinder() (*types.Result, error) {

	start := time.Now()

	tool, err := common.FindTool("subfinder")

	if err != nil {
		return nil, err
	}

	out, err := os.Create(common.SubfinderOutput())

	if err != nil {
		return nil, err
	}

	defer out.Close()

	cmd := exec.Command(

		tool,

		"-dL",
		common.TargetsFile(),

		"-all",

		"-silent",
	)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {

		return nil, err

	}

	scanner := bufio.NewScanner(stdout)

	seen := make(map[string]struct{})

	count := 0

	for scanner.Scan() {

		host := scanner.Text()

		if host == "" {
			continue
		}

		if _, ok := seen[host]; ok {
			continue
		}

		seen[host] = struct{}{}

		count++

		out.WriteString(host + "\n")
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if err := cmd.Wait(); err != nil {
		return nil, err
	}

	return &types.Result{

		Module: "Subfinder",

		OutputFile: common.SubfinderOutput(),

		Count: count,

		Started: start,

		Finished: time.Now(),

		Success: true,
	}, nil

}
