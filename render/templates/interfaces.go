package templates

// This file contains shared interfaces and types for dashboard templates

// DashboardRenderer defines the interface for dashboard renderers
type DashboardRenderer interface {
	// Content access
	GetContent() string
	GetFaviconURL() string
	GetLogoImageURL() string
	GetLogoRawHtml() string
	GetLogoRedirectURL() string
	GetTemplateName() string

	// Menu access
	GetMenuItems() []MenuItem
	GetMenuShowText() bool
	GetQuickAccessMenu() []MenuItem

	// User access
	GetUser() User
	GetUserMenu() []MenuItem
	GetLoginURL() string
	GetRegisterURL() string

	// Navbar access
	GetNavbarBackgroundColorMode() string
	GetNavbarBackgroundColor() string
	GetNavbarTextColor() string
}

// Template interface has been moved to template.go to avoid duplication

// MenuItem represents a single item in a menu
type MenuItem struct {
	ID       string
	Title    string
	URL      string
	Icon     string
	Children []MenuItem
	Active   bool
}

// User represents a user in the system
type User struct {
	ID        string
	Name      string
	Email     string
	AvatarURL string
}
