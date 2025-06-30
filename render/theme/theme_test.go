package theme_test

import (
	"fmt"
	"testing"

	"github.com/dracory/dashboard/render/theme/shared"
	"github.com/dracory/dashboard/render/theme/adminlte"
	"github.com/dracory/dashboard/render/theme/bootstrap"
	"github.com/dracory/dashboard/render/theme/tabler"
	"github.com/dracory/omni"
)

func TestThemeRendering(t *testing.T) {
	tests := []struct {
		name  string
		theme shared.Theme
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
			testCases := []struct {
				name     string
				testFunc func() error
			}{
				{
					name: "container",
					testFunc: func() error {
						atom := omni.NewAtom("container")
						if atomPtr, ok := atom.(*omni.Atom); ok {
							_, err := tt.theme.RenderAtom(atomPtr)
							return err
						}
						return fmt.Errorf("failed to convert omni.AtomInterface to *omni.Atom")
					},
				},
				{
					name: "header",
					testFunc: func() error {
						atom := omni.NewAtom("header")
						if atomPtr, ok := atom.(*omni.Atom); ok {
							_, err := tt.theme.RenderAtom(atomPtr)
							return err
						}
						return fmt.Errorf("failed to convert omni.AtomInterface to *omni.Atom")
					},
				},
				{
					name: "footer",
					testFunc: func() error {
						atom := omni.NewAtom("footer")
						if atomPtr, ok := atom.(*omni.Atom); ok {
							_, err := tt.theme.RenderAtom(atomPtr)
							return err
						}
						return fmt.Errorf("failed to convert omni.AtomInterface to *omni.Atom")
					},
				},
				{
					name: "menu",
					testFunc: func() error {
						atom := omni.NewAtom("menu")
						if atomPtr, ok := atom.(*omni.Atom); ok {
							_, err := tt.theme.RenderAtom(atomPtr)
							return err
						}
						return fmt.Errorf("failed to convert omni.AtomInterface to *omni.Atom")
					},
				},
				{
					name: "dashboard",
					testFunc: func() error {
						atom := omni.NewAtom("dashboard")
						if atomPtr, ok := atom.(*omni.Atom); ok {
							_, err := tt.theme.RenderDashboard(atomPtr)
							return err
						}
						return fmt.Errorf("failed to convert omni.AtomInterface to *omni.Atom")
					},
				},
			}

			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					if err := tc.testFunc(); err != nil {
						t.Fatalf("%s: Failed to render %s: %v", tt.name, tc.name, err)
					}
				})
			}
		})
	}
}
