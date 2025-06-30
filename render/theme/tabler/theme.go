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

// RenderHeader renders the header of the dashboard
func (t *TablerTheme) RenderHeader(d model.DashboardRenderer) *hb.Tag {
	header := hb.NewHeader().Class("navbar navbar-expand-md navbar-dark")

	// Container for header content
	container := hb.NewDiv().Class("container-fluid")

	// Logo on the left
	logoLink := hb.NewLink().Href("/").Class("navbar-brand me-0 me-md-3")
	logoImg := hb.NewImage().Src(d.GetLogoImageURL()).Alt("Logo").Style("height: 32px")
	logoLink.Child(logoImg)

	// Mobile menu toggle
	mobileToggle := hb.NewButton().Class("navbar-toggler")
	mobileToggle.Attr("type", "button")
	mobileToggle.Attr("data-bs-toggle", "collapse")
	mobileToggle.Attr("data-bs-target", "#navbar-menu")
	mobileToggle.Child(hb.NewSpan().Class("navbar-toggler-icon"))

	// Header controls (right side)
	headerControls := hb.NewDiv().Class("d-flex align-items-center ms-auto")

	// Search form
	searchForm := hb.NewForm().Class("d-none d-md-flex me-3")
	searchGroup := hb.NewDiv().Class("input-group input-group-flat")
	searchInput := hb.NewInput().Type("search").Class("form-control").Placeholder("Search...")
	searchButton := hb.NewButton().Class("btn btn-ghost-secondary btn-icon").Child(hb.NewI().Class("ti ti-search"))
	searchGroup.Child(searchInput).Child(searchButton)
	searchForm.Child(searchGroup)

	// Notifications
	notifications := hb.NewDiv().Class("dropdown me-3")
	notificationsButton := hb.NewButton().Class("btn btn-ghost-secondary btn-icon")
	notificationsButton.Child(hb.NewI().Class("ti ti-bell"))
	notifications.Child(notificationsButton)

	// User menu
	userMenu := hb.NewDiv().Class("dropdown")
	userButton := hb.NewButton().Class("btn btn-ghost-secondary btn-icon")
	userAvatar := hb.NewSpan().Class("avatar avatar-sm").Style("background-image: url('https://ui-avatars.com/api/?name=User&background=random')")
	userButton.Child(userAvatar)

	// User dropdown menu
	userDropdown := hb.NewDiv().Class("dropdown-menu dropdown-menu-end")
	
	// Add user menu items
	userDropdown.Child(hb.NewLink().Href("/profile").Class("dropdown-item").Text("Profile"))
	userDropdown.Child(hb.NewLink().Href("/settings").Class("dropdown-item").Text("Settings"))
	userDropdown.Child(hb.NewDiv().Class("dropdown-divider"))
	userDropdown.Child(hb.NewLink().Href("/logout").Class("dropdown-item").Text("Logout"))

	userMenu.Child(userButton).Child(userDropdown)

	// Add controls to header
	headerControls.Child(searchForm).Child(notifications).Child(userMenu)

	// Main navigation menu
	menu := hb.NewDiv().Class("collapse navbar-collapse").ID("navbar-menu")
	menuList := hb.NewDiv().Class("navbar-nav")

	// Add menu items
	for _, item := range d.GetMenuItems() {
		menuItemClass := "nav-item"
		if item.Active {
			menuItemClass += " active"
		}

		menuItem := hb.NewDiv().Class(menuItemClass)
		linkClass := "nav-link"
		if len(item.SubMenu) > 0 {
			linkClass += " dropdown-toggle"
		}
		link := hb.NewLink().Href(item.URL).Class(linkClass)
		
		// Add icon if exists
		if item.Icon != "" {
			icon := hb.NewI().Class(item.Icon + " me-1")
			link.Child(icon)
		}
		
		// Add text
		if item.Text != "" {
			link.Child(hb.Text(item.Text))
		}

		// Handle submenu if exists
		if len(item.SubMenu) > 0 {
			link.Attr("data-bs-toggle", "dropdown")
			dropdown := hb.NewDiv().Class("dropdown-menu")
			
			for _, child := range item.SubMenu {
				dropdown.Child(hb.NewLink().Href(child.URL).Class("dropdown-item").Text(child.Text))
			}
			
			menuItem.Child(link).Child(dropdown)
		} else {
			menuItem.Child(link)
		}

		menuList.Child(menuItem)
	}

	menu.Child(menuList)

	// Add all elements to container
	container.Child(logoLink)
	container.Child(mobileToggle)
	container.Child(menu)
	container.Child(headerControls)

	header.Child(container)

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
