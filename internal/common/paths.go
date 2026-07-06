package common

import "path/filepath"

const (
	WorkspaceDir = "workspace"

	ToolsDir = "tools"

	ScopeDir = "workspace/scope"

	SubdomainsDir = "workspace/subdomains"

	HTTPDir = "workspace/http"

	PortsDir = "workspace/ports"

	CrawlDir = "workspace/crawl"

	ScreenshotsDir = "workspace/screenshots"

	NucleiDir = "workspace/nuclei"

	ReportsDir = "workspace/reports"

	LogsDir = "workspace/logs"

	TempDir = "workspace/temp"
)

func TargetsFile() string {
	return filepath.Join(ScopeDir, "targets.txt")
}

func SubfinderOutput() string {
	return filepath.Join(SubdomainsDir, "subfinder.txt")
}

func HTTPJSON() string {
	return filepath.Join(HTTPDir, "httpx.json")
}

func LiveHostsFile() string {
	return filepath.Join(HTTPDir, "live.txt")
}

func Status200File() string {
	return filepath.Join(HTTPDir, "status200.txt")
}

func TechnologiesFile() string {
	return filepath.Join(HTTPDir, "technologies.txt")
}

func IPsFile() string {
	return filepath.Join(HTTPDir, "ips.txt")
}

func NaabuJSON() string {
	return filepath.Join(PortsDir, "naabu.json")
}

func OpenPortsFile() string {
	return filepath.Join(PortsDir, "open_ports.txt")
}

func Tool(name string) string {
	return filepath.Join(ToolsDir, name+".exe")
}
func KatanaJSON() string {
	return filepath.Join(CrawlDir, "katana.json")
}

func URLsFile() string {
	return filepath.Join(CrawlDir, "urls.txt")
}

func JavaScriptFile() string {
	return filepath.Join(CrawlDir, "javascript.txt")
}

func APIFile() string {
	return filepath.Join(CrawlDir, "api.txt")
}

func ParametersFile() string {
	return filepath.Join(CrawlDir, "parameters.txt")
}

func InterestingFile() string {
	return filepath.Join(CrawlDir, "interesting.txt")
}
func FuzzDir() string {
	return filepath.Join(WorkspaceDir, "fuzz")
}

func FuzzTargetsFile() string {
	return filepath.Join(FuzzDir(), "targets.txt")
}

func FuzzDirectoriesJSON() string {
	return filepath.Join(FuzzDir(), "directories.json")
}

func FuzzFilesJSON() string {
	return filepath.Join(FuzzDir(), "files.json")
}

func Fuzz200File() string {
	return filepath.Join(FuzzDir(), "200.txt")
}

func Fuzz301File() string {
	return filepath.Join(FuzzDir(), "301.txt")
}

func Fuzz302File() string {
	return filepath.Join(FuzzDir(), "302.txt")
}

func Fuzz401File() string {
	return filepath.Join(FuzzDir(), "401.txt")
}

func Fuzz403File() string {
	return filepath.Join(FuzzDir(), "403.txt")
}
func InterestingDirectoriesFile() string {

	return filepath.Join(
		FuzzDir(),
		"interesting_directories.txt",
	)

}
func NucleiJSON() string {

	return filepath.Join(
		NucleiDir,
		"results.jsonl",
	)

}

func NucleiSummary() string {

	return filepath.Join(
		NucleiDir,
		"summary.json",
	)

}
func FindingsJSON() string {

	return filepath.Join(
		NucleiDir,
		"findings.json",
	)

}
func GroupsJSON() string {

	return filepath.Join(
		NucleiDir,
		"groups.json",
	)

}
