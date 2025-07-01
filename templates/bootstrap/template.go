package bootstrap

import (
	"github.com/dracory/dashboard/types"
	"github.com/gouniverse/hb"
)

// Template implements the types.TemplateInterface for Bootstrap-based templates
type Template struct{}

// Ensure Template implements the TemplateInterface
var _ types.TemplateInterface = (*Template)(nil)

// ToHTML generates the complete HTML for the dashboard page
func (t *Template) ToHTML(dashboard types.DashboardInterface) string {
	// Create a new webpage
	webpage := hb.Webpage()

	// Set the page title
	webpage.SetTitle(dashboard.GetTitle())

	// Add favicon
	if favicon := favicon(); favicon != "" {
		webpage.SetFavicon(favicon)
	}

	// Add CSS
	if style := templateStyle(); style != "" {
		webpage.AddStyle(style)
	}

	// Add JavaScript
	if script := templateScript(); script != "" {
		webpage.AddScript(script)
	}

	// Add Bootstrap CSS
	webpage.AddStyleURL("https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/css/bootstrap.min.css")

	// Add Bootstrap Icons
	webpage.AddStyleURL("https://cdn.jsdelivr.net/npm/bootstrap-icons@1.8.1/font/bootstrap-icons.css")

	// Add Bootstrap JS Bundle with Popper
	webpage.AddScriptURL("https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/js/bootstrap.bundle.min.js")

	// Add custom CSS if any
	for _, style := range dashboard.GetStyles() {
		webpage.AddStyle(style)
	}

	// Add custom scripts if any
	for _, script := range dashboard.GetScripts() {
		webpage.AddScript(script)
	}

	// Get menu items
	menuItems := []types.MenuItem{}
	if menuProvider, ok := dashboard.(interface{ GetMenuItems() []types.MenuItem }); ok {
		menuItems = menuProvider.GetMenuItems()
	}

	// Get theme and branding settings
	navbarBackgroundColor := ""
	if bgProvider, ok := dashboard.(interface{ GetNavbarBackgroundColor() string }); ok {
		navbarBackgroundColor = bgProvider.GetNavbarBackgroundColor()
	}

	navbarTextColor := ""
	if textColorProvider, ok := dashboard.(interface{ GetNavbarTextColor() string }); ok {
		navbarTextColor = textColorProvider.GetNavbarTextColor()
	}

	// Create the main container
	container := hb.NewDiv().Class("container-fluid p-0").Style("min-height: 100vh;")

	// Add the top navigation
	navbar := topNavigation(dashboard, menuItems, "", "", "", navbarBackgroundColor, navbarTextColor)
	container.Child(hb.Raw(navbar))

	// Create the main content area
	contentRow := hb.NewDiv().Class("row g-0")

	// Add the sidebar if there are menu items
	if len(menuItems) > 0 {
		sidebar := hb.NewDiv().Class("col-md-3 col-lg-2 d-md-block bg-light sidebar collapse")
		sidebar.Child(hb.NewDiv().Class("position-sticky pt-3").Child(hb.Raw(dashboardMenuNavbar(menuItems))))
		contentRow.Child(sidebar)
	}

	// Add the main content area
	mainContent := hb.NewMain().Class("col-md-9 ms-sm-auto col-lg-10 px-md-4")
	mainContent.Child(hb.NewDiv().Class("container-fluid py-4").Child(hb.Raw(dashboard.GetContent())))

	contentRow.Child(mainContent)
	container.Child(contentRow)

	// Add the container to the webpage
	webpage.Body().Child(container)

	// Handle redirect if needed
	if redirectURL := dashboard.GetRedirectUrl(); redirectURL != "" {
		redirectTime := "0"
		if dashboard.GetRedirectTime() != "" {
			redirectTime = dashboard.GetRedirectTime()
		}
		webpage.Meta(hb.Meta().
			Attr("http-equiv", "refresh").
			Attr("content", redirectTime+"; url = "+redirectURL))
	}

	// Generate the final HTML
	return webpage.ToHTML()
}

// center is a helper function to center content
func center(content string) string {
	return `<div class="d-flex justify-content-center align-items-center" style="height: 100vh;">` + content + `</div>`
}
