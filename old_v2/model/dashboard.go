package model

// DashboardRenderer defines the interface for dashboard rendering data
type DashboardRenderer interface {
	// Content access
	GetContent() string
	GetFaviconURL() string
	GetLogoImageURL() string
	GetLogoRawHtml() string
	GetLogoRedirectURL() string
	GetThemeName() string

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

// MenuItem represents a menu item in the dashboard
type MenuItem struct {
	// The unique ID of the menu item
	ID string

	// The icon of the menu item
	Icon string

	// The text of the menu item
	Text string

	// The URL of the menu item
	URL string

	// Whether the menu item is active
	Active bool

	// The submenu items
	SubMenu []MenuItem

	// Optional. Whether to open in new window
	NewWindow bool

	// Optional. The badge text
	BadgeText string

	// Optional. The badge class
	BadgeClass string

	// Optional. The onclick JavaScript
	OnClick string
}

// User represents a user in the dashboard
type User struct {
	// The unique ID of the user
	ID string

	// The name of the user
	Name string

	// The email of the user
	Email string

	// The avatar URL of the user
	AvatarURL string
}
