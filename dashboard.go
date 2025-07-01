package dashboard

import (
	"strings"

	"github.com/dracory/dashboard/templates/bootstrap"
	"github.com/dracory/dashboard/types"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/icons"
	"github.com/samber/lo"
)

const TEMPLATE_BOOTSTRAP = "bootstrap"
const TEMPLATE_DEFAULT = TEMPLATE_BOOTSTRAP

// Menu type constants
const MENU_TYPE_MODAL = "modal"
const MENU_TYPE_OFFCANVAS = "offcanvas"

type dashboard struct {
	content                   string
	faviconURL                string
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
	redirectTime              string            // redirect time (if any, in seconds)
	redirectUrl               string            // redirect URL (if any)
	scripts                   []string          // custom scripts defined by the user
	scriptURLs                []string          // custom script URLs defined by the user
	styles                    []string          // custom styles defined by the user
	styleURLs                 []string          // custom style URLs defined by the user
	theme                     string            // color mode: default, dark, light
	themesRestrict            map[string]string // restricted theme options
	themeHandlerUrl           string            // URL for theme handler
	template                  string            // bootstrap (default), adminlte, tabler
	title                     string            // title of the webpage
	user                      *types.User       // user object (if any)
	loginURL                  string            // login URL
	registerURL               string            // register URL
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

// GetMenuType returns the menu type (modal or offcanvas)
func (d *dashboard) GetMenuType() string {
	if d.menuType == "" {
		return types.MENU_TYPE_OFFCANVAS // Default to offcanvas
	}
	return d.menuType
}

// SetMenuType sets the menu type (modal or offcanvas)
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

// Navbar dropdown methods
func (d *dashboard) NavbarDropdownQuickAccess(iconStyle string) string {
	if len(d.menuQuickAccessItems) == 0 {
		return ""
	}

	hasNavbarTextColor := d.navbarTextColor != ""
	buttonTheme := d.NavbarButtonThemeClass()

	button := hb.Button().
		ID("ButtonQuickAccess").
		Class("btn "+buttonTheme+" dropdown-toggle").
		Style("background:none;border:0px;").
		StyleIf(hasNavbarTextColor, "color: "+d.navbarTextColor+";").
		Type(hb.TYPE_BUTTON).
		Data("bs-toggle", "dropdown").
		Child(icons.Icon("bi bi-grid", 24, 24, ""))

	dropdownMenu := hb.Div().
		Class("dropdown-menu dropdown-menu-end").
		Style("min-width:300px;padding:0.5rem;")

	// Add quick access items from configuration
	if len(d.menuQuickAccessItems) > 0 {
		var menuItems []hb.TagInterface

		// Group items into rows of 3
		for i := 0; i < len(d.menuQuickAccessItems); i += 3 {
			end := i + 3
			if end > len(d.menuQuickAccessItems) {
				end = len(d.menuQuickAccessItems)
			}
			rowItems := d.menuQuickAccessItems[i:end]

			row := hb.Div().Class("row g-0")

			for _, item := range rowItems {
				icon := item.Icon
				if icon == "" {
					icon = "bi bi-app"
				}

				col := hb.Div().Class("col-4 text-center").
					Child(hb.Hyperlink().
						Class("dropdown-item d-flex flex-column align-items-center").
						Href(item.URL).
						Child(icons.Icon(icon, 24, 24, "")).
						Child(hb.Span().Text(item.Title).Class("mt-1")),
					)
				row.Child(col)
			}

			menuItems = append(menuItems, row)
		}

		dropdownMenu.Children(menuItems)
	}

	return hb.Div().
		Class("dropdown").
		Children([]hb.TagInterface{
			button,
			dropdownMenu,
		}).ToHTML()
}

func (d *dashboard) NavbarDropdownThemeSwitch() string {
	if d.themeHandlerUrl == "" {
		return ""
	}

	hasNavbarTextColor := d.navbarTextColor != ""
	buttonTheme := d.NavbarButtonThemeClass()
	isDark := d.IsThemeDark()

	// Import theme maps from bootstrap package
	// If themes are restricted, use those instead
	if len(d.themesRestrict) > 0 {
		// Use restricted themes
		// Create dropdown items
		var dropdownItems []hb.TagInterface
		for themeKey, themeName := range d.themesRestrict {
			icon := lo.TernaryF(d.IsThemeDark(), 
				func() string { return "bi bi-sun" }, 
				func() string { return "bi bi-moon-stars-fill" })

			active := lo.Ternary(d.theme == themeKey, " active", "")
			url := lo.Ternary(strings.Contains(d.themeHandlerUrl, "?"), 
				d.themeHandlerUrl+"&theme="+themeKey, 
				d.themeHandlerUrl+"?theme="+themeKey)

			dropdownItems = append(dropdownItems,
				hb.LI().Children([]hb.TagInterface{
					hb.Hyperlink().
						Class("dropdown-item"+active).
						Href(url).
						Child(hb.I().Class(icon+" me-2")).
						HTML(themeName),
				}),
			)
		}

		button := hb.Button().
			ID("buttonTheme").
			Class(buttonTheme+" dropdown-toggle").
			Style("background:none;border:0px;").
			StyleIf(hasNavbarTextColor, "color:"+d.navbarTextColor).
			Data("bs-toggle", "dropdown").
			Children([]hb.TagInterface{
				lo.Ternary(isDark, hb.I().Class("bi bi-sun"), hb.I().Class("bi bi-moon-stars-fill")),
			})

		return hb.Div().
			Class("dropdown").
			Child(button).
			Child(hb.UL().
				Class(buttonTheme + " dropdown-menu dropdown-menu-dark").
				Children(dropdownItems)).ToHTML()
	} else {
		// Use all available themes
		// Light Themes
		lightDropdownItems := lo.Map(lo.Keys(bootstrap.ThemesLight), func(theme string, index int) hb.TagInterface {
			name := bootstrap.ThemesLight[theme]
			active := lo.Ternary(d.theme == theme, " active", "")
			url := lo.Ternary(strings.Contains(d.themeHandlerUrl, "?"), 
				d.themeHandlerUrl+"&theme="+theme, 
				d.themeHandlerUrl+"?theme="+theme)

			return hb.LI().Children([]hb.TagInterface{
				hb.Hyperlink().
					Class("dropdown-item"+active).
					Child(hb.I().Class("bi bi-sun me-2")).
					HTML(name).
					Href(url).
					Attr("ref", "nofollow"),
			})
		})

		// Dark Themes
		darkDropdownItems := lo.Map(lo.Keys(bootstrap.ThemesDark), func(theme string, index int) hb.TagInterface {
			name := bootstrap.ThemesDark[theme]
			active := lo.Ternary(d.theme == theme, " active", "")
			url := lo.Ternary(strings.Contains(d.themeHandlerUrl, "?"), 
				d.themeHandlerUrl+"&theme="+theme, 
				d.themeHandlerUrl+"?theme="+theme)

			return hb.LI().Children([]hb.TagInterface{
				hb.Hyperlink().
					Class("dropdown-item"+active).
					Child(hb.I().Class("bi bi-moon-stars-fill me-2")).
					HTML(name).
					Href(url).
					Attr("ref", "nofollow"),
			})
		})

		button := hb.Button().
			ID("buttonTheme").
			Class(buttonTheme+" dropdown-toggle").
			Style("background:none;border:0px;").
			StyleIf(hasNavbarTextColor, "color:"+d.navbarTextColor).
			Data("bs-toggle", "dropdown").
			Children([]hb.TagInterface{
				lo.Ternary(isDark, hb.I().Class("bi bi-sun"), hb.I().Class("bi bi-moon-stars-fill")),
			})

		return hb.Div().
			Class("dropdown").
			Child(button).
			Child(hb.UL().
				Class(buttonTheme + " dropdown-menu dropdown-menu-dark").
				Children(lightDropdownItems).
				ChildIf(
					len(lo.Filter(darkDropdownItems, func(item hb.TagInterface, _ int) bool { return item != nil })) > 0 && 
					len(lo.Filter(lightDropdownItems, func(item hb.TagInterface, _ int) bool { return item != nil })) > 0,
					hb.LI().Children([]hb.TagInterface{
						hb.HR().Class("dropdown-divider"),
					}),
				).
				Children(darkDropdownItems)).ToHTML()
	}

}

func (d *dashboard) NavbarDropdownUser(iconStyle string) string {
	if d.user == nil {
		return ""
	}

	hasNavbarTextColor := d.navbarTextColor != ""
	buttonTheme := d.NavbarButtonThemeClass()
	userName := d.user.FirstName + " " + d.user.LastName

	dropdownUser := hb.Div().
		Class("dropdown").
		Children([]hb.TagInterface{
			hb.Button().
				ID("ButtonUser").
				Class("btn "+buttonTheme+" dropdown-toggle").
				Style("background:none;border:0px;").
				StyleIf(hasNavbarTextColor, "color: "+d.navbarTextColor+";").
				Type(hb.TYPE_BUTTON).
				Data("bs-toggle", "dropdown").
				Children([]hb.TagInterface{
					icons.Icon("bi-person", 24, 24, "").Style(iconStyle),
					hb.Span().
						Class("d-none d-md-inline-block").
						Text(userName).
						Style("margin-right:10px;"),
				}),
			hb.UL().
				Class("dropdown-menu dropdown-menu-dark").
				Class(buttonTheme).
				Children(lo.Map(d.menuUserItems, func(item types.MenuItem, _ int) hb.TagInterface {
					target := lo.Ternary(item.Target == "", "_self", item.Target)
					url := lo.Ternary(item.URL == "", "#", item.URL)

					return hb.LI().Children([]hb.TagInterface{
						hb.If(item.Title == "",
							hb.HR().
								Class("dropdown-divider"),
						),

						hb.If(item.Title != "",
							hb.Hyperlink().
								Class("dropdown-item").
								ChildIf(item.Icon != "", hb.Span().Class("icon").Style("margin-right: 5px;").HTML(item.Icon)).
								Text(item.Title).
								Href(url).
								Target(target),
						),
					})
				})),
		})

	return dropdownUser.ToHTML()
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
