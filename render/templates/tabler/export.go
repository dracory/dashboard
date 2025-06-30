package tabler

import (
	templates "github.com/dracory/dashboard/render/templates"
	"github.com/dracory/dashboard/render/templates/shared"
)

// NewTablerTemplate creates a new instance of the Tabler template
func NewTablerTemplate() shared.Template {
	return &TablerTheme{}
}

func init() {
	templates.RegisterTemplateInitializer(func() shared.Template {
		return NewTablerTemplate()
	})
}
