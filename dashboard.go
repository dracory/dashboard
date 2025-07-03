package dashboard

import (
	"github.com/dracory/dashboard/shared"
	"github.com/dracory/dashboard/templates/adminlte"
	"github.com/dracory/dashboard/templates/bootstrap"
	"github.com/dracory/dashboard/templates/tabler"
	"github.com/dracory/dashboard/types"
	"github.com/samber/lo"
)

type dashboard struct {
	content                   string
	subtitle                  string         // page subtitle
	actions                   []types.Action // header action buttons
	alerts                    []types.Alert  // alert messages to display
	modals                    []types.Modal  // modal dialogs
	faviconURL                string
	layout                    string // layout: some templates support different layouts
	logoImageURL              string
	logoRawHtml               string
	logoRedirectURL           string
	menuMainItems             []types.MenuItem
	menuShowText              bool   // controls whether to show text in menu items
	menuType                  string // modal or offcanvas
	menuUserItems             []types.MenuItem
	menuQuickAccessItems      []types.MenuItem
	navbarBackgroundColorMode string
	navbarBackgroundColor     string
	navbarTextColor           string
	redirectTime              string                 // redirect time (if any, in seconds)
	redirectUrl               string                 // redirect URL (if any)
	scripts                   []string               // custom scripts defined by the user
	scriptURLs                []string               // custom script URLs defined by the user
	sidebarCollapsed          bool                   // whether the sidebar is collapsed
	breadcrumb                []types.BreadcrumbItem // breadcrumb navigation items
	styles                    []string               // custom styles defined by the user
	styleURLs                 []string               // custom style URLs defined by the user
	theme                     string                 // color mode: default, dark, light
	themesRestrict            map[string]string      // restricted theme options
	themeHandlerUrl           string                 // URL for theme handler
	template                  string                 // bootstrap (default), adminlte, tabler
	title                     string                 // title of the webpage
	user                      *types.User            // user object (if any)
	loginURL                  string                 // login URL
	registerURL               string                 // register URL
}

var _ types.DashboardInterface = (*dashboard)(nil)

// ToHTML returns the HTML of the dashboard
func (d *dashboard) ToHTML() string {
	templateName, template := d.findTemplate()

	if template == nil {
		return "Template not found: " + templateName
	}

	return template.ToHTML(d)
}

// findTemplate returns the template name and template instance
// If template is not found, returns the default template
func (d *dashboard) findTemplate() (templateName string, template types.TemplateInterface) {
	templateName = d.GetTemplate()
	if templateName == "" {
		templateName = shared.TEMPLATE_DEFAULT
	}

	if templateName == shared.TEMPLATE_BOOTSTRAP {
		return templateName, &bootstrap.Template{}
	}

	if templateName == shared.TEMPLATE_TABLER {
		return templateName, &tabler.Template{}
	}

	if templateName == shared.TEMPLATE_ADMINLTE {
		return templateName, &adminlte.Template{}
	}

	return templateName, &bootstrap.Template{}
}

// ============================================================================
// == Sidebar State Methods
// ============================================================================

// GetSidebarCollapsed returns whether the sidebar is collapsed
func (d *dashboard) GetSidebarCollapsed() bool {
	return d.sidebarCollapsed
}

// SetSidebarCollapsed sets whether the sidebar is collapsed
func (d *dashboard) SetSidebarCollapsed(collapsed bool) {
	d.sidebarCollapsed = collapsed
}

// ============================================================================
// == Breadcrumb Methods
// ============================================================================

// GetBreadcrumb returns the breadcrumb navigation items
func (d *dashboard) GetBreadcrumb() []types.BreadcrumbItem {
	if d.breadcrumb == nil {
		return []types.BreadcrumbItem{}
	}
	return d.breadcrumb
}

// SetBreadcrumb sets the breadcrumb navigation items
func (d *dashboard) SetBreadcrumb(items []types.BreadcrumbItem) {
	d.breadcrumb = items
}

// ============================================================================
// == Subtitle Methods
// ============================================================================

// GetSubtitle returns the page subtitle
func (d *dashboard) GetSubtitle() string {
	return d.subtitle
}

// SetSubtitle sets the page subtitle
func (d *dashboard) SetSubtitle(subtitle string) {
	d.subtitle = subtitle
}

// ============================================================================
// == Action Methods
// ============================================================================

// GetActions returns the list of action buttons for the header
func (d *dashboard) GetActions() []types.Action {
	if d.actions == nil {
		return []types.Action{}
	}
	return d.actions
}

// SetActions sets the list of action buttons for the header
func (d *dashboard) SetActions(actions []types.Action) {
	d.actions = actions
}

// ============================================================================
// == Alert Methods
// ============================================================================

// GetAlerts returns the list of alerts to display
func (d *dashboard) GetAlerts() []types.Alert {
	if d.alerts == nil {
		return []types.Alert{}
	}
	return d.alerts
}

// AddAlert adds an alert to be displayed
func (d *dashboard) AddAlert(alert types.Alert) {
	if d.alerts == nil {
		d.alerts = []types.Alert{}
	}
	d.alerts = append(d.alerts, alert)
}

// ClearAlerts removes all alerts
func (d *dashboard) ClearAlerts() {
	d.alerts = []types.Alert{}
}

// ============================================================================
// == Modal Methods
// ============================================================================

// GetModals returns the list of modals to display
func (d *dashboard) GetModals() []types.Modal {
	if d.modals == nil {
		return []types.Modal{}
	}
	return d.modals
}

// AddModal adds a modal dialog to be displayed
func (d *dashboard) AddModal(modal types.Modal) {
	if d.modals == nil {
		d.modals = []types.Modal{}
	}
	d.modals = append(d.modals, modal)
}

// ClearModals removes all modals
func (d *dashboard) ClearModals() {
	d.modals = []types.Modal{}
}

// ============================================================================
// == Getters and Setters
// ============================================================================

// GetContent returns the content of the webpage
func (d *dashboard) GetContent() string {
	return d.content
}

// SetContent sets the content of the webpage
func (d *dashboard) SetContent(content string) {
	d.content = content
}

// GetFaviconURL returns the favicon URL of the webpage
func (d *dashboard) GetFaviconURL() string {
	return d.faviconURL
}

// SetFaviconURL sets the favicon URL of the webpage
func (d *dashboard) SetFaviconURL(faviconURL string) {
	d.faviconURL = faviconURL
}

// GetLogoImageURL returns the logo image URL of the webpage
func (d *dashboard) GetLogoImageURL() string {
	return d.logoImageURL
}

// SetLogoImageURL sets the logo image URL of the webpage
func (d *dashboard) SetLogoImageURL(logoImageURL string) {
	d.logoImageURL = logoImageURL
}

// GetLogoRawHtml returns the logo raw HTML of the webpage
func (d *dashboard) GetLogoRawHtml() string {
	return d.logoRawHtml
}

// SetLogoRawHtml sets the logo raw HTML of the webpage
func (d *dashboard) SetLogoRawHtml(logoRawHtml string) {
	d.logoRawHtml = logoRawHtml
}

// GetLogoRedirectURL returns the logo redirect URL of the webpage
func (d *dashboard) GetLogoRedirectURL() string {
	return d.logoRedirectURL
}

// SetLogoRedirectURL sets the logo redirect URL of the webpage
func (d *dashboard) SetLogoRedirectURL(logoRedirectURL string) {
	d.logoRedirectURL = logoRedirectURL
}

// GetMenuMainItems returns the menu items for the main menu
func (d *dashboard) GetMenuMainItems() []types.MenuItem {
	return d.menuMainItems
}

// SetMenuMainItems sets the menu items for the main menu
func (d *dashboard) SetMenuMainItems(menuItems []types.MenuItem) {
	d.menuMainItems = menuItems
}

// GetMenuUserItems returns the menu items for the user menu
func (d *dashboard) GetMenuUserItems() []types.MenuItem {
	return d.menuUserItems
}

// SetMenuUserItems sets the menu items for the user menu
func (d *dashboard) SetMenuUserItems(menuItems []types.MenuItem) {
	d.menuUserItems = menuItems
}

// GetMenuQuickAccessItems returns the menu items for the quick access menu
func (d *dashboard) GetMenuQuickAccessItems() []types.MenuItem {
	return d.menuQuickAccessItems
}

// SetMenuQuickAccessItems sets the menu items for the quick access menu
func (d *dashboard) SetMenuQuickAccessItems(menuItems []types.MenuItem) {
	d.menuQuickAccessItems = menuItems
}

// GetMenuType returns the menu type
func (d *dashboard) GetMenuType() string {
	return d.menuType
}

// SetMenuType sets the menu type
func (d *dashboard) SetMenuType(menuType string) {
	d.menuType = menuType
}

// GetRedirectTime returns the redirect time of the dashboard
func (d *dashboard) GetRedirectTime() string {
	return d.redirectTime
}

// SetRedirectTime sets the redirect time of the dashboard
func (d *dashboard) SetRedirectTime(redirectTime string) {
	d.redirectTime = redirectTime
}

// GetRedirectUrl returns the redirect URL of the dashboard
func (d *dashboard) GetRedirectUrl() string {
	return d.redirectUrl
}

// SetRedirectUrl sets the redirect URL of the dashboard
func (d *dashboard) SetRedirectUrl(redirectUrl string) {
	d.redirectUrl = redirectUrl
}

// GetScripts returns the custom scripts of the dashboard
func (d *dashboard) GetScripts() []string {
	return d.scripts
}

// SetScripts sets the custom scripts of the dashboard
func (d *dashboard) SetScripts(scripts []string) {
	d.scripts = scripts
}

// GetScriptURLs returns the custom script URLs of the dashboard
func (d *dashboard) GetScriptURLs() []string {
	return d.scriptURLs
}

// SetScriptURLs sets the custom script URLs of the dashboard
func (d *dashboard) SetScriptURLs(scriptURLs []string) {
	d.scriptURLs = scriptURLs
}

// GetStyles returns the custom styles of the dashboard
func (d *dashboard) GetStyles() []string {
	return d.styles
}

// SetStyles sets the custom styles of the dashboard
func (d *dashboard) SetStyles(styles []string) {
	d.styles = styles
}

// GetStyleURLs returns the custom style URLs of the dashboard
func (d *dashboard) GetStyleURLs() []string {
	return d.styleURLs
}

// SetStyleURLs sets the custom style URLs of the dashboard
func (d *dashboard) SetStyleURLs(styleURLs []string) {
	d.styleURLs = styleURLs
}

// GetTemplate returns the template name, or the default template if not set
func (d *dashboard) GetTemplate() string {
	if d.template == "" {
		return shared.TEMPLATE_DEFAULT
	}
	return d.template
}

// SetTemplate sets the template name
func (d *dashboard) SetTemplate(template string) {
	d.template = template
}

// GetTitle returns the title of the webpage
func (d *dashboard) GetTitle() string {
	return d.title
}

// SetTitle sets the title of the webpage
func (d *dashboard) SetTitle(title string) {
	d.title = title
}

// GetUser returns the user
func (d *dashboard) GetUser() *types.User {
	return d.user
}

// SetUser sets the user
func (d *dashboard) SetUser(user types.User) {
	d.user = &user
}

// Navbar theming methods
func (d *dashboard) GetNavbarBackgroundColorMode() string {
	return d.navbarBackgroundColorMode
}

func (d *dashboard) SetNavbarBackgroundColorMode(mode string) {
	d.navbarBackgroundColorMode = mode
}

func (d *dashboard) GetNavbarBackgroundColor() string {
	return d.navbarBackgroundColor
}

func (d *dashboard) SetNavbarBackgroundColor(color string) {
	d.navbarBackgroundColor = color
}

func (d *dashboard) GetNavbarTextColor() string {
	return d.navbarTextColor
}

func (d *dashboard) SetNavbarTextColor(color string) {
	d.navbarTextColor = color
}

// Login/register URLs
func (d *dashboard) GetLoginURL() string {
	return d.loginURL
}

func (d *dashboard) SetLoginURL(url string) {
	d.loginURL = url
}

func (d *dashboard) GetRegisterURL() string {
	return d.registerURL
}

func (d *dashboard) SetRegisterURL(url string) {
	d.registerURL = url
}

// Theme methods
func (d *dashboard) IsThemeDark() bool {
	// Check if the theme is in the list of dark themes
	darkThemes := []string{
		"cyborg", "darkly", "slate", "solar", "superhero", "vapor", "dark",
	}

	return lo.Contains(darkThemes, d.theme)
}

func (d *dashboard) GetTheme() string {
	return d.theme
}

func (d *dashboard) SetTheme(theme string) {
	d.theme = theme
}

// GetThemeHandlerUrl returns the URL for the theme handler endpoint
func (d *dashboard) GetThemeHandlerUrl() string {
	if d.themeHandlerUrl == "" {
		return "/theme"
	}
	return d.themeHandlerUrl
}

// SetThemeHandlerUrl sets the URL for the theme handler endpoint
func (d *dashboard) SetThemeHandlerUrl(url string) {
	d.themeHandlerUrl = url
}

// GetThemesRestrict returns the map of restricted themes
func (d *dashboard) GetThemesRestrict() map[string]string {
	return d.themesRestrict
}

// SetThemesRestrict sets the map of restricted themes
func (d *dashboard) SetThemesRestrict(themes map[string]string) {
	d.themesRestrict = themes
}

// GetMenuShowText returns whether to show text in menu items
func (d *dashboard) GetMenuShowText() bool {
	return d.menuShowText
}

// SetMenuShowText sets whether to show text in menu items
func (d *dashboard) SetMenuShowText(showText bool) {
	d.menuShowText = showText
}

// GetNavbarBackground returns the current navbar background class and true if it's set,
// or an empty string and false if using the default
func (d *dashboard) GetNavbarBackground() (string, bool) {
	// If a specific background color is set, use it
	if d.navbarBackgroundColor != "" {
		return d.navbarBackgroundColor, true
	}

	// Otherwise, determine background based on theme
	if d.IsThemeDark() {
		return "bg-dark", false
	}
	return "bg-light", false
}
