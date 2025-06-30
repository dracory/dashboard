package bootstrap

import "github.com/dracory/dashboard/render/templates/shared"

// NewBootstrapTemplate creates a new instance of the Bootstrap template
func NewBootstrapTemplate() shared.Template {
	return &BootstrapTemplate{}
}
