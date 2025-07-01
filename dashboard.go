package dashboard

import (
	"github.com/dracory/dashboard/templates/bootstrap"
	"github.com/dracory/dashboard/types"
)

const TEMPLATE_BOOTSTRAP = "bootstrap"
const TEMPLATE_DEFAULT = TEMPLATE_BOOTSTRAP

type dashboard struct {
	content                   string
	faviconURL                string
	logoImageURL              string
	logoRawHtml               string
	logoRedirectURL           string
	menuMainItems             []types.MenuItem
	menuShowText              bool // controls whether to show text in menu items
	menuUserItems             []types.MenuItem
	menuQuickAccessItems      []types.MenuItem
	navbarBackgroundColorMode string
	navbarBackgroundColor     string
	navbarTextColor           string
	redirectTime              string      // redirect time (if any, in seconds)
	redirectUrl               string      // redirect URL (if any)
	scripts                   []string    // custom scripts defined by the user
	scriptURLs                []string    // custom script URLs defined by the user
	styles                    []string    // custom styles defined by the user
	styleURLs                 []string    // custom style URLs defined by the user
	theme                     string      // color mode: default, dark, light
	template                  string      // bootstrap (default), adminlte, tabler
	title                     string      // title of the webpage
	user                      *types.User // user object (if any)
	loginURL                  string      // login URL
	registerURL               string      // register URL
}

var _ types.DashboardInterface = (*dashboard)(nil)

// ToHTML returns the HTML of the dashboard
func (d *dashboard) ToHTML() string {
	templateName := d.GetTemplate()
	if templateName == "" {
		templateName = TEMPLATE_DEFAULT
	}

	var template types.TemplateInterface

	if templateName == TEMPLATE_BOOTSTRAP {
		template = &bootstrap.Template{}
	} else {
		template = &bootstrap.Template{}
	}

	if template == nil {
		return "Template not found: " + templateName
	}

	return template.ToHTML(d)
}

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

// GetTemplate returns the template name
func (d *dashboard) GetTemplate() string {
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

// Navbar utility methods
func (d *dashboard) NavbarHasBackgroundThemeClass() bool {
	hasNavbarBackgroundColor := d.navbarBackgroundColor != ""
	hasNavbarBackgroundTheme := !hasNavbarBackgroundColor && d.navbarBackgroundColorMode != ""
	return hasNavbarBackgroundTheme
}

func (d *dashboard) NavbarBackgroundThemeClass() string {
	if d.NavbarHasBackgroundThemeClass() {
		return "bg-" + d.navbarBackgroundColorMode
	}
	return ""
}

func (d *dashboard) NavbarButtonThemeClass() string {
	if d.NavbarHasBackgroundThemeClass() {
		return "btn-" + d.navbarBackgroundColorMode
	}
	return ""
}

// Navbar dropdown methods - these are stubs that need to be implemented
// with actual HTML generation logic
func (d *dashboard) NavbarDropdownQuickAccess(iconStyle string) string {
	// TODO: Implement actual quick access dropdown HTML generation
	return ""
}

func (d *dashboard) NavbarDropdownThemeSwitch() string {
	// TODO: Implement actual theme switch dropdown HTML generation
	return ""
}

func (d *dashboard) NavbarDropdownUser(iconStyle string) string {
	// TODO: Implement actual user dropdown HTML generation
	return ""
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
	// TODO: Implement actual theme detection logic
	return d.theme == "dark"
}

func (d *dashboard) GetTheme() string {
	return d.theme
}

func (d *dashboard) SetTheme(theme string) {
	d.theme = theme
}

// GetThemeHandlerUrl returns the URL for the theme handler endpoint
func (d *dashboard) GetThemeHandlerUrl() string {
	// Default to "/theme" if not set
	return "/theme"
}

// GetMenuShowText returns whether to show text in menu items
func (d *dashboard) GetMenuShowText() bool {
	return d.menuShowText
}
