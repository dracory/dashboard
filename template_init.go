// Package dashboard provides the main dashboard functionality.
package dashboard

import (
	"github.com/dracory/dashboard/render/templates"
	"github.com/dracory/dashboard/render/templates/shared"
	adminlte "github.com/dracory/dashboard/render/templates/adminlte"
	bootstrap "github.com/dracory/dashboard/render/templates/bootstrap"
	tabler "github.com/dracory/dashboard/render/templates/tabler"
)

// init registers all available templates
func init() {
	// Register Tabler template
	templates.RegisterTemplateInitializer(func() shared.Template {
		return tabler.NewTablerTemplate()
	})

	// Register Bootstrap 5 template
	templates.RegisterTemplateInitializer(func() shared.Template {
		return bootstrap.NewBootstrapTemplate()
	})

	// Register AdminLTE template
	templates.RegisterTemplateInitializer(func() shared.Template {
		return adminlte.NewAdminLTETemplate()
	})

	// Initialize all registered templates
	templates.InitializeRegisteredTemplates()
}
