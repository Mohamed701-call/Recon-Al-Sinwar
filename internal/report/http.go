package report

import (
	"bufio"
	"encoding/json"
	"os"

	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/common"
	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/types"
)

func ParseHTTPReport() error {

	file, err := os.Open(common.HTTPJSON())

	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		var r types.HTTPResult

		if err := json.Unmarshal(scanner.Bytes(), &r); err != nil {

			continue

		}

		if r.URL != "" {

			common.AppendLine(
				common.LiveHostsFile(),
				r.URL,
			)

		}

		if r.IP != "" {

			common.AppendLine(
				common.IPsFile(),
				r.IP,
			)

		}

		for _, tech := range r.Tech {

			common.AppendLine(
				common.TechnologiesFile(),
				tech,
			)

		}

		if r.StatusCode == 200 {

			common.AppendLine(
				common.Status200File(),
				r.URL,
			)

		}

	}

	return scanner.Err()

}
