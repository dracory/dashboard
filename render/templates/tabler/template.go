package tabler

import (
	"fmt"

	"github.com/dracory/dashboard/config"
	"github.com/dracory/dashboard/model"
	"github.com/dracory/dashboard/render/templates/shared"
	hb "github.com/gouniverse/hb"
)

// TablerTemplate implements the shared.Template interface for Tabler
type TablerTemplate struct {
	shared.DefaultTemplate // Embed DefaultTemplate to inherit default implementations
}

// New creates a new instance of the Tabler template
func New() *TablerTemplate {
	return &TablerTemplate{}
}

// Ensure TablerTemplate implements shared.Template
var _ shared.Template = (*TablerTemplate)(nil)

// GetName returns the name of the template
func (t *TablerTemplate) GetName() string {
	return config.TEMPLATE_TABLER
}

// GetCSSLinks returns the CSS link tags for the template
func (t *TablerTemplate) GetCSSLinks(isDarkMode bool) []*hb.Tag {
	return GetTablerCDNLinks(isDarkMode)
}

// GetJSScripts returns the JavaScript script tags for the template
func (t *TablerTemplate) GetJSScripts() []*hb.Tag {
	return GetTablerCDNScripts()
}

// GetCustomCSS returns any custom CSS for the template
func (t *TablerTemplate) GetCustomCSS() string {
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
func (t *TablerTemplate) GetCustomJS() string {
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

// RenderPage renders a complete page with the given content and dashboard renderer
func (t *TablerTemplate) RenderPage(content string, d shared.DashboardRenderer) (*hb.Tag, error) {
	// Create the head section
	head := hb.NewTag("head").
		Child(hb.NewTag("meta").Attr("charset", "utf-8")).
		Child(hb.NewTag("meta").Attr("name", "viewport").Attr("content", "width=device-width, initial-scale=1, viewport-fit=cover")).
		Child(hb.NewTag("meta").Attr("http-equiv", "X-UA-Compatible").Attr("content", "ie=edge")).
		Child(hb.NewTag("title").Text("Dashboard"))

	// Add favicon if available
	if d.GetFaviconURL() != "" {
		head.Child(hb.NewTag("link").Attr("rel", "icon").Attr("href", d.GetFaviconURL()))
	}

	// Add theme CSS
	cssLinks := t.GetCSSLinks(t.isDarkColorScheme(d))
	for _, link := range cssLinks {
		head.Child(link)
	}

	// Create the body section with Tabler classes
	bodyAttrs := map[string]string{
		"class": "antialiased",
	}
	if t.isDarkColorScheme(d) {
		bodyAttrs["data-bs-theme"] = "dark"
	}

	body := hb.NewTag("body").Attrs(bodyAttrs)

	// Add header
	header := t.RenderHeader(d)
	if header != nil {
		body.Child(header)
	}

	// Create page wrapper
	pageWrapper := hb.NewDiv().Class("page-wrapper")

	// Add main content
	contentContainer := hb.NewDiv().Class("container-xl").
		AddChild(hb.NewHTML(content))

	pageContent := hb.NewDiv().Class("page-body").
		AddChild(contentContainer)

	pageWrapper.Child(pageContent)

	// Add footer
	footer := t.RenderFooter(d)
	if footer != nil {
		pageWrapper.Child(footer)
	}

	// Create page container
	pageContainer := hb.NewDiv().Class("page").
		AddChild(pageWrapper)

	body.Child(pageContainer)

	// Add JavaScript
	for _, script := range t.GetJSScripts() {
		body.Child(script)
	}

	// Add custom JavaScript
	if customJS := t.GetCustomJS(); customJS != "" {
		body.Child(hb.NewTag("script").Text(customJS))
	}

	// Create HTML document
	html := hb.NewTag("html").Attr("lang", "en").
		Child(head).
		Child(body)

	return hb.Wrap().
		Child(hb.NewHTML("<!DOCTYPE html>")).
		Child(html), nil
}

// isDarkColorScheme checks if the color scheme should be dark
func (t *TablerTemplate) isDarkColorScheme(d shared.DashboardRenderer) bool {
	return d.GetNavbarBackgroundColorMode() == "dark"
}

// RenderHeader renders the header of the dashboard
func (t *TablerTemplate) RenderHeader(d shared.DashboardRenderer) *hb.Tag {
	fmt.Printf("[DEBUG] TablerTheme.RenderHeader called for theme: %s, dashboard theme: %s\n", t.GetName(), d.GetTemplateName())

	// Create header container
	header := hb.NewHeader().Class("navbar navbar-expand-md d-print-none")
	header.Style("background-color: #206bc4;")
	container := hb.NewDiv().Class("container-xl")

	// First row container
	firstRow := hb.NewDiv().Class("d-flex flex-row w-100 align-items-center")

	// Logo on left side of first row
	logoLink := hb.NewLink().Href("/").Class("navbar-brand navbar-brand-autodark me-0 me-md-3")
	logoURL := d.GetLogoImageURL()
	if logoURL != "" {
		logoLink.Child(hb.NewImage().Src(logoURL).Alt("Logo").Style("height: 32px;"))
	} else {
		logoLink.Child(hb.Text("Dashboard"))
	}

	// User menu on right side of first row
	userMenu := hb.NewDiv().Class("ms-auto d-flex")
	user := d.GetUser()
	userName := "User"
	if user.Name != "" {
		userName = user.Name
	}

	// User avatar
	avatarURL := "https://ui-avatars.com/api/?name=" + userName + "&background=0D8ABC&color=fff"
	if user.AvatarURL != "" {
		avatarURL = user.AvatarURL
	}

	// User dropdown
	dropdownMenu := hb.NewDiv().Class("dropdown-menu dropdown-menu-end dropdown-menu-arrow")

	// Add user menu items
	for _, item := range d.GetUserMenu() {
		dropdownMenu.Child(hb.NewLink().Href(item.URL).Class("dropdown-item").Text(item.Text))
	}

	// Add dropdown to user menu
	userMenu.Child(hb.NewDiv().Class("nav-item dropdown").
		Child(hb.NewLink().Href("#").Class("nav-link d-flex lh-1 text-white p-0").
			Attr("data-bs-toggle", "dropdown").
			Child(hb.NewSpan().Class("avatar avatar-sm me-2").
				Child(hb.NewImage().Src(avatarURL).Class("avatar-img rounded-circle").Alt(""))).
			Child(hb.NewDiv().Class("d-none d-xl-block ps-1").
				Child(hb.NewDiv().Text(userName)).
				Child(hb.NewDiv().Class("mt-1 small text-white-50").Text(user.Email)))))

	// Mobile toggle
	mobileToggle := hb.NewButton().Class("navbar-toggler")
	mobileToggle.Attr("type", "button")
	mobileToggle.Attr("data-bs-toggle", "collapse")
	mobileToggle.Attr("data-bs-target", "#navbar-menu")
	mobileToggle.Attr("aria-controls", "navbar-menu")
	mobileToggle.Attr("aria-expanded", "false")
	mobileToggle.Attr("aria-label", "Toggle navigation")
	mobileToggle.Child(hb.NewSpan().Class("navbar-toggler-icon"))

	// Add logo and user menu to first row
	firstRow.Child(logoLink)
	firstRow.Child(userMenu)

	// Second row container for main navigation
	secondRow := hb.NewDiv().Class("navbar-expand-md")
	navbarCollapse := hb.NewDiv().Class("collapse navbar-collapse").ID("navbar-menu")
	navbarNav := hb.NewTag("ul").Class("navbar-nav")

	// Primary navigation container
	// Main navigation items (First row)
	navbarNavInner := hb.NewTag("ul").Class("navbar-nav me-auto")

	// Add main menu items (first level)
	for _, item := range d.GetMenuItems() {
		navbarNavInner.Child(t.renderNavItem(item)) // Start with level 0 for top-level items
	}

	navbarNav.Child(navbarNavInner)

	// Secondary navigation (Second row) - Only shown on larger screens
	secondaryNav := hb.NewTag("ul").Class("navbar-nav d-none d-lg-flex ms-auto")

	// Add secondary menu items if available
	if secondaryItems, ok := d.(interface{ GetSecondaryMenuItems() []model.MenuItem }); ok {
		for _, item := range secondaryItems.GetSecondaryMenuItems() {
			secondaryNav.Child(t.renderNavItem(item)) // Start with level 0 for top-level items
		}
	}

	// Add both navigation rows to the navbar collapse
	navbarCollapse.Child(navbarNav)

	// Check if we have any secondary navigation items by checking if the secondary nav has any children
	// The ToHTML() check ensures we don't add an empty nav
	if secondaryNav.ToHTML() != "<ul class=\"navbar-nav d-none d-lg-flex ms-auto\"></ul>" {
		navbarCollapse.Child(hb.NewDiv().Class("navbar-nav-secondary").Child(secondaryNav))
	}

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

	// User dropdown container
	userDropdownContainer := hb.NewDiv().Class("dropdown")
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

	userDropdownContainer.Child(userButton).Child(userDropdown)

	// Add controls to header
	headerControls.Child(searchForm).Child(notifications).Child(userMenu)

	// Main navigation menu is already rendered in navbarNav and navbarNavInner

	// Add menu items to navigation
	for _, item := range d.GetMenuItems() {
		navbarNav.Child(t.renderNavItem(item))
	}

	// Add navigation to second row
	navbarCollapse.Child(navbarNav)

	// Add mobile toggle and navbar to second row
	secondRow.Child(mobileToggle)
	secondRow.Child(navbarCollapse)

	// Add rows to container
	container.Child(firstRow)
	container.Child(secondRow)

	header.Child(container)

	// Add dropdown menu to user menu
	header.Child(dropdownMenu)

	return header
}

// renderNavItem creates a navigation item for the header with proper Tabler styling
func (t *TablerTemplate) renderNavItem(item model.MenuItem) *hb.Tag {
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
		navLink.Class("nav-link dropdown-toggle text-white")
		navLink.Attr("data-bs-toggle", "dropdown")
		navLink.Attr("role", "button")
		navLink.Attr("aria-expanded", "false")
	} else {
		navLink.Class("nav-link text-white")
	}

	// Add hover and active states
	navLink.Class("px-3 py-2 d-flex align-items-center")
	if item.Active {
		navLink.Class("active fw-bold")
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
		dropdownMenu := hb.NewTag("ul").Class("dropdown-menu dropdown-menu-arrow dropdown-menu-dark")
		dropdownMenu.Attr("data-bs-popper", "static")

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

// RenderFooter renders the Tabler template footer
func (t *TablerTemplate) RenderFooter(d shared.DashboardRenderer) *hb.Tag {
	fmt.Printf("[DEBUG] TablerTheme.RenderFooter called for theme: %s, dashboard template: %s\n", t.GetName(), d.GetTemplateName())

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
