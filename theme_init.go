// Package dashboard provides the main dashboard functionality.
package dashboard

import (
	"github.com/dracory/dashboard/render/theme"
	adminlte "github.com/dracory/dashboard/render/theme/adminlte"
	bootstrap "github.com/dracory/dashboard/render/theme/bootstrap"
	shared "github.com/dracory/dashboard/render/theme/shared"
	tabler "github.com/dracory/dashboard/render/theme/tabler"
)

// init registers all available themes
func init() {
	// Register Tabler theme
	theme.RegisterThemeInitializer(func() shared.Theme {
		return tabler.NewTablerTheme()
	})

	// Register Bootstrap 5 theme
	theme.RegisterThemeInitializer(func() shared.Theme {
		return bootstrap.NewBootstrapTheme()
	})

	// Register AdminLTE theme
	theme.RegisterThemeInitializer(func() shared.Theme {
		return adminlte.NewAdminLTETheme()
	})

	// Initialize all registered themes
	theme.InitializeRegisteredThemes()
}
