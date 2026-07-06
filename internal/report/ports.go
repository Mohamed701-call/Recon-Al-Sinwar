package report

import (
	"bufio"
	"encoding/json"
	"os"
	"strconv"

	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/common"
	"github.com/Mohamed701-call/Recon-Al-Sinwar/internal/types"
)

func ParsePorts() error {

	file, err := os.Open(common.NaabuJSON())
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		var result types.PortResult

		if err := json.Unmarshal(scanner.Bytes(), &result); err != nil {
			continue
		}

		line := result.Host + ":" + strconv.Itoa(result.Port)

		if err := common.AppendLine(common.OpenPortsFile(), line); err != nil {
			return err
		}
	}

	return scanner.Err()
}
