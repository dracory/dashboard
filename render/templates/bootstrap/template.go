package bootstrap

import (
	"github.com/dracory/dashboard/render"
	"github.com/dracory/dashboard/render/templates/shared"
	"github.com/gouniverse/hb"
)

// BootstrapTheme implements the shared.Theme interface for Bootstrap 5
type BootstrapTemplate struct {
	shared.DefaultTemplate // Embed DefaultTemplate to inherit default implementations
}

// New creates a new instance of the Bootstrap theme
func New() *BootstrapTemplate {
	return &BootstrapTemplate{}
}

// RenderHeader renders the Bootstrap theme header
func (t *BootstrapTemplate) RenderHeader(d shared.DashboardRenderer) *hb.Tag {
	header := hb.NewTag("header").Class("navbar navbar-expand-lg navbar-light bg-light mb-4")
	container := hb.NewTag("div").Class("container-fluid")

	// Brand/logo
	logoLink := hb.NewTag("a").Attr("href", "/").Class("navbar-brand")
	logoLink.Child(hb.NewTag("span").Text("Dashboard"))

	// Toggle button for mobile
	toggleBtn := hb.NewTag("button").
		Class("navbar-toggler").
		Attr("type", "button").
		Attr("data-bs-toggle", "collapse").
		Attr("data-bs-target", "#navbarNav").
		Child(hb.NewTag("span").Class("navbar-toggler-icon"))

	// Navbar content
	navbarCollapse := hb.NewTag("div").Class("collapse navbar-collapse").ID("navbarNav")
	navbarNav := hb.NewTag("ul").Class("navbar-nav me-auto")
	navbarNav.Child(hb.NewTag("li").Class("nav-item").
		Child(hb.NewTag("a").Class("nav-link").Attr("href", "/").Text("Home")))
	navbarCollapse.Child(navbarNav)

	// User menu
	user := d.GetUser()
	if user.Name != "" {
		userMenu := hb.NewTag("div").Class("dropdown")
		userButton := hb.NewTag("button").
			Class("btn btn-outline-secondary dropdown-toggle").
			Attr("type", "button").
			Attr("data-bs-toggle", "dropdown").
			Text(user.Name)

		dropdownMenu := hb.NewTag("ul").Class("dropdown-menu dropdown-menu-end")
		dropdownMenu.Child(hb.NewTag("li").
			Child(hb.NewTag("a").Class("dropdown-item").
				Attr("href", "/profile").
				Text("Profile")))
		dropdownMenu.Child(hb.NewTag("li").
			Child(hb.NewTag("hr").Class("dropdown-divider")))
		dropdownMenu.Child(hb.NewTag("li").
			Child(hb.NewTag("a").Class("dropdown-item").
				Attr("href", "/logout").
				Text("Logout")))

		userMenu.Child(userButton)
		userMenu.Child(dropdownMenu)
		navbarCollapse.Child(hb.NewTag("div").Class("ms-auto").Child(userMenu))
	} else {
		loginBtn := hb.NewTag("a").
			Class("btn btn-outline-primary me-2").
			Attr("href", "/login").
			Text("Login")
		signupBtn := hb.NewTag("a").
			Class("btn btn-primary").
			Attr("href", "/register").
			Text("Sign up")
		navbarCollapse.Child(hb.NewTag("div").Class("ms-auto d-flex").
			Child(loginBtn).
			Child(signupBtn))
	}

	container.Child(logoLink)
	container.Child(toggleBtn)
	container.Child(navbarCollapse)
	header.Child(container)

	return header
}

// RenderFooter renders the Bootstrap theme footer
func (t *BootstrapTemplate) RenderFooter(d shared.DashboardRenderer) *hb.Tag {
	footer := hb.NewTag("footer").Class("bg-light py-4 mt-5")
	container := hb.NewTag("div").Class("container")

	// Footer content
	row := hb.NewTag("div").Class("row")

	// Left side (links)
	leftCol := hb.NewTag("div").Class("col-md-6")
	linkList := hb.NewTag("ul").Class("list-inline mb-0")
	linkList.Child(hb.NewTag("li").Class("list-inline-item").
		Child(hb.NewTag("a").Attr("href", "/about").Text("About")))
	linkList.Child(hb.NewTag("li").Class("list-inline-item").
		Text("•").Class("mx-2").Style("color: #6c757d;"))
	linkList.Child(hb.NewTag("li").Class("list-inline-item").
		Child(hb.NewTag("a").Attr("href", "/privacy").Text("Privacy")))
	leftCol.Child(linkList)

	// Right side (copyright)
	rightCol := hb.NewTag("div").Class("col-md-6 text-md-end")
	rightCol.Child(hb.NewTag("p").Class("mb-0 text-muted").
		Text("© 2023 Dashboard. All rights reserved."))

	row.Child(leftCol)
	row.Child(rightCol)
	container.Child(row)
	footer.Child(container)

	return footer
}

// Ensure BootstrapTheme implements shared.Template
var _ shared.Template = (*BootstrapTemplate)(nil)

// GetName returns the name of the theme
func (t *BootstrapTemplate) GetName() string {
	return render.THEME_BOOTSTRAP
}

// GetCSSLinks returns the CSS link tags for the theme
func (t *BootstrapTemplate) GetCSSLinks(isDarkMode bool) []*hb.Tag {
	return GetBootstrapCDNLinks()
}

// GetJSScripts returns the JavaScript script tags for the theme
func (t *BootstrapTemplate) GetJSScripts() []*hb.Tag {
	return GetBootstrapCDNScripts()
}

// GetCustomCSS returns any custom CSS for the theme
func (t *BootstrapTemplate) GetCustomCSS() string {
	return `
		/* Bootstrap 5 custom styles */
		.sidebar {
			min-height: 100vh;
			background: #f8f9fa;
			padding: 20px 0;
		}
		.main-content {
			padding: 20px;
		}
		.navbar-brand img {
			height: 2rem;
		}
	`
}

// GetCustomJS returns any custom JavaScript for the theme
func (t *BootstrapTemplate) GetCustomJS() string {
	return `
		// Initialize tooltips
		var tooltipTriggerList = [].slice.call(document.querySelectorAll('[data-bs-toggle="tooltip"]'));
		var tooltipList = tooltipTriggerList.map(function (tooltipTriggerEl) {
			return new bootstrap.Tooltip(tooltipTriggerEl);
		});

		// Enable popovers
		var popoverTriggerList = [].slice.call(document.querySelectorAll('[data-bs-toggle="popover"]'));
		var popoverList = popoverTriggerList.map(function (popoverTriggerEl) {
			return new bootstrap.Popover(popoverTriggerEl);
		});

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
func (t *BootstrapTemplate) RenderPage(content string, d shared.DashboardRenderer) (*hb.Tag, error) {
	// Create the head section
	head := hb.NewTag("head").
		Child(hb.NewTag("meta").Attr("charset", "utf-8")).
		Child(hb.NewTag("meta").Attr("name", "viewport").Attr("content", "width=device-width, initial-scale=1")).
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

	// Create the body section with Bootstrap classes
	bodyAttrs := map[string]string{
		"class": "d-flex flex-column min-vh-100",
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

	// Create main content area
	mainContent := hb.NewTag("main").Class("flex-grow-1 py-3")

	// Create container for content
	container := hb.NewDiv().Class("container")
	container.Child(hb.NewHTML(content))
	mainContent.Child(container)

	body.Child(mainContent)

	// Add footer
	footer := t.RenderFooter(d)
	if footer != nil {
		body.Child(footer)
	}

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
func (t *BootstrapTemplate) isDarkColorScheme(d shared.DashboardRenderer) bool {
	return d.GetNavbarBackgroundColorMode() == "dark"
}
