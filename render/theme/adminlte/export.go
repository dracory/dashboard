package adminlte

import (
	"github.com/dracory/dashboard/render/theme"
	"github.com/dracory/dashboard/render/theme/shared"
)

// NewAdminLTETheme creates a new instance of the AdminLTE theme
func NewAdminLTETheme() shared.Theme {
	return &AdminLTETheme{}
}

func init() {
	theme.RegisterThemeInitializer(func() shared.Theme {
		return NewAdminLTETheme()
	})
}
