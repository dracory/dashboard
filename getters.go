package dashboard

import (
	"github.com/dracory/dashboard/model"
)

// GetUser returns the user
func (d *Dashboard) GetUser() model.User {
	return d.user
}

// GetLoginURL returns the login URL
func (d *Dashboard) GetLoginURL() string {
	return d.loginURL
}

// GetRegisterURL returns the register URL
func (d *Dashboard) GetRegisterURL() string {
	return d.registerURL
}

// GetUserMenu returns the user menu
func (d *Dashboard) GetUserMenu() []model.MenuItem {
	return d.userMenu
}

// GetMenuItems returns the menu items
func (d *Dashboard) GetMenuItems() []model.MenuItem {
	return d.menuItems
}

// GetQuickAccessMenu returns the quick access menu
func (d *Dashboard) GetQuickAccessMenu() []model.MenuItem {
	return d.quickAccessMenu
}

// GetLogoImageURL returns the logo image URL
func (d *Dashboard) GetLogoImageURL() string {
	return d.logoImageURL
}

// GetLogoRawHtml returns the logo raw HTML
func (d *Dashboard) GetLogoRawHtml() string {
	return d.logoRawHtml
}

// GetLogoRedirectURL returns the logo redirect URL
func (d *Dashboard) GetLogoRedirectURL() string {
	return d.logoRedirectURL
}

// GetFaviconURL returns the favicon URL
func (d *Dashboard) GetFaviconURL() string {
	return d.faviconURL
}

// GetThemeName returns the theme name
func (d *Dashboard) GetThemeName() string {
	return d.themeName
}

// GetContent returns the content
func (d *Dashboard) GetContent() string {
	return d.content
}

// GetMenuShowText returns whether to show menu text
func (d *Dashboard) GetMenuShowText() bool {
	return d.menuShowText
}

// GetMenuType returns the menu type
func (d *Dashboard) GetMenuType() string {
	return d.menuType
}

// GetNavbarBackgroundColor returns the navbar background color
func (d *Dashboard) GetNavbarBackgroundColor() string {
	return d.navbarBackgroundColor
}

// GetNavbarBackgroundColorMode returns the navbar background color mode
func (d *Dashboard) GetNavbarBackgroundColorMode() string {
	return d.navbarBackgroundColorMode
}

// GetNavbarTextColor returns the navbar text color
func (d *Dashboard) GetNavbarTextColor() string {
	return d.navbarTextColor
}
