package tabler

import (
	"github.com/dracory/dashboard/model"
	"github.com/dracory/dashboard/render/theme/shared"
	"github.com/gouniverse/hb"
)

// TablerTheme implements the shared.Theme interface for Tabler
type TablerTheme struct{}

// New creates a new instance of the Tabler theme
func New() *TablerTheme {
	return &TablerTheme{}
}

// Ensure TablerTheme implements shared.Theme
var _ shared.Theme = (*TablerTheme)(nil)

// GetName returns the name of the theme
func (t *TablerTheme) GetName() string {
	return "tabler"
}

// GetCSSLinks returns the CSS link tags for the theme
func (t *TablerTheme) GetCSSLinks(isDarkMode bool) []*hb.Tag {
	return GetTablerCDNLinks(isDarkMode)
}

// GetJSScripts returns the JavaScript script tags for the theme
func (t *TablerTheme) GetJSScripts() []*hb.Tag {
	return GetTablerCDNScripts()
}

// GetCustomCSS returns any custom CSS for the theme
func (t *TablerTheme) GetCustomCSS() string {
	return `
		.navbar-brand-image {
			height: 2rem;
		}
		.navbar-vertical.navbar-expand-lg {
			width: 15rem;
		}
		.navbar-vertical.navbar-expand-lg .navbar-collapse {
			margin: 0 -0.5rem;
		}
	`
}

// GetCustomJS returns any custom JavaScript for the theme
func (t *TablerTheme) GetCustomJS() string {
	return `
		// Theme switcher
		document.querySelectorAll('[data-bs-theme-value]').forEach(function(element) {
			element.addEventListener('click', function() {
				var theme = this.getAttribute('data-bs-theme-value');
				document.body.setAttribute('data-bs-theme', theme);
				localStorage.setItem('theme', theme);
			});
		});
		
		// Set theme from localStorage
		var theme = localStorage.getItem('theme');
		if (theme) {
			document.body.setAttribute('data-bs-theme', theme);
		}
	`
}

// RenderHeader renders the Tabler theme header
func (t *TablerTheme) RenderHeader(d shared.DashboardRenderer) *hb.Tag {
	header := hb.NewTag("header").Class("navbar navbar-expand-md navbar-light d-print-none")
	container := hb.NewTag("div").Class("container-xl")
	
	// Toggle button for mobile
	toggleBtn := hb.NewTag("button").
		Class("navbar-toggler").
		Attr("type", "button").
		Attr("data-bs-toggle", "collapse").
		Attr("data-bs-target", "#navbar-menu").
		Child(hb.NewTag("span").Class("navbar-toggler-icon"))

	// Logo/Brand
	logoLink := hb.NewTag("a").Attr("href", "/").Class("navbar-brand navbar-brand-autodark d-none-navbar-horizontal pe-0 pe-md-3")
	logoImg := hb.NewTag("img").Attr("src", "/static/logo.svg").Attr("alt", "Tabler").Class("navbar-brand-image")
	logoLink.Child(logoImg)

	// Header content
	headerContent := hb.Div().Class("navbar-nav flex-row order-md-last")
	
	// Theme toggle
	themeToggle := hb.NewTag("div").Class("nav-item dropdown d-none d-md-flex me-3")
	themeBtn := hb.NewTag("a").Attr("href", "#").Class("nav-link px-0").Attr("data-bs-toggle", "dropdown")
	themeIcon := hb.NewTag("i").Class("ti ti-bulb")
	themeBtn.Child(themeIcon)
	
	themeMenu := hb.Div().Class("dropdown-menu dropdown-menu-end")
	themeMenu.Child(hb.A().Href("#").Class("dropdown-item").Attr("data-bs-theme-value", "light").Text("Light"))
	themeMenu.Child(hb.A().Href("#").Class("dropdown-item").Attr("data-bs-theme-value", "dark").Text("Dark"))
	themeMenu.Child(hb.A().Href("#").Class("dropdown-item active").Attr("data-bs-theme-value", "system").Text("System"))
	
	themeToggle.Child(themeBtn)
	themeToggle.Child(themeMenu)
	headerContent.Child(themeToggle)

	// User menu
	user := d.GetUser()
	if user != (model.User{}) {
		userMenu := hb.Div().Class("nav-item dropdown")
		userLink := hb.A().Href("#").Class("nav-link d-flex lh-1 text-reset p-0").Attr("data-bs-toggle", "dropdown")
		
		// User avatar
		avatar := hb.NewTag("span").Class("avatar avatar-sm")
		if user.AvatarURL != "" {
			avatar.Child(hb.NewTag("img").Attr("src", user.AvatarURL).Class("rounded-circle"))
		} else {
			initials := ""
			if len(user.Name) > 0 {
				initials = string(user.Name[0])
			}
			avatar.Child(hb.NewTag("span").Text(initials).Class("avatar-initial"))
		}
		
		// User info
		userInfo := hb.Div().Class("d-none d-xl-block ps-2")
		userInfo.Child(hb.Div().Text(user.Name).Class("text-muted"))
		
		userLink.Child(avatar)
		userLink.Child(userInfo)
		
		// Dropdown menu
		dropdownMenu := hb.NewTag("div").Class("dropdown-menu dropdown-menu-end dropdown-menu-arrow")
		dropdownMenu.Child(hb.NewTag("a").Attr("href", "/profile").Class("dropdown-item").Text("Profile"))
		dropdownMenu.Child(hb.NewTag("a").Attr("href", "/settings").Class("dropdown-item").Text("Settings"))
		dropdownMenu.Child(hb.NewTag("hr").Class("dropdown-divider"))
		dropdownMenu.Child(hb.NewTag("a").Attr("href", "/logout").Class("dropdown-item").Text("Logout"))
		
		userMenu.Child(userLink)
		userMenu.Child(dropdownMenu)
		headerContent.Child(userMenu)
	}

	// Build header structure
	container.Child(toggleBtn)
	container.Child(logoLink)
	
	// Collapsible content
	collapse := hb.Div().Class("navbar-collapse").ID("navbar-menu")
	collapseInner := hb.Div().Class("d-flex flex-column flex-md-row flex-fill align-items-stretch align-items-md-center")
	
	nav := hb.NewDiv().Class("flex-grow-1")
	
	// Add search bar
	searchForm := hb.NewForm().Class("flex-grow-1").Action("/search").Method("GET")
	searchGroup := hb.NewDiv().Class("input-icon")
	searchInput := hb.NewInput().Type("text").Class("form-control").Placeholder("Search...").Attr("aria-label", "Search")
	searchIcon := hb.NewSpan().Class("input-icon-addon")
	searchIcon.AddChild(hb.NewI().Class("ti ti-search"))
	searchGroup.AddChild(searchInput)
	searchGroup.AddChild(searchIcon)
	searchForm.AddChild(searchGroup)
	nav.AddChild(searchForm)
	
	collapseInner.AddChild(nav)
	collapse.AddChild(collapseInner)
	container.AddChild(collapse)
	container.AddChild(headerContent)
	header.AddChild(container)
	
	return header
}

// RenderFooter renders the Tabler theme footer
func (t *TablerTheme) RenderFooter(d shared.DashboardRenderer) *hb.Tag {
	footer := hb.NewTag("footer").Class("footer footer-transparent d-print-none")
	container := hb.NewTag("div").Class("container-xl")
	
	// Footer content
	row := hb.Div().Class("row text-center align-items-center flex-row-reverse")
	
	// Left side (copyright)
	leftCol := hb.NewTag("div").Class("col-lg-auto ms-lg-auto")
	leftList := hb.NewTag("ul").Class("list-inline list-inline-dots mb-0")
	leftList.Child(hb.NewTag("li").Class("list-inline-item").Child(hb.NewTag("a").Attr("href", "/").Text("Home")))
	leftList.Child(hb.NewTag("li").Class("list-inline-item").Child(hb.NewTag("a").Attr("href", "/about").Text("About")))
	leftCol.Child(leftList)
	
	// Right side (copyright)
	rightCol := hb.NewTag("div").Class("col-12 col-lg-auto mt-3 mt-lg-0")
	rightCol.Child(hb.NewTag("ul").Class("list-inline list-inline-dots mb-0").
		Child(hb.NewTag("li").Class("list-inline-item").
			Text(" 2023 ").
			Child(hb.NewTag("a").Attr("href", "https://tabler.io/").Class("link-secondary").Text("Tabler").Attr("target", "_blank")).
			Text(" Dashboard")))
	
	row.Child(leftCol)
	row.Child(rightCol)
	container.Child(row)
	footer.Child(container)
	footer.AddChild(container)
	
	return footer
}
