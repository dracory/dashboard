package adminlte

import (
	"github.com/dracory/dashboard/types"
	"github.com/gouniverse/hb"
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

	// Add AdminLTE CSS
	webpage.AddStyle("https://cdn.jsdelivr.net/npm/admin-lte@3.2.0/dist/css/adminlte.min.css")

	// Add Font Awesome
	webpage.AddStyle("https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.4.0/css/all.min.css")

	// Add Google Font: Source Sans Pro
	webpage.AddStyle("https://fonts.googleapis.com/css?family=Source+Sans+Pro:300,400,400i,700&display=fallback")

	// Add custom CSS
	if style := templateStyle(); style != "" {
		webpage.AddStyle(style)
	}

	// Add jQuery (required for Bootstrap and AdminLTE)
	webpage.AddScript("https://code.jquery.com/jquery-3.6.0.min.js")

	// Add Bootstrap 4 JS Bundle with Popper (required for AdminLTE)
	webpage.AddScript("https://cdn.jsdelivr.net/npm/bootstrap@4.6.2/dist/js/bootstrap.bundle.min.js")

	// Add overlayScrollbars (required by AdminLTE)
	webpage.AddScript("https://cdn.jsdelivr.net/npm/overlayscrollbars@1.13.1/js/jquery.overlayScrollbars.min.js")

	// Add AdminLTE JS
	webpage.AddScript("https://cdn.jsdelivr.net/npm/admin-lte@3.2.0/dist/js/adminlte.min.js")

	// Add custom JavaScript
	if script := templateScript(); script != "" {
		webpage.AddScript(script)
	}

	// Add theme CSS or default AdminLTE CSS
	theme := dashboard.GetTheme()
	if theme != "" && theme != ThemeDefault {
		// Check if it's a known theme
		if theme == ThemeDark || theme == ThemeLight {
			webpage.AddStyle("https://cdn.jsdelivr.net/npm/admin-lte@3.2.0/dist/css/skins/" + theme + ".min.css")
		}
	} else {
		// Default theme
		webpage.AddStyle("https://cdn.jsdelivr.net/npm/admin-lte@3.2.0/dist/css/skins/skin-blue.min.css")
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
	footer := hb.NewFooter().Class("main-footer")
	footer.Child(hb.NewDiv().Class("float-right d-none d-sm-block").Child(
		hb.NewDiv().Style("font-size: 85%;").Text("AdminLTE v3.2.0"),
	))
	footer.Child(hb.NewStrong().Text("Copyright &copy; 2023").Child(
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

	// Add initialization script
	webpage.AddScript(`
		document.addEventListener('DOMContentLoaded', function() {
			// Initialize AdminLTE
			if (typeof $ !== 'undefined' && $.fn.AdminLTE) {
				$('body').AdminLTE({
					sidebarExpandOnHover: false,
					enableControlSidebar: true,
					controlSidebarOptions: {}
				});
			}
		});
	`)

	// Generate the final HTML
	return webpage.ToHTML()
}
