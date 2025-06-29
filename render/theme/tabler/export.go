package tabler

import (
	"github.com/dracory/dashboard/render/theme"
	"github.com/dracory/dashboard/render/theme/shared"
)

// NewTablerTheme creates a new instance of the Tabler theme
func NewTablerTheme() shared.Theme {
	return &TablerTheme{}
}

func init() {
	theme.RegisterThemeInitializer(func() shared.Theme {
		return NewTablerTheme()
	})
}
