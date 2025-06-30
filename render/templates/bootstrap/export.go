package bootstrap

import (
	templates "github.com/dracory/dashboard/render/templates"
	"github.com/dracory/dashboard/render/templates/shared"
)

// NewBootstrapTemplate creates a new instance of the Bootstrap 5 template
func NewBootstrapTemplate() shared.Template {
	return &BootstrapTemplate{}
}

func init() {
	templates.RegisterTemplateInitializer(func() shared.Template {
		return NewBootstrapTemplate()
	})
}
