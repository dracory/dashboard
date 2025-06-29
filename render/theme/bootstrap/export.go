package bootstrap

import (
	"github.com/dracory/dashboard/render/theme"
	"github.com/dracory/dashboard/render/theme/shared"
)

// NewBootstrapTheme creates a new instance of the Bootstrap 5 theme
func NewBootstrapTheme() shared.Theme {
	return &BootstrapTheme{}
}

func init() {
	theme.RegisterThemeInitializer(func() shared.Theme {
		return NewBootstrapTheme()
	})
}
