package atomizer

import (
	"github.com/dracory/dashboard/model"
	"github.com/dracory/omni"
)

// Transformer converts dashboard model to Omni atoms
type Transformer interface {
	// TransformDashboard converts a dashboard model to an Omni atom tree
	TransformDashboard(dashboard model.DashboardRenderer) (*omni.Atom, error)

	// TransformHeader converts dashboard header to an Omni atom
	TransformHeader(dashboard model.DashboardRenderer) (*omni.Atom, error)

	// TransformFooter converts dashboard footer to an Omni atom
	TransformFooter(dashboard model.DashboardRenderer) (*omni.Atom, error)

	// TransformMenu converts a menu item to an Omni atom
	TransformMenu(menu []model.MenuItem) (*omni.Atom, error)

	// TransformUserMenu converts the user menu to an Omni atom
	TransformUserMenu(user model.User, menu []model.MenuItem) (*omni.Atom, error)
}

// NewTransformer creates a new dashboard to Omni transformer
func NewTransformer() Transformer {
	return &defaultTransformer{}
}

// defaultTransformer implements the Transformer interface
type defaultTransformer struct{}
