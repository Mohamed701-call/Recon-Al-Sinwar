package report

import (
	"bufio"
	"encoding/json"
	"os"
	"strings"

	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/common"
	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/types"
)

func ParseKatana() error {

	file, err := os.Open(common.KatanaJSON())
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		var r types.CrawlResult

		if err := json.Unmarshal(scanner.Bytes(), &r); err != nil {
			continue
		}

		if r.URL == "" {
			continue
		}

		_ = common.AppendLine(common.URLsFile(), r.URL)

		url := strings.ToLower(r.URL)

		if strings.Contains(url, ".js") {
			_ = common.AppendLine(common.JavaScriptFile(), r.URL)
		}

		if strings.Contains(url, "/api") ||
			strings.Contains(url, "graphql") ||
			strings.Contains(url, "swagger") {

			_ = common.AppendLine(common.APIFile(), r.URL)
		}

		if strings.Contains(url, "?") {
			_ = common.AppendLine(common.ParametersFile(), r.URL)
		}

		if isInteresting(url) {
			_ = common.AppendLine(common.InterestingFile(), r.URL)
		}
	}

	return scanner.Err()
}

func isInteresting(url string) bool {

	keywords := []string{
		"admin",
		"login",
		"dashboard",
		"upload",
		"backup",
		"config",
		"debug",
		"api",
		"graphql",
		"swagger",
		"oauth",
		"auth",
		"token",
		".env",
		".git",
	}

	for _, k := range keywords {

		if strings.Contains(url, k) {
			return true
		}

	}

	return false
}
