package fuzz

import (
	"bufio"
	"os"
	"os/exec"

	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/common"
)

func RunFFUF() error {

	if !CheckWordlists() {

		if err := DownloadWordlists(); err != nil {

			return err

		}

	}

	file, err := os.Open(common.FuzzTargetsFile())

	if err != nil {

		return err

	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		target := scanner.Text()

		if target == "" {

			continue

		}

		err := runDirectories(target)

		if err != nil {

			continue

		}

	}

	return nil

}

func runDirectories(target string) error {

	tool, err := common.FindTool("ffuf")

	if err != nil {

		return err

	}

	cmd := exec.Command(

		tool,

		"-u", target+"FUZZ",

		"-w", "tools/wordlists/directories/raft-medium-directories.txt",

		"-mc", "200,204,301,302,307,401,403",

		"-ac",

		"-ignore-body",

		"-rate", "150",

		"-recursion",

		"-recursion-depth", "2",

		"-of", "json",

		"-o",

		common.FuzzDirectoriesJSON(),
	)

	return cmd.Run()

}
