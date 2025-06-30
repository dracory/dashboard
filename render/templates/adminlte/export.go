package adminlte

import (
	templates "github.com/dracory/dashboard/render/templates"
	"github.com/dracory/dashboard/render/templates/shared"
)

// NewAdminLTETemplate creates a new instance of the AdminLTE template
func NewAdminLTETemplate() shared.Template {
	return &AdminLTETemplate{}
}

func init() {
	templates.RegisterTemplateInitializer(func() shared.Template {
		return NewAdminLTETemplate()
	})
}
