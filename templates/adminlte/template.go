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
func (t *Template) layout(dashboard types.DashboardInterface) string {
	content := dashboard.GetContent()
	layout := hb.NewBorderLayout()
	layout.AddTop(topNavigation(dashboard), hb.BORDER_LAYOUT_ALIGN_LEFT, hb.BORDER_LAYOUT_ALIGN_MIDDLE)
	layout.AddCenter(hb.Raw(content), hb.BORDER_LAYOUT_ALIGN_LEFT, hb.BORDER_LAYOUT_ALIGN_TOP)
	return layout.ToHTML()
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
	wrapper := hb.NewDiv().Class("wrapper")

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

	// Add content wrapper
	contentWrapper := hb.NewDiv().Class("content-wrapper")
	contentWrapper.Child(hb.Raw(t.layout(dashboard)))
	wrapper.Child(contentWrapper)

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
