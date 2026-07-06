package common

import "os"

func EnsureWorkspace() error {

	dirs := []string{
		WorkspaceDir,
		ScopeDir,
		SubdomainsDir,
		HTTPDir,
		PortsDir,
		CrawlDir,
		ScreenshotsDir,
		NucleiDir,
		ReportsDir,
		LogsDir,
		TempDir,
		FuzzDir(),
	}

	for _, dir := range dirs {

		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}

	}

	return nil
}
