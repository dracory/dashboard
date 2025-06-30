package dashboard

import "net/http"

// TemplateNameContextKey is a key for storing template name in context
type TemplateNameContextKey struct{}

// TEMPLATE_COOKIE_KEY is the name of the cookie that stores the template
var TEMPLATE_COOKIE_KEY = "template"

// Config holds configuration for the dashboard
type Config struct {
	// The content to display in the dashboard
	Content string
	
	// The favicon URL
	FaviconURL string
	
	// The HTTP request (for template detection)
	HTTPRequest *http.Request
	
	// The URL of the logo image
	LogoURL string

	// Should the menu text be shown
	MenuShowText bool

	// The menu items to display in the main menu
	MenuItems []MenuItem

	// MenuType can be MENU_TYPE_OFFCANVAS or MENU_TYPE_MODAL
	MenuType string

	// Optional. The URL of the logo image
	LogoImageURL string

	// Optional. Raw HTML of the logo, if set will be used instead of logoImageURL
	LogoRawHtml string

	// Optional. The redirect URL of the logo image
	LogoRedirectURL string

	// Optional The background color for the navbar: light, dark (default), primary, secondary, success, warning, info, danger
	NavbarBackgroundColorMode string

	// Optional. The background color for the navbar (default none)
	NavbarBackgroundColor string

	// Optional. The text color for the navbar (default light)
	NavbarTextColor string

	// Optional. The template name to use for the dashboard
	// Defaults to "tabler"
	TemplateName string

	// Optional. The URL of the login page to use (if user is not provided)
	LoginURL string

	// Optional. The URL of the register page to use (if user is not provided)
	RegisterURL string

	// Optional. Menu for Quick Access
	QuickAccessMenu []MenuItem
	
	// Optional. The currently logged in user
	User User
	
	// Optional. The menu items for the user dropdown
	UserMenu []MenuItem
}

// NewFromConfig creates a new dashboard from a config
func NewFromConfig(config Config) *Dashboard {
	dashboard := New()
	
	if config.Content != "" {
		dashboard.SetContent(config.Content)
	}
	
	if config.FaviconURL != "" {
		dashboard.SetFaviconURL(config.FaviconURL)
	}
	
	if config.LogoImageURL != "" {
		dashboard.SetLogoImageURL(config.LogoImageURL)
	}
	
	if config.LogoRawHtml != "" {
		dashboard.SetLogoRawHtml(config.LogoRawHtml)
	}
	
	if config.LogoRedirectURL != "" {
		dashboard.SetLogoRedirectURL(config.LogoRedirectURL)
	}
	
	if config.MenuType != "" {
		dashboard.SetMenuType(config.MenuType)
	}
	
	if len(config.MenuItems) > 0 {
		dashboard.SetMenuItems(config.MenuItems)
	}
	
	dashboard.SetMenuShowText(config.MenuShowText)
	
	if config.NavbarBackgroundColor != "" {
		dashboard.SetNavbarBackgroundColor(config.NavbarBackgroundColor)
	}
	
	if config.NavbarBackgroundColorMode != "" {
		dashboard.SetNavbarBackgroundColorMode(config.NavbarBackgroundColorMode)
	}
	
	if config.NavbarTextColor != "" {
		dashboard.SetNavbarTextColor(config.NavbarTextColor)
	}
	
	if config.LoginURL != "" {
		dashboard.loginURL = config.LoginURL
	}
	
	if config.RegisterURL != "" {
		dashboard.registerURL = config.RegisterURL
	}
	
	if len(config.QuickAccessMenu) > 0 {
		dashboard.SetQuickAccessMenu(config.QuickAccessMenu)
	}
	
	if config.TemplateName != "" {
		dashboard.SetTemplateName(config.TemplateName)
	}
	
	if config.User.ID != "" {
		dashboard.SetUser(config.User)
	}
	
	if len(config.UserMenu) > 0 {
		dashboard.SetUserMenu(config.UserMenu)
	}
	
	// Detect template from cookie if HTTP request is provided
	if config.HTTPRequest != nil {
		cookie, err := config.HTTPRequest.Cookie(TEMPLATE_COOKIE_KEY)
		if err == nil && cookie.Value != "" {
			dashboard.SetTemplateName(cookie.Value)
		}
	}
	
	return dashboard
}
