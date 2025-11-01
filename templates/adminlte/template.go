package adminlte

import (
	"github.com/dracory/dashboard/templates/shared"
	"github.com/dracory/dashboard/types"
	"github.com/dracory/hb"
	"github.com/samber/lo"
)

// Template implements the types.TemplateInterface for AdminLTE-based templates
type Template struct{}

// Ensure Template implements the TemplateInterface
var _ types.TemplateInterface = (*Template)(nil)

// layout generates the main layout structure for the dashboard
func (t *Template) layout(dashboard types.DashboardInterface) *hb.Tag {
	// Main content wrapper
	contentWrapper := hb.Div().Class("content-wrapper")

	// Content header
	contentHeader := hb.Div().Class("content-header")
	containerFluid := hb.Div().Class("container-fluid")

	// Page title row
	row := hb.Div().Class("row mb-2")

	// Left side (page title)
	colSm6Left := hb.Div().Class("col-sm-6")
	colSm6Left.Child(hb.H1().Class("m-0").HTML(dashboard.GetTitle()))
	row.Child(colSm6Left)

	// Right side (breadcrumb)
	colSm6Right := hb.Div().Class("col-sm-6")
	breadcrumb := hb.Ol().Class("breadcrumb float-sm-right")

	homeLink := hb.Hyperlink().Href("/").HTML("Home")
	breadcrumb.Child(hb.Li().Class("breadcrumb-item").Child(homeLink))

	dashboardItem := hb.Li().Class("breadcrumb-item active").HTML("Dashboard")
	breadcrumb.Child(dashboardItem)

	colSm6Right.Child(breadcrumb)
	row.Child(colSm6Right)

	containerFluid.Child(row)
	contentHeader.Child(containerFluid)
	contentWrapper.Child(contentHeader)

	// Main content
	contentSection := hb.Section().Class("content")
	contentContainer := hb.Div().Class("container-fluid")
	contentContainer.Child(hb.Raw(dashboard.GetContent()))
	contentSection.Child(contentContainer)
	contentWrapper.Child(contentSection)

	return contentWrapper
}

func (t *Template) getStylesAndScripts(dashboard types.DashboardInterface) (
	styleURLs []string,
	scriptURLs []string,
	styles []string,
	scripts []string,
) {
	styleURLs = make([]string, 0)
	scriptURLs = make([]string, 0)
	styles = make([]string, 0)
	scripts = make([]string, 0)

	// Add CSS
	if style := templateStyle(); style != "" {
		styles = append(styles, style)
	}

	// Add JavaScript
	if script := templateScript(); script != "" {
		scripts = append(scripts, script)
	}

	// AdminLTE CSS
	styleURLs = append(styleURLs, "https://cdn.jsdelivr.net/npm/admin-lte@3.2.0/dist/css/adminlte.min.css")
	// Font Awesome
	styleURLs = append(styleURLs, "https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.4.0/css/all.min.css")
	// Google Font: Source Sans Pro
	styleURLs = append(styleURLs, "https://fonts.googleapis.com/css?family=Source+Sans+Pro:300,400,400i,700&display=fallback")
	// jQuery
	scriptURLs = append(scriptURLs, "https://code.jquery.com/jquery-3.6.0.min.js")
	// Bootstrap 4 JS Bundle with Popper
	scriptURLs = append(scriptURLs, "https://cdn.jsdelivr.net/npm/bootstrap@4.6.2/dist/js/bootstrap.bundle.min.js")
	// overlayScrollbars
	scriptURLs = append(scriptURLs, "https://cdn.jsdelivr.net/npm/overlayscrollbars@1.13.1/js/jquery.overlayScrollbars.min.js")
	// AdminLTE JS
	scriptURLs = append(scriptURLs, "https://cdn.jsdelivr.net/npm/admin-lte@3.2.0/dist/js/adminlte.min.js")
	// Initialize AdminLTE with default options
	scripts = append(scripts, "$(document).ready(function() { $('body').addClass('sidebar-mini'); });")

	styles = append(styles, "@import url('https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700&display=swap');")

	styleURLs = append(styleURLs, dashboard.GetStyleURLs()...)
	scriptURLs = append(scriptURLs, dashboard.GetScriptURLs()...)

	styles = append(styles, dashboard.GetStyles()...)
	scripts = append(scripts, dashboard.GetScripts()...)

	return lo.Uniq(styleURLs), lo.Uniq(scriptURLs), lo.Uniq(styles), lo.Uniq(scripts)
}

// ToHTML generates the complete HTML for the dashboard page
func (t *Template) ToHTML(dashboard types.DashboardInterface) string {
	styleURLs, scriptURLs, styles, scripts := t.getStylesAndScripts(dashboard)

	// Create a new webpage
	webpage := hb.Webpage()

	// Set the page title
	webpage.SetTitle(dashboard.GetTitle())

	// Add favicon
	webpage.SetFavicon(shared.Favicon())

	// Add styles URLs
	for _, styleURL := range styleURLs {
		if styleURL != "" {
			webpage.AddStyleURL(styleURL)
		}
	}

	// Add styles
	for _, style := range styles {
		if style != "" {
			webpage.AddStyle(style)
		}
	}

	// Add scripts URLs
	for _, scriptURL := range scriptURLs {
		if scriptURL != "" {
			webpage.AddScriptURL(scriptURL)
		}
	}

	// Add scripts
	for _, script := range scripts {
		if script != "" {
			webpage.AddScript(script)
		}
	}

	// Apply theme classes to body
	theme := dashboard.GetTheme()
	body := webpage.Body()

	// Add dark mode class if theme is dark
	if theme == ThemeDark {
		body.Class("dark-mode")
	}

	// Apply navbar and sidebar theming
	switch theme {
	case ThemeDark:
		body.Class("sidebar-dark-primary")
	case "blue":
		body.Class("sidebar-light-primary")
	default:
		// Default to light theme with primary color
		body.Class("sidebar-light-primary")
	}

	// Create main wrapper
	wrapper := hb.Div().Class("wrapper")

	// Add navigation
	navbar := topNavigation(dashboard)
	if navbar != nil {
		wrapper.Child(navbar)
	}

	// Add sidebar if not in modal mode
	if dashboard.GetMenuType() != "modal" {
		sidebar := menuOffcanvas(dashboard)
		if sidebar != nil {
			wrapper.Child(sidebar)
		}
	}

	// Add content wrapper with proper AdminLTE structure
	wrapper.Child(t.layout(dashboard))

	// Add footer
	footer := hb.NewFooter().
		Class("main-footer").
		Child(hb.NewDiv().
			Class("float-right d-none d-sm-block").
			Child(
				hb.NewDiv().Style("font-size: 85%;").Text("AdminLTE v3.2.0"),
			)).
		Child(hb.NewStrong().
			Text("Copyright &copy; 2023").
			Child(
				hb.NewA().Href("#").Text("AdminLTE.io"),
			).Text(". All rights reserved."))

	wrapper.Child(footer)

	// Add control sidebar
	controlSidebar := hb.NewAside().Class("control-sidebar control-sidebar-dark")
	wrapper.Child(controlSidebar)

	// Add main wrapper to body
	webpage.Body().Child(wrapper)

	// Add modal menu if needed
	if dashboard.GetMenuType() == "modal" {
		modal := menuModal(dashboard)
		if modal != nil {
			webpage.Body().Child(modal)
		}
	}

	// Generate the final HTML
	return webpage.ToHTML()
}
