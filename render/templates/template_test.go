package templates_test

import (
	"fmt"
	"testing"

	"github.com/dracory/dashboard/render/templates/adminlte"
	"github.com/dracory/dashboard/render/templates/bootstrap"
	"github.com/dracory/dashboard/render/templates/shared"
	"github.com/dracory/dashboard/render/templates/tabler"
	"github.com/dracory/omni"
)

func TestThemeRendering(t *testing.T) {
	tests := []struct {
		name  string
		theme shared.Template
	}{
		{
			name:  "Bootstrap",
			theme: bootstrap.NewBootstrapTemplate(),
		},
		{
			name:  "AdminLTE",
			theme: adminlte.NewAdminLTETemplate(),
		},
		{
			name:  "Tabler",
			theme: tabler.NewTablerTemplate(),
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
