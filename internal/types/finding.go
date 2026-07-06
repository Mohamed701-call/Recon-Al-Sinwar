package types

type Finding struct {
	TemplateID string `json:"template_id"`
	Name       string `json:"name"`
	Severity   string `json:"severity"`

	Host    string `json:"host"`
	Matched string `json:"matched"`

	Description string `json:"description"`

	Reference []string `json:"reference"`

	CVE []string `json:"cve"`

	CVSS float64 `json:"cvss"`

	Tags []string `json:"tags"`
}
type FindingGroup struct {
	TemplateID  string `json:"template_id"`
	Name        string `json:"name"`
	Severity    string `json:"severity"`
	Description string `json:"description"`

	Hosts []string `json:"hosts"`

	References []string `json:"references"`

	Tags []string `json:"tags"`
}
