package atomizer

import (
	"fmt"

	"github.com/dracory/dashboard/model"
	"github.com/dracory/omni"
)

// TransformDashboard converts a dashboard model to an Omni atom tree
func (t *defaultTransformer) TransformDashboard(dashboard model.DashboardRenderer) (*omni.Atom, error) {
	dashboardAtom := NewAtom(AtomTypeDashboard)

	// Add header
	header, err := t.TransformHeader(dashboard)
	if err != nil {
		return nil, fmt.Errorf("failed to transform header: %w", err)
	}
	dashboardAtom.AddChild(header)

	// Add content
	content := NewAtom(AtomTypeContent,
		WithText(dashboard.GetContent()),
	)
	dashboardAtom.AddChild(content)

	// Add footer
	footer, err := t.TransformFooter(dashboard)
	if err != nil {
		return nil, fmt.Errorf("failed to transform footer: %w", err)
	}
	dashboardAtom.AddChild(footer)

	return dashboardAtom, nil
}

// TransformHeader converts the dashboard header to an Omni atom
func (t *defaultTransformer) TransformHeader(dashboard model.DashboardRenderer) (*omni.Atom, error) {
	header := NewAtom(AtomTypeHeader)

	// Add logo if available
	if logoURL := dashboard.GetLogoImageURL(); logoURL != "" {
		logo := NewAtom(AtomTypeLink,
			omni.WithProperties(
				omni.NewProperty(PropHref, dashboard.GetLogoRedirectURL()),
			),
		)

		img := NewAtom(AtomTypeImage,
			omni.WithProperties(
				omni.NewProperty(PropSrc, logoURL),
				omni.NewProperty(PropAlt, "Logo"),
			),
		)
		logo.AddChild(img)
		header.AddChild(logo)
	}

	// Add main navigation menu
	if menuItems := dashboard.GetMenuItems(); len(menuItems) > 0 {
		menu, err := t.TransformMenu(menuItems)
		if err != nil {
			return nil, fmt.Errorf("failed to transform main menu: %w", err)
		}
		header.AddChild(menu)
	}

	// Add user menu if user is logged in
	user := dashboard.GetUser()
	if user.ID != "" { // Check if user has an ID (logged in)
		userMenu, err := t.TransformUserMenu(user, dashboard.GetUserMenu())
		if err != nil {
			return nil, fmt.Errorf("failed to transform user menu: %w", err)
		}
		header.AddChild(userMenu)
	}

	return header, nil
}

// TransformFooter converts the dashboard footer to an Omni atom
func (t *defaultTransformer) TransformFooter(dashboard model.DashboardRenderer) (*omni.Atom, error) {
	footer := NewAtom(AtomTypeFooter)
	// Add footer content here
	return footer, nil
}

// TransformMenu converts a list of menu items to an Omni menu atom
func (t *defaultTransformer) TransformMenu(menuItems []model.MenuItem) (*omni.Atom, error) {
	menu := NewAtom(AtomTypeMenu)

	for _, item := range menuItems {
		menuItem, err := t.transformMenuItem(item)
		if err != nil {
			return nil, fmt.Errorf("failed to transform menu item: %w", err)
		}
		menu.AddChild(menuItem)
	}

	return menu, nil
}

// transformMenuItem converts a single menu item to an Omni atom
func (t *defaultTransformer) transformMenuItem(item model.MenuItem) (*omni.Atom, error) {
	// Create menu item with type
	menuItem := NewAtom(AtomTypeMenuItem)

	// Set properties using omni.NewProperty
	if item.Text != "" {
		menuItem.SetProperty(omni.NewProperty(PropText, item.Text))
	}

	// Add URL property if present
	if item.URL != "" {
		menuItem.SetProperty(omni.NewProperty(PropHref, item.URL))
	}

	// Add active state if true
	if item.Active {
		menuItem.SetProperty(omni.NewProperty(PropActive, "true"))
	}

	// Add submenu items if any
	if len(item.SubMenu) > 0 {
		submenu, err := t.TransformMenu(item.SubMenu)
		if err != nil {
			return nil, fmt.Errorf("failed to transform submenu: %w", err)
		}
		menuItem.AddChild(submenu)
	}

	return menuItem, nil
}

// TransformUserMenu converts the user menu to an Omni atom
func (t *defaultTransformer) TransformUserMenu(user model.User, menuItems []model.MenuItem) (*omni.Atom, error) {
	userMenu := NewAtom("user_menu")

	// Add user info
	userInfo := NewAtom("user_info")
	userInfo.SetProperty(omni.NewProperty(PropText, user.Name))
	userMenu.AddChild(userInfo)

	// Add user menu items
	if len(menuItems) > 0 {
		menu, err := t.TransformMenu(menuItems)
		if err != nil {
			return nil, fmt.Errorf("failed to transform user menu items: %w", err)
		}
		userMenu.AddChild(menu)
	}

	return userMenu, nil
}

// Helper function to set a property on an atom
// func WithProperty(key string, value interface{}) omni.AtomOption {
// 	return func(a *omni.Atom) {
// 		// Convert value to string using cast.ToString
// 		strVal := cast.ToString(value)
// 		a.SetProperty(omni.NewProperty(key, strVal))
// 	}
// }
