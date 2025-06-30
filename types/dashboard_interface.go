package types

type DashboardInterface interface {

	// SetTitle(title string) DashboardInterface
	// SetContent(content string) DashboardInterface
	// SetFaviconURL(faviconURL string) DashboardInterface
	// SetLogoImageURL(logoImageURL string) DashboardInterface
	// SetLogoRawHtml(logoRawHtml string) DashboardInterface
	// SetLogoRedirectURL(logoRedirectURL string) DashboardInterface
	// SetThemeName(themeName string) DashboardInterface
	// SetMenuItems(menuItems []MenuItem) DashboardInterface
	// SetMenuShowText(menuShowText bool) DashboardInterface
	// SetQuickAccessMenu(quickAccessMenu []MenuItem) DashboardInterface

	// GetTitle() string
	// SetTitle(title string) DashboardInterface

	GetContent() string
	SetContent(content string)

	// GetUser() User
	// SetUser(user User) DashboardInterface

	// GetUserMenu() []MenuItem

	// SetUserMenu(userMenu []MenuItem) DashboardInterface

	// SetLoginURL(loginURL string) DashboardInterface
	// SetRegisterURL(registerURL string) DashboardInterface
	// SetNavbarBackgroundColorMode(navbarBackgroundColorMode string) DashboardInterface
	// SetNavbarBackgroundColor(navbarBackgroundColor string) DashboardInterface
	// SetNavbarTextColor(navbarTextColor string) DashboardInterface
	ToHTML() string
}
