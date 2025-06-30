package tabler

import "github.com/dracory/dashboard/render/templates/shared"

// NewTablerTemplate creates a new instance of the Tabler template
func NewTablerTemplate() shared.Template {
	return &TablerTemplate{}
}
