package report

import (
	"bufio"
	"encoding/json"
	"os"
	"sort"
	"strings"

	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/common"
	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/types"
)

func BuildFuzzTargets(limit int) error {

	file, err := os.Open(common.InterestingFile())
	if err != nil {
		return err
	}
	defer file.Close()

	seen := make(map[string]bool)
	var targets []types.FuzzTarget

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		url := normalize(scanner.Text())

		if url == "" {
			continue
		}

		if seen[url] {
			continue
		}

		seen[url] = true

		targets = append(targets, types.FuzzTarget{
			URL:   url,
			Score: score(url),
		})
	}

	sort.Slice(targets, func(i, j int) bool {
		return targets[i].Score > targets[j].Score
	})

	if limit > len(targets) {
		limit = len(targets)
	}

	out, err := os.Create(common.FuzzTargetsFile())
	if err != nil {
		return err
	}
	defer out.Close()

	for i := 0; i < limit; i++ {
		_, _ = out.WriteString(targets[i].URL + "\n")
	}

	return nil
}

func normalize(url string) string {

	url = strings.TrimSpace(url)

	if url == "" {
		return ""
	}

	parts := strings.Split(url, "/")

	if len(parts) < 4 {
		return url
	}

	return strings.Join(parts[:4], "/") + "/"
}

func score(url string) int {

	s := 0

	keywords := map[string]int{
		"admin":     10,
		"login":     10,
		"dashboard": 9,
		"api":       9,
		"graphql":   10,
		"swagger":   10,
		"upload":    9,
		"backup":    9,
		"config":    9,
		".git":      10,
		".env":      10,
		"auth":      8,
		"oauth":     8,
		"token":     8,
	}

	u := strings.ToLower(url)

	for k, v := range keywords {
		if strings.Contains(u, k) {
			s += v
		}
	}

	return s
}

func ParseFFUF() error {

	file, err := os.Open(common.FuzzDirectoriesJSON())
	if err != nil {
		return err
	}
	defer file.Close()

	var result types.FFUFResult

	if err := json.NewDecoder(file).Decode(&result); err != nil {
		return err
	}

	for _, r := range result.Results {

		saveByStatus(r)

		if isInterestingDirectory(r.URL) {
			_ = common.AppendLine(
				common.InterestingDirectoriesFile(),
				r.URL,
			)
		}
	}

	return nil
}

func saveByStatus(r types.FFUFEntry) {

	switch r.Status {

	case 200:
		_ = common.AppendLine(common.Fuzz200File(), r.URL)

	case 301:
		_ = common.AppendLine(common.Fuzz301File(), r.URL)

	case 302:
		_ = common.AppendLine(common.Fuzz302File(), r.URL)

	case 401:
		_ = common.AppendLine(common.Fuzz401File(), r.URL)

	case 403:
		_ = common.AppendLine(common.Fuzz403File(), r.URL)
	}
}

func isInterestingDirectory(url string) bool {

	url = strings.ToLower(url)

	keywords := []string{
		"admin",
		"login",
		"dashboard",
		"backup",
		"config",
		"upload",
		"private",
		"internal",
		"dev",
		"stage",
		"staging",
		"test",
		"debug",
		"swagger",
		"graphql",
		"jenkins",
		"grafana",
		"kibana",
		"gitlab",
		"phpmyadmin",
		".git",
		".env",
	}

	for _, k := range keywords {
		if strings.Contains(url, k) {
			return true
		}
	}

	return false
}
