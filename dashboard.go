package dashboard

import (
	"github.com/dracory/dashboard/model"
)

const MENU_TYPE_MODAL = "modal"
const MENU_TYPE_OFFCANVAS = "offcanvas"

// MenuItem is an alias for model.MenuItem for backward compatibility
type MenuItem = model.MenuItem

// User is an alias for model.User for backward compatibility
type User = model.User

// Dashboard represents the main dashboard structure
type Dashboard struct {
	// The menu items for the main menu
	menuItems []MenuItem

	// Whether to show the menu text (default false)
	menuShowText bool

	// Optional. The type of the main menu (see MENU_TYPE_* constants)
	menuType string

	// The currently logged in user (default nil)
	user User

	// The menu items for the user dropdown
	userMenu []MenuItem

	// Optional. Menu for quick access to various pages
	quickAccessMenu []MenuItem

	// Optional. The background color for the navbar, light or dark (default)
	navbarBackgroundColorMode string

	// Optional. The background color for the navbar (default dark)
	navbarBackgroundColor string

	// Optional. The text color for the navbar (default light)
	navbarTextColor string

	// Optional. The URL of the login page to use (if user is not provided)
	loginURL string

	// Optional. The URL of the register page to use (if user is not provided)
	registerURL string

	// Optional. The URL of the logo image
	logoImageURL string

	// Optional. Raw HTML of the logo, if set will be used instead of logoImageURL
	logoRawHtml string

	// Optional. The redirect URL of the logo image
	logoRedirectURL string

	// Optional. The favicon URL
	faviconURL string

	// Optional. The theme name
	themeName string

	// Optional. The content to display in the dashboard
	content string
}

// New creates a new dashboard instance
func New() *Dashboard {
	return &Dashboard{
		menuType:                  MENU_TYPE_OFFCANVAS,
		navbarBackgroundColorMode: "dark",
		navbarTextColor:           "light",
		menuShowText:              true,
	}
}

// SetContent sets the content of the dashboard
func (d *Dashboard) SetContent(content string) *Dashboard {
	d.content = content
	return d
}

// SetFaviconURL sets the favicon URL
func (d *Dashboard) SetFaviconURL(faviconURL string) *Dashboard {
	d.faviconURL = faviconURL
	return d
}

// SetLogoImageURL sets the logo image URL
func (d *Dashboard) SetLogoImageURL(logoImageURL string) *Dashboard {
	d.logoImageURL = logoImageURL
	return d
}

// SetLogoRawHtml sets the logo raw HTML
func (d *Dashboard) SetLogoRawHtml(logoRawHtml string) *Dashboard {
	d.logoRawHtml = logoRawHtml
	return d
}

// SetLogoRedirectURL sets the logo redirect URL
func (d *Dashboard) SetLogoRedirectURL(logoRedirectURL string) *Dashboard {
	d.logoRedirectURL = logoRedirectURL
	return d
}

// SetMenuItems sets the menu items
func (d *Dashboard) SetMenuItems(menuItems []MenuItem) *Dashboard {
	d.menuItems = menuItems
	return d
}

// SetMenuShowText sets whether to show menu text
func (d *Dashboard) SetMenuShowText(menuShowText bool) *Dashboard {
	d.menuShowText = menuShowText
	return d
}

// SetMenuType sets the menu type
func (d *Dashboard) SetMenuType(menuType string) *Dashboard {
	d.menuType = menuType
	return d
}

// SetNavbarBackgroundColor sets the navbar background color
func (d *Dashboard) SetNavbarBackgroundColor(navbarBackgroundColor string) *Dashboard {
	d.navbarBackgroundColor = navbarBackgroundColor
	return d
}

// SetNavbarBackgroundColorMode sets the navbar background color mode
func (d *Dashboard) SetNavbarBackgroundColorMode(navbarBackgroundColorMode string) *Dashboard {
	d.navbarBackgroundColorMode = navbarBackgroundColorMode
	return d
}

// SetNavbarTextColor sets the navbar text color
func (d *Dashboard) SetNavbarTextColor(navbarTextColor string) *Dashboard {
	d.navbarTextColor = navbarTextColor
	return d
}

// SetQuickAccessMenu sets the quick access menu
func (d *Dashboard) SetQuickAccessMenu(quickAccessMenu []MenuItem) *Dashboard {
	d.quickAccessMenu = quickAccessMenu
	return d
}

// SetUser sets the user
func (d *Dashboard) SetUser(user User) *Dashboard {
	d.user = user
	return d
}

// SetUserMenu sets the user menu
func (d *Dashboard) SetUserMenu(userMenu []MenuItem) *Dashboard {
	d.userMenu = userMenu
	return d
}

// SetThemeName sets the theme name
func (d *Dashboard) SetThemeName(themeName string) *Dashboard {
	d.themeName = themeName
	return d
}
