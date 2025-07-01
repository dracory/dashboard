package types

// MenuType constants for dashboard menu types
const (
	MENU_TYPE_MODAL     = "modal"
	MENU_TYPE_OFFCANVAS = "offcanvas"
)

// BreadcrumbItem represents a single item in the breadcrumb navigation
type BreadcrumbItem struct {
	Title string // Display text for the breadcrumb item
	URL   string // URL for the breadcrumb link (empty for current/last item)
}

// Modal represents a Bootstrap/Tabler modal dialog
type Modal struct {
	ID          string // Unique ID for the modal
	Title       string // Modal title
	Content     string // Modal content (HTML)
	Size        string // Modal size (sm, lg, xl, or empty for default)
	Footer      string // Modal footer content (HTML)
	CloseButton bool   // Whether to show the close button (default: true)
}

// Alert represents a notification message to display to the user
type Alert struct {
	Type    string // Alert type (e.g., "success", "danger", "warning", "info")
	Message string // The message to display
}

// Action represents an action button in the header
type Action struct {
	ID      string // Optional ID for the button
	Title   string // Button text
	Icon    string // Icon name (without the 'ti ti-' prefix)
	Primary bool   // Whether this is a primary action
	OnClick string // JavaScript to execute on click
}

type DashboardInterface interface {
	// GetContent returns the content of the webpage
	GetContent() string
	// SetContent sets the content of the webpage
	SetContent(content string)

	// GetSubtitle returns the subtitle of the webpage
	GetSubtitle() string
	// SetSubtitle sets the subtitle of the webpage
	SetSubtitle(subtitle string)

	// GetFaviconURL returns the favicon URL of the dashboard
	GetFaviconURL() string
	// SetFaviconURL sets the favicon URL of the dashboard
	SetFaviconURL(faviconURL string)

	// GetLogoImageURL returns the logo image URL of the dashboard
	GetLogoImageURL() string
	// SetLogoImageURL sets the logo image URL of the dashboard
	SetLogoImageURL(logoImageURL string)

	// GetLogoRawHtml returns the logo raw HTML of the dashboard
	GetLogoRawHtml() string
	// SetLogoRawHtml sets the logo raw HTML of the dashboard
	SetLogoRawHtml(logoRawHtml string)

	// GetLogoRedirectURL returns the logo redirect URL of the dashboard
	GetLogoRedirectURL() string
	// SetLogoRedirectURL sets the logo redirect URL of the dashboard
	SetLogoRedirectURL(logoRedirectURL string)

	// GetMenuMainItems returns the menu items for the main menu
	GetMenuMainItems() []MenuItem
	// SetMenuMainItems sets the menu items for the main menu
	SetMenuMainItems(menuItems []MenuItem)

	// GetMenuUserItems returns the menu items for the user menu
	GetMenuUserItems() []MenuItem
	// SetMenuUserItems sets the menu items for the user menu
	SetMenuUserItems(menuItems []MenuItem)

	// GetMenuQuickAccessItems returns the menu items for the quick access menu
	GetMenuQuickAccessItems() []MenuItem
	// SetMenuQuickAccessItems sets the menu items for the quick access menu
	SetMenuQuickAccessItems(menuItems []MenuItem)

	// GetTitle returns the title of the webpage
	GetTitle() string
	// SetTitle sets the title of the webpage
	SetTitle(title string)

	// GetTemplate returns the template of the dashboard
	GetTemplate() string
	// SetTemplate sets the template of the dashboard
	SetTemplate(template string)

	// GetUser returns the user of the dashboard
	GetUser() *User
	// SetUser sets the user of the dashboard
	SetUser(user User)

	// GetRedirectTime returns the redirect time of the dashboard
	GetRedirectTime() string
	// SetRedirectTime sets the redirect time of the dashboard
	SetRedirectTime(redirectTime string)

	// GetRedirectUrl returns the redirect URL of the dashboard
	GetRedirectUrl() string
	// SetRedirectUrl sets the redirect URL of the dashboard
	SetRedirectUrl(redirectUrl string)

	// GetScripts returns the scripts of the dashboard
	GetScripts() []string
	// SetScripts sets the scripts of the dashboard
	SetScripts(scripts []string)

	// GetScriptURLs returns the script URLs of the dashboard
	GetScriptURLs() []string
	// SetScriptURLs sets the script URLs of the dashboard
	SetScriptURLs(scriptURLs []string)

	// GetStyles returns the styles of the dashboard
	GetStyles() []string
	// SetStyles sets the styles of the dashboard
	SetStyles(styles []string)

	// GetStyleURLs returns the style URLs of the dashboard
	GetStyleURLs() []string
	// SetStyleURLs sets the style URLs of the dashboard
	SetStyleURLs(styleURLs []string)

	// Navbar theming methods
	GetNavbarBackgroundColorMode() string
	SetNavbarBackgroundColorMode(mode string)
	GetNavbarBackgroundColor() string
	SetNavbarBackgroundColor(color string)
	GetNavbarTextColor() string
	SetNavbarTextColor(color string)

	// Navbar utility methods
	NavbarHasBackgroundThemeClass() bool
	NavbarBackgroundThemeClass() string
	NavbarButtonThemeClass() string

	// Navbar dropdown methods
	NavbarDropdownQuickAccess(iconStyle string) string
	NavbarDropdownThemeSwitch() string
	NavbarDropdownUser(iconStyle string) string

	// Login/register URLs
	GetLoginURL() string
	SetLoginURL(url string)
	GetRegisterURL() string
	SetRegisterURL(url string)

	// Theme methods
	IsThemeDark() bool
	GetTheme() string
	SetTheme(theme string)
	GetThemeHandlerUrl() string
	SetThemeHandlerUrl(url string)
	GetThemesRestrict() map[string]string
	SetThemesRestrict(themes map[string]string)

	// UI Configuration
	GetMenuShowText() bool
	SetMenuShowText(showText bool)

	// Menu Type Configuration
	GetMenuType() string
	SetMenuType(menuType string)

	// Navbar background
	GetNavbarBackground() (string, bool)

	// Sidebar state
	GetSidebarCollapsed() bool
	SetSidebarCollapsed(collapsed bool)

	// Breadcrumb
	GetBreadcrumb() []BreadcrumbItem
	SetBreadcrumb(items []BreadcrumbItem)

	// Actions
	GetActions() []Action
	SetActions(actions []Action)

	// Alerts
	GetAlerts() []Alert
	AddAlert(alert Alert)
	ClearAlerts()

	// Modals
	GetModals() []Modal
	AddModal(modal Modal)
	ClearModals()

	ToHTML() string
}
