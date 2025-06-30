package tabler

import (
	"fmt"
	"github.com/dracory/dashboard/model"
	"github.com/dracory/dashboard/render"
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
	return render.THEME_TABLER
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
	fmt.Printf("[DEBUG] TablerTheme.RenderHeader called for theme: %s, dashboard theme: %s\n", t.GetName(), d.GetThemeName())
	
	// Main header container
	header := hb.NewHeader().Class("navbar navbar-expand-md navbar-light")
	header.Attr("data-bs-theme", "light") // Force light theme for header

	// Container for header content
	container := hb.NewDiv().Class("container-fluid")

	// Logo on the left
	logoLink := hb.NewLink().Href("/").Class("navbar-brand")
	logoImg := hb.NewImage().Src(d.GetLogoImageURL()).Alt("Logo").Class("navbar-brand-image")
	logoLink.Child(logoImg)

	// Mobile menu toggle
	mobileToggle := hb.NewButton().Class("navbar-toggler")
	mobileToggle.Attr("type", "button")
	mobileToggle.Attr("data-bs-toggle", "collapse")
	mobileToggle.Attr("data-bs-target", "#navbar-menu")
	mobileToggle.Child(hb.NewSpan().Class("navbar-toggler-icon"))

	// Main navigation wrapper
	navbarCollapse := hb.NewDiv().Class("collapse navbar-collapse")
	navbarCollapse.Attr("id", "navbar-menu")

	// Primary navigation container
	navbarNav := hb.NewDiv().Class("d-flex flex-column flex-md-row flex-fill align-items-stretch align-items-md-center")
	
	// Main navigation items
	navbarNavInner := hb.NewTag("ul").Class("navbar-nav")
	
	// Add menu items
	for _, item := range d.GetMenuItems() {
		navbarNavInner.Child(renderNavItem(item))
	}
	
	navbarNav.Child(navbarNavInner)
	navbarCollapse.Child(navbarNav)

	// Header controls (right side)
	headerControls := hb.NewDiv().Class("d-flex order-lg-2 ms-auto")

	// Search form
	searchForm := hb.NewDiv().Class("d-none d-md-flex me-3")
	searchGroup := hb.NewDiv().Class("input-icon")
	searchInput := hb.NewInput()
	searchInput.Attr("type", "text").Class("form-control").Attr("placeholder", "Search...")
	searchButton := hb.NewSpan().Class("input-icon-addon")
	searchButton.Child(hb.NewI().Class("ti ti-search"))
	searchGroup.Child(searchInput).Child(searchButton)
	searchForm.Child(searchGroup)

	// Notifications
	notifications := hb.NewDiv().Class("dropdown me-3")
	notificationsButton := hb.NewButton().Class("btn btn-icon btn-ghost-secondary")
	notificationsButton.Attr("type", "button")
	notificationsButton.Child(hb.NewI().Class("ti ti-bell"))
	notifications.Child(notificationsButton)

	// User menu
	userMenu := hb.NewDiv().Class("dropdown")
	userButton := hb.NewButton().Class("btn btn-ghost-secondary")
	userButton.Attr("type", "button")
	userButton.Attr("data-bs-toggle", "dropdown")
	userButton.Attr("aria-expanded", "false")
	
	userAvatar := hb.NewSpan().Class("avatar avatar-sm")
	userAvatar.Style("background-image: url('https://ui-avatars.com/api/?name=User&background=random')")
	userButton.Child(userAvatar)

	// User dropdown menu
	userDropdown := hb.NewDiv().Class("dropdown-menu dropdown-menu-end")
	
	// Add user menu items
	userDropdown.Child(hb.NewLink().Href("/profile").Class("dropdown-item").Child(hb.NewI().Class("ti ti-user me-2")).Text("Profile"))
	userDropdown.Child(hb.NewLink().Href("/settings").Class("dropdown-item").Child(hb.NewI().Class("ti ti-settings me-2")).Text("Settings"))
	userDropdown.Child(hb.NewDiv().Class("dropdown-divider"))
	userDropdown.Child(hb.NewLink().Href("/logout").Class("dropdown-item").Child(hb.NewI().Class("ti ti-logout me-2")).Text("Logout"))

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

// renderNavItem creates a navigation item for the header with proper Tabler styling
func renderNavItem(item model.MenuItem) *hb.Tag {
	navItem := hb.NewTag("li").Class("nav-item")
	if len(item.SubMenu) > 0 {
		navItem.Class("dropdown")
	}

	// Create the main nav link
	linkAttributes := make(map[string]string)
	linkAttributes["href"] = item.URL
	if item.NewWindow {
		linkAttributes["target"] = "_blank"
	}
	if item.OnClick != "" {
		linkAttributes["onclick"] = item.OnClick
	}
	
	navLink := hb.NewTag("a")
	if len(item.SubMenu) > 0 {
		navLink.Class("nav-link dropdown-toggle")
		navLink.Attr("data-bs-toggle", "dropdown")
		navLink.Attr("role", "button")
		navLink.Attr("aria-expanded", "false")
	} else {
		navLink.Class("nav-link")
	}
	
	// Set attributes
	for k, v := range linkAttributes {
		navLink.Attr(k, v)
	}
	
	// Add icon if present
	if item.Icon != "" {
		navLink.Child(hb.NewI().Class(item.Icon).Class("me-2"))
	}
	
	// Add text
	navLink.Child(hb.NewSpan().Text(item.Text))
	
	// Add badge if present
	if item.BadgeText != "" {
		badgeClass := "badge ms-2"
		if item.BadgeClass != "" {
			badgeClass += " bg-" + item.BadgeClass
		} else {
			badgeClass += " bg-primary"
		}
		navLink.Child(hb.NewSpan().Class(badgeClass).Text(item.BadgeText))
	}
	
	// Handle submenu if present
	if len(item.SubMenu) > 0 {
		dropdownMenu := hb.NewTag("ul").Class("dropdown-menu dropdown-menu-arrow")
		
		for _, subItem := range item.SubMenu {
			dropdownItem := hb.NewTag("li")
			
			// Handle divider
			if subItem.ID == "divider" {
				dropdownMenu.Child(hb.NewTag("li").Child(hb.NewTag("hr").Class("dropdown-divider")))
				continue
			}
			
			subLink := hb.NewTag("a").Class("dropdown-item")
			subLink.Attr("href", subItem.URL)
			
			if subItem.NewWindow {
				subLink.Attr("target", "_blank")
			}
			if subItem.OnClick != "" {
				subLink.Attr("onclick", subItem.OnClick)
			}
			
			// Add icon if present
			if subItem.Icon != "" {
				subLink.Child(hb.NewI().Class(subItem.Icon).Class("me-2"))
			}
			
			// Add text
			subLink.Child(hb.NewSpan().Text(subItem.Text))
			
			// Add badge if present
			if subItem.BadgeText != "" {
				badgeClass := "badge ms-2"
				if subItem.BadgeClass != "" {
					badgeClass += " bg-" + subItem.BadgeClass
				} else {
					badgeClass += " bg-primary"
				}
				subLink.Child(hb.NewSpan().Class(badgeClass).Text(subItem.BadgeText))
			}
			
			dropdownItem.Child(subLink)
			dropdownMenu.Child(dropdownItem)
		}
		
		navItem.Child(navLink).Child(dropdownMenu)
	} else {
		navItem.Child(navLink)
	}
	
	return navItem
}

// RenderFooter renders the Tabler theme footer
func (t *TablerTheme) RenderFooter(d model.DashboardRenderer) *hb.Tag {
	fmt.Printf("[DEBUG] TablerTheme.RenderFooter called for theme: %s, dashboard theme: %s\n", t.GetName(), d.GetThemeName())
	
	// Create footer with proper Tabler classes
	footer := hb.NewTag("footer").Class("footer footer-transparent d-print-none")
	container := hb.NewTag("div").Class("container-xl")
	
	// Footer content row
	row := hb.Div().Class("row text-center align-items-center flex-row-reverse")
	
	// Left side (menu links)
	leftCol := hb.NewTag("div").Class("col-lg-auto ms-lg-auto")
	leftList := hb.NewTag("ul").Class("list-inline list-inline-dots mb-0")
	leftList.Child(hb.NewTag("li").Class("list-inline-item").Child(hb.NewTag("a").Attr("href", "/").Text("Home")))
	leftList.Child(hb.NewTag("li").Class("list-inline-item").Child(hb.NewTag("a").Attr("href", "/about").Text("About")))
	leftCol.Child(leftList)
	
	// Right side (copyright)
	rightCol := hb.NewTag("div").Class("col-12 col-lg-auto mt-3 mt-lg-0")
	rightCol.Child(hb.NewTag("ul").Class("list-inline list-inline-dots mb-0").
		Child(hb.NewTag("li").Class("list-inline-item").
			Text("© 2023 ").
			Child(hb.NewTag("a").Attr("href", "https://tabler.io/").Class("link-secondary").Text("Tabler").Attr("target", "_blank")).
			Text(" Dashboard")))
	
	// Assemble the footer
	row.Child(leftCol)
	row.Child(rightCol)
	container.Child(row)
	footer.Child(container)
	
	return footer
}
