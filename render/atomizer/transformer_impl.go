package atomizer

import (
	"fmt"

	"github.com/dracory/dashboard/model"
	"github.com/dracory/omni"
)

// TransformDashboard converts a dashboard model to an Omni atom tree
func (t *DefaultTransformer) TransformDashboard(dashboard model.DashboardRenderer) (omni.AtomInterface, error) {
	dashboardAtom := NewAtom(AtomTypeDashboard)

	// Add header
	header, err := t.TransformHeader(dashboard)
	if err != nil {
		return nil, fmt.Errorf("failed to transform header: %w", err)
	}
	dashboardAtom.ChildAdd(header)

	// Add content
	content := NewAtom(AtomTypeContent,
		WithText(dashboard.GetContent()),
	)
	dashboardAtom.ChildAdd(content)

	// Add footer
	footer, err := t.TransformFooter(dashboard)
	if err != nil {
		return nil, fmt.Errorf("failed to transform footer: %w", err)
	}
	dashboardAtom.ChildAdd(footer)

	return dashboardAtom, nil
}

// TransformHeader converts the dashboard header to an Omni atom
func (t *DefaultTransformer) TransformHeader(dashboard model.DashboardRenderer) (omni.AtomInterface, error) {
	header := NewAtom(AtomTypeHeader)

	// Add logo if available
	if logoURL := dashboard.GetLogoImageURL(); logoURL != "" {
		logo := NewAtom(AtomTypeLink,
			omni.WithProperties(map[string]string{
				PropText: "Logo",
				PropHref: dashboard.GetLogoRedirectURL(),
			}),
		)

		img := NewAtom(AtomTypeImage,
			omni.WithProperties(map[string]string{
				PropSrc: logoURL,
				PropAlt: "Logo",
			}),
		)
		logo.ChildAdd(img)
		header.ChildAdd(logo)
	}

	// Add main navigation menu
	if menuItems := dashboard.GetMenuItems(); len(menuItems) > 0 {
		menu, err := t.TransformMenu(menuItems)
		if err != nil {
			return nil, fmt.Errorf("failed to transform main menu: %w", err)
		}
		header.ChildAdd(menu)
	}

	// Add user menu if user is logged in
	user := dashboard.GetUser()
	if user.ID != "" { // Check if user has an ID (logged in)
		userMenu, err := t.TransformUserMenu(user, dashboard.GetUserMenu())
		if err != nil {
			return nil, fmt.Errorf("failed to transform user menu: %w", err)
		}
		header.ChildAdd(userMenu)
	}

	return header, nil
}

// TransformFooter converts the dashboard footer to an Omni atom
func (t *DefaultTransformer) TransformFooter(dashboard model.DashboardRenderer) (omni.AtomInterface, error) {
	footer := NewAtom(AtomTypeFooter)
	// Add footer content here
	return footer, nil
}

// TransformMenu converts a list of menu items to an Omni menu atom
func (t *DefaultTransformer) TransformMenu(menuItems []model.MenuItem) (omni.AtomInterface, error) {
	menu := NewAtom(AtomTypeMenu)

	for _, item := range menuItems {
		menuItem, err := t.transformMenuItem(item)
		if err != nil {
			return nil, fmt.Errorf("failed to transform menu item: %w", err)
		}
		menu.ChildAdd(menuItem)
	}

	return menu, nil
}

// transformMenuItem converts a single menu item to an Omni atom
func (t *DefaultTransformer) transformMenuItem(item model.MenuItem) (omni.AtomInterface, error) {
	// Create menu item with type
	menuItem := NewAtom(AtomTypeMenuItem)

	// Set properties using omni.NewProperty
	if item.Text != "" {
		menuItem.Set(PropText, item.Text)
	}

	// Add URL property if present
	if item.URL != "" {
		menuItem.Set(PropHref, item.URL)
	}

	// Add active state if true
	if item.Active {
		menuItem.Set(PropActive, "true")
	}

	// Add submenu items if any
	if len(item.Children) > 0 {
		submenu, err := t.TransformMenu(item.Children)
		if err != nil {
			return nil, fmt.Errorf("failed to transform submenu: %w", err)
		}
		menuItem.ChildAdd(submenu)
	}

	return menuItem, nil
}

// TransformUserMenu converts the user menu to an Omni atom
func (t *DefaultTransformer) TransformUserMenu(user model.User, menuItems []model.MenuItem) (omni.AtomInterface, error) {
	userMenu := NewAtom("user_menu")

	// Add user info
	userInfo := NewAtom("user_info")
	userInfo.Set(PropText, user.Name)
	userMenu.ChildAdd(userInfo)

	// Add user menu items
	if len(menuItems) > 0 {
		menu, err := t.TransformMenu(menuItems)
		if err != nil {
			return nil, fmt.Errorf("failed to transform user menu items: %w", err)
		}
		userMenu.ChildAdd(menu)
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
