package theme_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"dracorys.io/dashboard/render/theme"
	"dracorys.io/dashboard/render/theme/adminlte"
	"dracorys.io/dashboard/render/theme/bootstrap"
	"dracorys.io/dashboard/render/theme/tabler"
	"dracorys.io/dashboard/render/omni"
)

func TestThemeRendering(t *testing.T) {
	tests := []struct {
		name  string
		theme theme.Theme
	}{
		{
			name:  "Bootstrap",
			theme: bootstrap.NewBootstrapTheme(),
		},
		{
			name:  "AdminLTE",
			theme: adminlte.NewAdminLTETheme(),
		},
		{
			name:  "Tabler",
			theme: tabler.NewTablerTheme(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test container rendering
			container := omni.NewAtom("container")
			_, err := tt.theme.RenderAtom(container)
			assert.NoError(t, err, "%s: Failed to render container", tt.name)

			// Test header rendering
			header := omni.NewAtom("header")
			_, err = tt.theme.RenderAtom(header)
			assert.NoError(t, err, "%s: Failed to render header", tt.name)

			// Test footer rendering
			footer := omni.NewAtom("footer")
			_, err = tt.theme.RenderAtom(footer)
			assert.NoError(t, err, "%s: Failed to render footer", tt.name)

			// Test menu rendering
			menu := omni.NewAtom("menu")
			_, err = tt.theme.RenderAtom(menu)
			assert.NoError(t, err, "%s: Failed to render menu", tt.name)

			// Test dashboard rendering
			dashboard := omni.NewAtom("dashboard")
			_, err = tt.theme.RenderDashboard(dashboard)
			assert.NoError(t, err, "%s: Failed to render dashboard", tt.name)
		})
	}
}
